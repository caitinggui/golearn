package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"strings"
)

func SimpleGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get err:", err)
		return

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func TimeoutGet(url string) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get err:", err)
		return

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func SimplePostJson(url string) {
	contentType := "application/json"
	values := map[string]string{"username": "username", "password": "password"}
	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("get err:", err)
		return

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func SimplePostForm(url string) {
	contentType := "application/x-www-form-urlencoded"
	values := strings.NewReader("ClientId=123&password=456")

	resp, err := http.Post(url, contentType, values)
	if err != nil {
		fmt.Println("get err:", err)
		return

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func AdvanceClient(url string) {
	contentType := "application/json"
	originValues := map[string]string{"username": "put", "password": "password"}
	jsonValue, _ := json.Marshal(originValues)
	values := bytes.NewBuffer(jsonValue)
	//contentType := "application/x-www-form-urlencoded"
	//values := strings.NewReader("ClientId=put&password=456")

	client := http.Client{Timeout: 3 * time.Second}

	req, _ := http.NewRequest("PUT", url, values)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Cookie", "name=54")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get err:", err)
		return

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func main() {
	url := "http://127.0.0.1:8088?hell=/dkjk&kjd=jkeh%jjkf+jdk"
	SimpleGet(url)
	TimeoutGet(url)
	SimplePostForm(url)
	SimplePostJson(url)
	AdvanceClient(url)

}
