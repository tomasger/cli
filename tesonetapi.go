package main

import (
	"bytes"
	"encoding/json"
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
	tokenRequest, _ := json.Marshal(values)
	resp, err := http.Post(authUrl, "application/json", bytes.NewBuffer(tokenRequest))
	if err != nil {
		return "", ErrNetwork.Wrap(err,"no connection could be made")
	}
	if resp.StatusCode >= 400 { // error codes start at 400
			msg := "Request was refused: "
			msg += string(resp.StatusCode) + resp.Status
		return "", ErrHTTP.New(msg)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var jsonResponse map[string]string
	err_json := JsonBytesToStruct(body, &jsonResponse)
	if err_json != nil {
		return "", ErrJSON.Wrap(err_json,"Could not parse token data")
	}
	return jsonResponse["token"], nil
}
func GetServers(token string) ([]byte, error) {
	serverUrl := "http://playground.tesonet.lt/v1/servers"
	client := &http.Client{}
	request, _ := http.NewRequest("GET", serverUrl, nil)
	authval := "Bearer " + token
	request.Header.Add("Authorization", authval)
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrNetwork.Wrap(err,"no connection could be made")
	}
	if response.StatusCode >= 400 { // error codes start at 400
		msg := "Request was refused: "
		msg += string(response.StatusCode) + response.Status
		return nil, ErrHTTP.New(msg)
	}
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
	}
