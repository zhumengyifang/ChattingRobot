package main

import (
	"fmt"
	"main/Util"
	"time"
)

const userId1 = "userId1"
const userId2 = "userId2"

func Chatting()  {
	var chattingInfo = Util.BuildChatting("你好！", userId1)
	var userId = userId1
	for i := 0; i < 10; i++ {
		var json = Util.ConvertJson(*chattingInfo)
		var resultText = chatting(userId, json)
		time.Sleep(time.Second)

		if i%2 != 0 {
			userId = userId1
		} else {
			userId = userId2
		}
		Util.UpdateChatting(userId, resultText, chattingInfo)
	}
}

func chatting(userId string, json string) string {
	var body = Util.HttpPost(json)
	var resultText = Util.GetText(body)
	fmt.Println()
	fmt.Println(userId + ":" + resultText)
	return resultText
}