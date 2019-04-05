package Util

import (
	"main/model"
)

const appKey = "6f884bbd0b9841bf846f13947a85b13c"

func UpdateChatting(userId string, text string, chattingInfo model.Chatting) model.Chatting {
	chattingInfo.UserInfo.UserId = userId
	chattingInfo.Perception.InputText.Text = text
	return chattingInfo
}

func BuildChatting(text string, userId string) model.Chatting {
	chatting := model.Chatting{ReqType: 0}
	chatting.Perception = buildPerception(text)
	chatting.UserInfo = buildUserInfo(userId)
	return chatting
}

func buildPerception(text string) model.Perception {
	perception := model.Perception{buildInputText(text)}
	return perception
}

func buildInputText(text string) model.InputText {
	inputText := model.InputText{text}
	return inputText
}

func buildUserInfo(userId string) model.UserInfo {
	return model.UserInfo{appKey, userId}
}
