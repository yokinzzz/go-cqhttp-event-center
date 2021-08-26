package handler

import (
	"encoding/json"

	"github.com/yokinzzz/go-cqhttp-event-center/event"
)

type GroupRequestQuickReply struct {
	Approve bool   `json:"approve"` // 是否同意请求／邀请
	Reason  string `json:"reason"`  // 拒绝理由 ( 仅在拒绝时有效 )
}

func (reply GroupRequestQuickReply) String() string {
	jsonStr, _ := json.Marshal(&reply)
	return string(jsonStr)
}

func GroupRequestHandler(event event.InviteGroupEvent) GroupRequestQuickReply {
	return GroupRequestQuickReply{false, "暂时不接受加群请求"}
}
