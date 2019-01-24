package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)
type testingWebRequest struct {
	json string
}
func (t testingWebRequest) GetBytes(url, method string, headers map[string]string, postdata []byte) ([]byte, error) {

	return []byte(t.json), nil //some fake data comes from here
}
func TestGetTokenWithCorrectAuth(t *testing.T) {
	token, err := GetToken("tesonet", "partyanimal")
	if err != nil || token == "" {
		t.Errorf("Failed to get token with correct credentials. Expected: non-empty token. Got: %v", err.Error())
	}
}
func TestGetTokenWithIncorrectAuth(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	_, err := GetToken("tesonet", "wrongpassword")
	if a, ok := err.(CliError); !(ok && a.errorType == ErrWeb){
		t.Errorf("Failed to receive error with incorrect credentials. Expected: %v(ErrWeb). Got: %v", a.errorType, err.Error())
	}
}

func TestGetServersWithCorrectToken(t *testing.T) {
	_, err := GetServers("f9731b590611a5a9377fbd02f247fcdf")
	if err != nil {
		t.Errorf("Failed to get server list with correct token. Expected: server list. Got: %v", err.Error())
	}
}
func TestGetServersWithIncorrectToken(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	_, err := GetServers("f9731b590611a5a9377fbd02f247abcd")
	if a, ok := err.(CliError); !(ok && a.errorType == ErrWeb){
		t.Errorf("Failed to receive error with incorrect token. Expected: %v(ErrWeb). Got: %v", a.errorType, err.Error())
	}
}
func TestGetBytesWithIncorrectUrl(t *testing.T) {
	req := WebRequestData{url:"http://nonexistingwebsite404.lt"}
	_, err := req.GetBytes()
	if a, ok := err.(CliError); !(ok && a.errorType == ErrWeb){
		t.Errorf("Failed to receive error connecting to a non-existing website. Expected: %v(ErrWeb). Got: %v", a.errorType, err.Error())
	}

}