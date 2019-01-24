package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var storagePath string = "storage/"
var fileNames FileNames = FileNames{credentials: "cred", servers:"servers"}
type FileNames struct {
	credentials, servers string
}

func SaveLoginData(username string, password string) {
	data := []byte(username + "\n" + password)
	ensureStorageAccessibility()
	err := ioutil.WriteFile(storagePath + fileNames.credentials, data, 0644)
	if err != nil {
		logrus.Panic("Failed to write to credentials file. Error: ", err)
	}
}
func LoadLoginData() (string, string, error) {
	data, err := ioutil.ReadFile(storagePath + fileNames.credentials)
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
	ensureStorageAccessibility()
	err := ioutil.WriteFile(storagePath + fileNames.servers, serverlist, 0644)
	if err != nil {
		logrus.Panic("Failed to write to server file. Error: ", err)
	}
}
func LoadServerData() ([]byte, error) {
	data, err := ioutil.ReadFile(storagePath + fileNames.servers)
	if err != nil {
		logrus.Error("Failed reading from servers file. Error: ", err)
		return nil, ErrFile.Wrap(err,"Failed reading from servers file")
	}
	return data, nil
}
func ensureStorageAccessibility() {
	// check if storage path exists. Create the folder for it in case it does not
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		os.MkdirAll(storagePath, os.ModePerm)
	}
}