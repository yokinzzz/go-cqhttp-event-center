package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestHandleApiGWRequest(t *testing.T) {
	var jsonReq IntegrationRequest
	//jsonStr := "{\"requestContext\":{\"serviceId\":\"service-f94sy04v\",\"path\":\"/test/{path}\",\"httpMethod\":\"POST\",\"requestId\":\"c6af9ac6-7b61-11e6-9a41-93e8deadbeef\",\"identity\":{\"secretId\":\"abdcdxxxxxxxsdfs\"},\"sourceIp\":\"10.0.2.14\",\"stage\":\"release\"},\"headers\":{\"accept-Language\":\"en-US,en,cn\",\"accept\":\"text/html,application/xml,application/json\",\"host\":\"service-3ei3tii4-251000691.ap-guangzhou.apigateway.myqloud.com\",\"user-Agent\":\"User Agent String\"},\"body\":\"{\\\"font\\\":0,\\\"message\\\":[{\\\"data\\\":{\\\"text\\\":\\\"jdjd\\\"},\\\"type\\\":\\\"text\\\"}],\\\"message_id\\\":324791218,\\\"message_type\\\":\\\"private\\\",\\\"post_type\\\":\\\"message\\\",\\\"raw_message\\\":\\\"jdjd\\\",\\\"self_id\\\":1007626060,\\\"sender\\\":{\\\"age\\\":0,\\\"nickname\\\":\\\"Veleteen\\\",\\\"sex\\\":\\\"unknown\\\",\\\"user_id\\\":535175166},\\\"sub_type\\\":\\\"friend\\\",\\\"target_id\\\":1007626060,\\\"time\\\":1629346409,\\\"user_id\\\":535175166}\",\"pathParameters\":{\"path\":\"value\"},\"queryStringParameters\":{\"foo\":\"bar\"},\"headerParameters\":{\"Refer\":\"10.0.2.14\"},\"stageVariables\":{\"stage\":\"release\"},\"path\":\"/test/value\",\"queryString\":{\"foo\":\"bar\",\"bob\":\"alice\"},\"httpMethod\":\"POST\"}"
	jsonStr := "{\"requestContext\":{\"serviceId\":\"service-f94sy04v\",\"path\":\"/test/{path}\",\"httpMethod\":\"POST\",\"requestId\":\"c6af9ac6-7b61-11e6-9a41-93e8deadbeef\",\"identity\":{\"secretId\":\"abdcdxxxxxxxsdfs\"},\"sourceIp\":\"10.0.2.14\",\"stage\":\"release\"},\"headers\":{\"accept-Language\":\"en-US,en,cn\",\"accept\":\"text/html,application/xml,application/json\",\"host\":\"service-3ei3tii4-251000691.ap-guangzhou.apigateway.myqloud.com\",\"user-Agent\":\"User Agent String\"},\"body\":\"{\\\"font\\\":0,\\\"message\\\":[{\\\"data\\\":{\\\"text\\\":\\\"jdjd\\\"},\\\"type\\\":\\\"text\\\"}],\\\"message_id\\\":324791218,\\\"message_type\\\":\\\"private\\\",\\\"post_type\\\":\\\"message\\\",\\\"raw_message\\\":\\\"jdjd\\\",\\\"self_id\\\":1007626060,\\\"sender\\\":{\\\"age\\\":0,\\\"nickname\\\":\\\"Veleteen\\\",\\\"sex\\\":\\\"unknown\\\",\\\"user_id\\\":535175166},\\\"sub_type\\\":\\\"friend\\\",\\\"target_id\\\":1007626060,\\\"time\\\":1629346409,\\\"user_id\\\":535175166}\",\"pathParameters\":{\"path\":\"value\"},\"queryStringParameters\":{\"foo\":\"bar\"},\"headerParameters\":{\"Refer\":\"10.0.2.14\"},\"stageVariables\":{\"stage\":\"release\"},\"path\":\"/test/value\",\"queryString\":{\"foo\":\"bar\",\"bob\":\"alice\"},\"httpMethod\":\"POST\"}"
	err := json.Unmarshal([]byte(jsonStr), &jsonReq)
	jsonReq.Body = "{\"anonymous\":null,\"font\":0,\"group_id\":398438555,\"message\":[{\"data\":{\"qq\":\"1007626060\"},\"type\":\"at\"},{\"data\":{\"text\":\" .r 5d2\"},\"type\":\"text\"}],\"message_id\":-1516515882,\"message_seq\":1124531,\"message_type\":\"group\",\"post_type\":\"message\",\"raw_message\":\"[CQ:at,qq=1007626060] 测试\",\"self_id\":1007626060,\"sender\":{\"age\":0,\"area\":\"\",\"card\":\"有猫\",\"level\":\"\",\"nickname\":\"Veleteen\",\"role\":\"admin\",\"sex\":\"unknown\",\"title\":\"伟酱\",\"user_id\":535175166},\"sub_type\":\"normal\",\"time\":1629346489,\"user_id\":535175166}"
	if err != nil {
		panic(err)
	} else {
		fmt.Println("no error")
	}
	actual, _ := handleApiGWRequest(context.TODO(), jsonReq)
	fmt.Println(actual)
}
