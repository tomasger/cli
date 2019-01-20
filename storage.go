package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func SaveLoginData(username string, password string) {
	data := []byte(username + "\n" + password)
	checkStorageAccessibility()
	err := ioutil.WriteFile("storage/cred", data, 0644)
	if err != nil {
		panic(err)
	}
}
func LoadLoginData() (string, string, error) {
	dat, err := ioutil.ReadFile("storage/cred")
	if err != nil {
		return "", "", ErrFile.Wrap(err,"Failed reading from credentials file")
	}
	data := strings.Split(string(dat), "\n")
	return data[0], data[1], nil
}
func SaveServerData(serverlist []byte) {
	checkStorageAccessibility()
	err := ioutil.WriteFile("storage/servers", serverlist, 0644)
	if err != nil {
		panic(err)
	}
}
func LoadServerData() ([]byte, error) {
	data, err := ioutil.ReadFile("storage/servers")
	if err != nil {
		return nil, ErrFile.Wrap(err,"Failed reading from credentials file")
	}
	return data, nil
}
func checkStorageAccessibility() {
	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		os.Mkdir("storage", 0644)
	}
}