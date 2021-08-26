package event

type Segment struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

/* 消息段格式化接口 */
type SegmentFormatter interface {
	ConvertToSegment() Segment
}

// @消息段
type MentionSegment struct {
	QQ   int64  `json:"qq"`   // @的 QQ 号, all 表示全体成员
	Name string `json:"name"` // 当在群中找不到此QQ号的名称时才会生效
}

func (mention *MentionSegment) ConvertToSegment() Segment {
	return Segment{Type: At, Data: map[string]interface{}{
		"qq":   mention.QQ,
		"name": mention.Name,
	}}
}

// 文本消息段
type TextSegment struct {
	Text string `json:"text"`
}

func (text *TextSegment) ConvertToSegment() Segment {
	return Segment{Type: Text, Data: map[string]interface{}{
		"text": text.Text,
	}}
}

// 消息段字段名
const (
	Type string = "type" // 消息类型
	Data string = "data" // 消息数据
)

// 消息类型
const (
	At   string = "at"   // @消息类型
	Text string = "text" // 文本消息类型
)
