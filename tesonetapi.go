package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)
type servers struct {
	Name string
	Distance int
}
type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type WebRequest interface {
	GetBytes(url, method string, headers map[string]string, postdata []byte) ([]byte, error)
}
type GetWebRequest struct {}
var requester GetWebRequest
func (GetWebRequest) GetBytes(url, method string, headers map[string]string, postdata []byte) ([]byte, error) {
	var requestBody io.Reader
	if postdata == nil {
		requestBody = nil
	} else {
		requestBody = bytes.NewBuffer(postdata)
	}
	request, _ := http.NewRequest(method, url, requestBody)
	for hname, hvalue:= range headers {
		request.Header.Add(hname, hvalue)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrNetwork.Wrap(err,"no connection could be made")
	}
	if response.StatusCode >= 400 { // error codes start at 400
		msg := "Request was refused: " + response.Status
		return nil, ErrHTTP.New(msg)
	}
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
}
func GetToken(username, password string) (string, error){
	authUrl := "http://playground.tesonet.lt/v1/tokens"
	tokenRequest, _ := json.Marshal(login{username,password})
	headers := map[string]string{"Content-type": "application/json"}
	body, _ := requester.GetBytes(authUrl,"POST", headers, tokenRequest)
	var jsonResponse map[string]string
	err_json := JsonBytesToStruct(body, &jsonResponse)
	if err_json != nil {
		return "", ErrJSON.Wrap(err_json,"Could not parse token data")
	}
	return jsonResponse["token"], nil
}
func GetServers(token string) ([]byte, error) {
	serverUrl := "http://playground.tesonet.lt/v1/servers"
	headers := map[string]string{"Content-type": "application/json",
		"Authorization": "Bearer " + token}
	body, err := requester.GetBytes(serverUrl,"GET",headers,nil)
	if err != nil {
		return nil, ErrHTTP.Wrap(err,"Failed to fetch data from API")
	}
	return body, nil
	}
