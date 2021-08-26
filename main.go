package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"

	"github.com/mitchellh/mapstructure"
	"github.com/yokinzzz/go-cqhttp-event-center/event"
	"github.com/yokinzzz/go-cqhttp-event-center/handler"
)

/* IntegrationRequest and IntegrationResponse according to tencent scf doc:
https://cloud.tencent.com/document/product/583/12513 */

type IntegrationRequest struct {
	RequestContext
	Body       string // 记录实际请求转换为String字符串后的内容
	Path       string // 记录实际请求的完整 Path 信息
	HttpMethod string // 记录实际请求的 HTTP 方法
	// todo: modify below type
	Headers               map[string]string
	PathParameters        map[string]string
	QueryStringParameters map[string]string
	HeaderParameters      map[string]string
	QueryString           map[string]string
}

type RequestContext struct {
	ServiceId  string // API 网关的服务 ID
	Path       string // API 的路径
	HttpMethod string // API 的Http方法
	RequestId  string // 标识当前这次请求的唯一 ID
	SourceIp   string // 标识请求来源 IP
	Stage      string // 指向请求来源 API 所在的环境
	Identity
}

/* 标识用户的认证方法和认证的信息 */
type Identity struct {
	SecretId string
}

type IntegrationResponse struct {
	IsBase64Encoded bool  `json:"isBase64Encoded"` // 指明 body 内的内容是否为 Base64 编码后的二进制内容，取值需要为 JSON 格式的 true 或 false
	StatusCode      int32 `json:"statusCode"`      // HTTP 返回的状态码，取值需要为 Integer 值
	// todo: modify below type
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// post_type
const (
	Message         string = "message" // 消息
	GROUP_MESSAGE   string = "group"   // 群消息
	PRIVATE_MESSAGE string = "private" // 私聊

	Notice string = "notice" // 通知

	Request string = "request" // 请求
)

// request_type
const (
	Group  string = "group"  // 加群请求／邀请
	Friend string = "friend" // 加好友请求
)

// field name
const (
	POST_TYPE    string = "post_type"
	REQUEST_TYPE string = "request_type"
	MESSAGE_TYPE string = "message_type"
	NOTICE_TYPE  string = "notice_type"
)

/* dispatch api gateway integration request to different message handler depending on message type */
func handleApiGWRequest(ctx context.Context, request IntegrationRequest) (IntegrationResponse, error) {
	fmt.Printf("%s\n", request.Body)
	headers := map[string]string{"Content-Type": "application/json"}
	var body map[string]interface{}
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		fmt.Printf("json unmarshal IntegrationRequest.Body failed, Body: %s\nError: %s\n", request.Body, err.Error())
	}
	switch body[POST_TYPE] {
	case Message:
		if body[MESSAGE_TYPE] == GROUP_MESSAGE {
			event := &event.GroupMessageEvent{}
			mapstructure.Decode(body, event)
			return IntegrationResponse{Headers: headers, StatusCode: 200, Body: handler.GroupMessageHandler(*event).String()}, nil
		} else if body[MESSAGE_TYPE] == PRIVATE_MESSAGE {
			event := &event.PrivateMessageEvent{}
			mapstructure.Decode(body, &event)
			return IntegrationResponse{Headers: headers, StatusCode: 200, Body: handler.PrivateMessageHandler(*event).String()}, nil
		}
	case Notice:
		// todo 撤回消息
	case Request:
		if body[REQUEST_TYPE] == Group {
			event := &event.InviteGroupEvent{}
			mapstructure.Decode(body, &event)
			return IntegrationResponse{Headers: headers, StatusCode: 200, Body: handler.GroupRequestHandler(*event).String()}, nil
		}
	}
	return IntegrationResponse{}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(handleApiGWRequest)
}
