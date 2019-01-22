package main

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
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
type WebRequestData struct {
	url string
	method string
	headers map[string]string
	postdata []byte
}
type WebRequest interface {
	GetBytes(url, method string, headers map[string]string, postdata []byte) ([]byte, error)
}
func (requester WebRequestData) GetBytes() ([]byte, error) {
	var requestBody io.Reader
	if requester.postdata == nil {
		requestBody = nil
	} else {
		requestBody = bytes.NewBuffer(requester.postdata)
	}
	request, _ := http.NewRequest(requester.method, requester.url, requestBody)
	for hname, hvalue:= range requester.headers {
		request.Header.Add(hname, hvalue)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrWeb.Wrap(err,"no connection could be made")
	}
	if response.StatusCode >= 400 { // error codes start at 400
		msg := "Request was refused: " + response.Status
		return nil, ErrWeb.New(msg)
	}
	body, _ := ioutil.ReadAll(response.Body)
	return body, nil
}
func GetToken(username, password string) (string, error){
	authUrl := "http://playground.tesonet.lt/v1/tokens"
	tokenRequest, _ := json.Marshal(login{username,password})
	headers := map[string]string{"Content-type": "application/json"}
	loginRequest := WebRequestData{authUrl, "POST", headers, tokenRequest}
	body, err := loginRequest.GetBytes()
	if err != nil {
		logrus.Error("Could not fetch the token from the API")
		return "", ErrWeb.Wrap(err,"Failed to fetch data from API")
	}
	var jsonResponse map[string]string
	err_json := JsonBytesToStruct(body, &jsonResponse)
	if err_json != nil {
		logrus.Error("Failed to parse response from token API")
		return "", ErrJSON.Wrap(err_json,"Could not parse token data")
	}
	return jsonResponse["token"], nil
}
func GetServers(token string) ([]byte, error) {
	serverUrl := "http://playground.tesonet.lt/v1/servers"
	headers := map[string]string{"Content-type": "application/json",
		"Authorization": "Bearer " + token}
	serversReqest := WebRequestData{serverUrl, "GET", headers, nil}
	body, err := serversReqest.GetBytes()
	if err != nil {
		logrus.Error("Could not fetch the server list from the API")
		return nil, ErrWeb.Wrap(err,"Failed to fetch data from API")
	}
	return body, nil
	}
