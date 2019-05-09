package Util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const url = "http://openapi.tuling123.com/openapi/api/v2"

func HttpPost(json string) string {
	jsonStr := []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
