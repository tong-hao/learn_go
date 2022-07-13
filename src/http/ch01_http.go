package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpPostJson() {
	url:= "https://www.denlery.top/api/v1/login"
	jsonStr :=[]byte(`{ "username": "auto", "password": "auto123123" }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	// handle error
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		fmt.Println(statusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statusCode)
	fmt.Println(resp.Header)
}

func main() {
	httpPostJson()
}