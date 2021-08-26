package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/yokinzzz/go-cqhttp-event-center/event"
)

type GroupMessageQuickReply struct {
	Reply       string `json:"reply"`        // 要回复的内容
	AutoEscape  bool   `json:"auto_escape"`  // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 reply 字段是字符串时有效
	AtSender    bool   `json:"at_sender"`    // 是否要在回复开头 at 发送者 ( 自动添加 ) , 发送者是匿名用户时无效
	Delete      bool   `json:"delete"`       // 撤回该条消息
	Kick        bool   `json:"kick"`         // 把发送者踢出群组 ( 需要登录号权限足够 ) , 不拒绝此人后续加群请求, 发送者是匿名用户时无效
	Ban         bool   `json:"ban"`          // 把发送者禁言 ban_duration 指定时长, 对匿名用户也有效
	BanDuration int64  `json:"ban_duration"` // 禁言时长
}

/* 群组消息 */

func (reply GroupMessageQuickReply) String() string {
	jsonStr, _ := json.Marshal(&reply)
	return string(jsonStr)
}

func GroupMessageHandler(groupEvent event.GroupMessageEvent) GroupMessageQuickReply {
	var hasBeenMentioned bool
	for _, segment := range groupEvent.Message {
		messageType := segment.Type
		switch messageType {
		case event.At:
			var messageData event.MentionSegment
			mapstructure.Decode(segment.Data, &messageData)
			if groupEvent.SelfId == messageData.QQ {
				hasBeenMentioned = true
			}
			continue
		case event.Text:
			if hasBeenMentioned {
				var messageData event.TextSegment
				mapstructure.Decode(segment.Data, &messageData)
				return textHandler(messageData.Text)
			}
			continue
		}
	}
	return GroupMessageQuickReply{}
}

// handler text segment
func textHandler(text string) GroupMessageQuickReply {
	text = strings.TrimSpace(text)
	// match string starts with . following with lower case command and command content comes after a space
	regex, _ := regexp.Compile("^.[a-z]+ ")
	matchIndex := regex.FindStringIndex(text)
	if matchIndex != nil {
		command := strings.TrimSpace(text[matchIndex[0]:matchIndex[1]]) // get rid of command following space
		content := text[matchIndex[1]:]
		switch command {
		case ROLL:
			return GroupMessageQuickReply{Reply: roll(content)}
		default:
			return GroupMessageQuickReply{Reply: fmt.Sprintf("命令[%s]不存在", command)}
		}
	}
	return GroupMessageQuickReply{}
}

// command
const (
	ROLL string = ".r"
)

// roll within DnD rule
func roll(exp string) string {
	var strBuilder strings.Builder
	var sum int = 0
	diceIndex := strings.Index(exp, "d")
	count, _ := strconv.Atoi(exp[0:diceIndex])
	max, _ := strconv.Atoi(exp[diceIndex+1:])
	if count > 512 {
		return "这么多次掷得过来吗？"
	}
	if max > 512 {
		return "这是多少面骰子..."
	}
	if count == 1 {
		return strconv.Itoa(rand.Intn(max))
	}
	for i := 0; i < count; i++ {
		random := rand.Intn(max)
		sum += random
		strBuilder.WriteString(strconv.Itoa(random))
		if i < count-1 {
			strBuilder.WriteString(" + ")
		} else {
			strBuilder.WriteString(" = " + strconv.Itoa(sum))
		}
	}
	return strBuilder.String()
}

/* 私聊消息 */

type PrivateMessageQuickReply struct {
	Reply      string `json:"reply"`       // 要回复的内容
	AutoEscape bool   `json:"auto_escape"` // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 reply 字段是字符串时有效
}

func (reply PrivateMessageQuickReply) String() string {
	jsonStr, _ := json.Marshal(&reply)
	return string(jsonStr)
}

func PrivateMessageHandler(event event.PrivateMessageEvent) PrivateMessageQuickReply {
	return PrivateMessageQuickReply{Reply: "Hi~"}
}
