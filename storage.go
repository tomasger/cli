package main

import (
	"io/ioutil"
	"strings"
)

func SaveLoginData(username string, password string) {
	data := []byte(username + "\n" + password)
	err := ioutil.WriteFile("storage/cred", data, 0644)
	if err != nil {
		panic(err)
	}
}
func LoadLoginData() (string, string) {
	dat, err := ioutil.ReadFile("storage/cred")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(dat), "\n")
	return data[0], data[1]
}
func SaveServerData(serverlist []byte) {
	err := ioutil.WriteFile("storage/servers", serverlist, 0644)
	if err != nil {
		panic(err)
	}
}
func LoadServerData() ([]byte) {
	data, err := ioutil.ReadFile("storage/servers")
	if err != nil {
		panic(err)
	}
	return data
}