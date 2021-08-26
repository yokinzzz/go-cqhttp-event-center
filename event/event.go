package event

/* 消息事件 */
// 私聊消息事件
type PrivateMessageEvent struct {
	Time                 int64           `json:"time"`         // 事件发生的时间戳
	SelfId               int64           `json:"self_id"`      // 收到事件的机器人 QQ 号
	PostType             string          `json:"post_type"`    // 上报类型
	MessageType          string          `json:"message_type"` // 消息类型
	SubType              string          `json:"sub_type"`     // 消息子类型, 如果是好友则是 friend, 如果是群临时会话则是 group, 如果是在群中自身发送则是 group_self
	TempSource           string          `json:"temp_source"`  // 临时会话来源
	MessageId            int32           `json:"message_id"`   // 消息 ID
	UserId               int64           `json:"user_id"`      // 发送者 QQ 号
	RawMessage           string          `json:"raw_message"`  // 原始消息内容
	Font                 int32           `json:"font"`         // 字体
	PrivateMessageSender `json:"sender"` // 发送人信息
	Message              []Segment       `json:"message"` // 消息段
}

type PrivateMessageSender struct {
	UserId   int64  `json:"user_id"`  // 发送者 QQ 号
	Nickname string `json:"nickname"` // 昵称
	Sex      string `json:"sex"`      // 性别, male 或 female 或 unknown
	Age      int32  `json:"age"`      // 年龄
}

// 群聊消息事件
type GroupMessageEvent struct {
	Time               int64              `json:"time"`         // 事件发生的时间戳
	SelfId             int64              `json:"self_id"`      // 收到事件的机器人 QQ 号
	PostType           string             `json:"post_type"`    // 上报类型
	MessageType        string             `json:"message_type"` // 消息类型
	SubType            string             `json:"sub_type"`     // 消息子类型, 如果是好友则是 friend, 如果是群临时会话则是 group, 如果是在群中自身发送则是 group_self
	MessageId          int32              `json:"message_id"`   // 消息 ID
	GroupId            int64              `json:"group_id"`     // 群号
	UserId             int64              `json:"user_id"`      // 发送者 QQ 号
	RawMessage         string             `json:"raw_message"`  // 原始消息内容
	Font               int32              `json:"font"`         // 字体
	GroupMessageSender `json:"sender"`    // 发送人信息
	Anonymous          `json:"anonymous"` // 匿名信息, 如果不是匿名消息则为 null
	Message            []Segment          `json:"message"` // 消息段
}

type GroupMessageSender struct {
	UserId   int64  `json:"user_id"`  // 发送者 QQ 号
	Nickname string `json:"nickname"` // 昵称
	Sex      string `json:"sex"`      // 性别, male 或 female 或 unknown
	Age      int32  `json:"age"`      // 年龄
	Area     string `json:"area"`     // 地区
	Card     string `json:"card"`     // 群名片／备注
	Level    string `json:"level"`    // 成员等级
	Role     string `json:"role"`     // 角色, owner 或 admin 或 member
	Title    string `json:"title"`    // 专属头衔
}

type Anonymous struct {
	ID   int64  `json:"id"`   // 匿名用户 ID
	Name string `json:"name"` // 匿名用户名称
	Flag string `json:"flag"` // 匿名用户 flag, 在调用禁言 API 时需要传入
}

/* 通知事件 */
// 群消息撤回
type GroupMessageRecallEvent struct {
	Time       int64  `json:"time"`        // 事件发生的时间戳
	SelfId     int64  `json:"self_id"`     // 收到事件的机器人 QQ 号
	PostType   string `json:"post_type"`   // 上报类型
	NoticeType string `json:"notice_type"` // 通知类型
	GroupId    int64  `json:"group_id"`    // 群号
	UserId     int64  `json:"user_id"`     // 消息发送者 QQ 号
	OperatorId int64  `json:"operator_id"` // 操作者 QQ 号
	MessageId  int64  `json:"message_id"`  // 被撤回的消息 ID
}

/* 请求事件 */
// 加群请求/邀请
type InviteGroupEvent struct {
	Time        int64  `json:"time"`         // 事件发生的时间戳
	SelfId      int64  `json:"self_id"`      // 收到事件的机器人 QQ 号
	PostType    string `json:"post_type"`    // 上报类型
	RequestType string `json:"request_type"` // 通知类型
	SubType     string `json:"sub_type"`     //请求子类型, 分别表示加群请求、邀请登录号入群
	GroupId     int64  `json:"group_id"`     // 群号
	UserId      int64  `json:"user_id"`      // 发送请求的 QQ 号
	Comment     string `json:"comment"`      // 验证信息
	Flag        string `json:"flag"`         // 请求 flag, 在调用处理请求的 API 时需要传入
}
