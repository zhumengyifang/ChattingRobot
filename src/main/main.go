package main

import (
	"fmt"
	"main/Util"
	"time"
)

const userId1 = "userId1"
const userId2 = "userId2"

func main() {
	var chattingInfo = Util.BuildChatting("你好！", userId1)

	var userId = userId1
	for i := 0; i < 10; i++ {
		var json = Util.ConvertJson(chattingInfo)

		var resultText = chatting(userId, json)
		time.Sleep(time.Second)

		if i%2 != 0 {
			userId = userId1
		} else {
			userId = userId2
		}

		chattingInfo = Util.UpdateChatting(userId, resultText, chattingInfo)
	}

}

func chatting(userId string, json string) string {
	var body, resultErrs = Util.HttpPost(json)
	if resultErrs != nil {
		fmt.Println(resultErrs.Error())
	}

	var resultText, errs = Util.GetText(body)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println()
	fmt.Println(userId + ":" + resultText)
	return resultText
}
