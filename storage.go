package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var storagePath string = "storage/"
var name fileNames = fileNames{credentials:"cred", servers:"servers"}
type fileNames struct {
	credentials, servers string
}

func SaveLoginData(username string, password string) {
	data := []byte(username + "\n" + password)
	checkStorageAccessibility()
	err := ioutil.WriteFile(storagePath + name.credentials, data, 0644)
	if err != nil {
		logrus.Panic("Failed to write to credentials file. Error: ", err)
	}
}
func LoadLoginData() (string, string, error) {
	data, err := ioutil.ReadFile(storagePath + name.credentials)
	if err != nil {
		logrus.Error("Failed reading from credentials file. Error: ", err)
		return "", "", ErrFile.Wrap(err,"Failed reading from credentials file")
	}
	cred, parse_err := ParseLoginData(data)
	if parse_err != nil {
		return "", "", ErrParse.Wrap(parse_err, "Couldn't load credentials data")
	}
	return cred[0], cred[1], nil
}
func SaveServerData(serverlist []byte) {
	checkStorageAccessibility()
	err := ioutil.WriteFile(storagePath + name.servers, serverlist, 0644)
	if err != nil {
		logrus.Panic("Failed to write to server file. Error: ", err)
	}
}
func LoadServerData() ([]byte, error) {
	data, err := ioutil.ReadFile(storagePath + name.servers)
	if err != nil {
		logrus.Error("Failed reading from servers file. Error: ", err)
		return nil, ErrFile.Wrap(err,"Failed reading from servers file")
	}
	return data, nil
}
func checkStorageAccessibility() {
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		os.MkdirAll(storagePath, os.ModePerm)
	}
}