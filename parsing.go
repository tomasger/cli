package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

func JsonBytesToStruct(data []byte, container interface{}) error {
	err := json.Unmarshal(data, &container)
	if err != nil {
		msg := ""
		switch t := container.(type) {
		case *[]servers:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed", t)
		case *map[string]string:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed", t)
		default:
			msg = fmt.Sprintf("JSON Bytes to %T conversion is not implemented", t)
		}
		return ErrJSON.Wrap(err, msg)
	}
	return nil
}

func ParseLoginData(data []byte) ([]string, error) {
	splitFn := func(c rune) bool {
		return c == '\n'
	}
	credentials := strings.FieldsFunc(string(data), splitFn)
	if len(credentials) != 2 {
		logrus.Error("Credentials file contains invalid information.")
		return nil, ErrParse.New("Credentials file contains invalid information")
	}
	return credentials, nil
}
func ParseServerData(data []byte) ([]servers, error) {
	var serverlist []servers
	err := JsonBytesToStruct(data, &serverlist)
	if err != nil {
		logrus.Error("Servers file contains invalid information")
		return nil, ErrParse.Wrap(err, "Servers data contains invalid information")
	}
	return serverlist, nil
}
