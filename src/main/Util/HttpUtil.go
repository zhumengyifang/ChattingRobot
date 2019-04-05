package Util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://openapi.tuling123.com/openapi/api/v2"

func HttpPost(json string) (string, error) {
	jsonStr := []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if !(resp.StatusCode == 200) {
		return "", errors.New("Get Text Error")
	}
	return string(body), nil
}
