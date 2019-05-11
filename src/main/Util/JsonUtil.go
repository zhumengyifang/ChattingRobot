package Util

import (
	"chattingRobot/src/main/model"
	"encoding/json"
	"github.com/gogf/gf/g/encoding/gjson"
)

func GetText(data string) string {
	if j, err := gjson.DecodeToJson([]byte(data)); err != nil {
		panic("Get Text Error")
	} else {
		var result = j.GetString("results.0.values.text")
		return result
	}
}

func ConvertJson(chattingInfo model.Chatting) string {
	jsons, errs := json.Marshal(chattingInfo)
	if errs != nil {
		panic(errs.Error())
	}
	return string(jsons)
}
