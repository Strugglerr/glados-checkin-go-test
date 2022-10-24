package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Push to send messages by pushplus(https://www.pushplus.plus/)

func Push(content string) {
	url := "https://www.pushplus.plus/send"

	type Body struct {
		Token   string `json:"token"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var body Body
	body.Token = os.Getenv("TOKEN")
	// to get pushplus token here
	body.Title = "GLADOS Check-in result reminder"
	body.Content = content
	// content to send
	byte, _ := json.Marshal(body)

	pushplus, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(byte)))
	if err != nil {
		panic(err)
	}
	pushplus.Header.Set("Accept", "application/json, text/plain, */*")
	pushplus.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// set Content-Type, json is necessary

	request, err := (&http.Client{}).Do(pushplus)
	if err != nil {
		fmt.Println(err)
	}
	defer request.Body.Close()
	// to send message by pushplus
}
