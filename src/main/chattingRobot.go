package main

import (
	"chattingRobot/src/main/Util"
	"chattingRobot/src/main/model"
	"fmt"
	"time"
)

func Chatting(count int) {
	var chattingInfo = Util.BuildChatting("你好！", model.UserId1)
	var userId = model.UserId1
	for i := 0; i < count; i++ {
		var json = Util.ConvertJson(*chattingInfo)
		var resultText = chatting(userId, json)
		time.Sleep(time.Second)

		if i%2 != 0 {
			userId = model.UserId1
		} else {
			userId = model.UserId2
		}
		Util.UpdateChatting(userId, resultText, chattingInfo)
	}
}

func chatting(userId string, json string) string {
	var body = Util.HttpPost(json)
	var resultText = Util.GetText(body)
	fmt.Println(userId + ":" + resultText)
	return resultText
}
