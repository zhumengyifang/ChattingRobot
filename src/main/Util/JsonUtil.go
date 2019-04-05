package Util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/g/encoding/gjson"
	"main/model"
)

func GetText(data string) (string, error) {
	if j, err := gjson.DecodeToJson([]byte(data)); err != nil {
		return "", errors.New("Get Text Error")
	} else {
		var result = j.GetString("results.0.values.text")
		return result, nil
	}
}

func ConvertJson(chattingInfo model.Chatting) string {
	jsons, errs := json.Marshal(chattingInfo)
	if errs != nil {
		fmt.Print(errs.Error())
	}
	return string(jsons)
}
