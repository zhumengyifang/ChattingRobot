package Util

import (
	"bytes"
	"chattingRobot/src/main/model"
	"io/ioutil"
	"net/http"
)

func HttpPost(json string) string {
	jsonStr := []byte(json)
	req, err := http.NewRequest("POST", model.Url, bytes.NewBuffer(jsonStr))
	if err!=nil{
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	c:=make(chan string)
	defer close(c)
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if !(resp.StatusCode == 200) {
			panic("Get Text Error")
		}
		c<-string(body)
	}()
	return <-c
}
