package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type servers struct {
	Name string
	Distance int
}
func GetToken(username, password string) (string, error){
	authUrl := "http://playground.tesonet.lt/v1/tokens"
	values := map[string]string{"username": username, "password": password}
	tokenRequest, enc_error := json.Marshal(values)
	if enc_error != nil {
		fmt.Println("FATAL: map to json conversion failed.")
	}
	resp, err := http.Post(authUrl, "application/json", bytes.NewBuffer(tokenRequest))
	if err != nil {
		fmt.Println("there's an error.", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonResponse map[string]string
	//dec_error := json.Unmarshal(body, &jsonResponse)
	//if dec_error != nil {
	//	fmt.Println("FATAL: api response to json conversion failed.")
	//}
	JsonBytesToStruct(body, &jsonResponse)
	if val, ok := jsonResponse["token"]; ok {
		return val, nil
	} else {
		return "Unauthorized", nil
	}
}
func GetServers(token string) ([]byte, error) {
	serverUrl := "http://playground.tesonet.lt/v1/servers"
	client := &http.Client{}
	request, _ := http.NewRequest("GET", serverUrl, nil)
	authval := "Bearer " + token
	request.Header.Add("Authorization", authval)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("there's an error")
	}
	body, _ := ioutil.ReadAll(response.Body)
	//serverlist := make([]servers, 0)
	//var serverlist []servers
	//dec_error := json.Unmarshal(body, &serverlist)
	//if dec_error != nil {
	//	fmt.Println("FATAL: api response to json conversion failed.")
	//}
	//fmt.Println(body)
	return body, nil
	}
