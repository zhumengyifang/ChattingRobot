package Util

import (
	"chattingRobot/src/main/model"
)

func UpdateChatting(userId string, text string, chattingInfo *model.Chatting) {
	chattingInfo.UserInfo.UserId = userId
	chattingInfo.Perception.InputText.Text = text
}

func BuildChatting(text string, userId string) *model.Chatting {
	chatting := model.Chatting{ReqType: 0}
	chatting.Perception = buildPerception(text)
	chatting.UserInfo = buildUserInfo(userId)
	return &chatting
}

func buildPerception(text string) *model.Perception {
	perception := model.Perception{InputText: buildInputText(text)}
	return &perception
}

func buildInputText(text string) *model.InputText {
	inputText := model.InputText{Text: text}
	return &inputText
}

func buildUserInfo(userId string) *model.UserInfo {
	return &model.UserInfo{ApiKey: model.AppKey, UserId: userId}
}
