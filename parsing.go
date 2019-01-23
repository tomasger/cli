package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func JsonBytesToStruct(data []byte, container interface{}) error {
	err := json.Unmarshal(data, &container)
	if err != nil {
		msg := ""
		switch t := container.(type) {
		case *[]servers:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed" , t)
		case *map[string]string:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed" , t)
		default:
			msg = fmt.Sprintf("JSON Bytes to %T conversion is not implemented" , t)
		}
		return ErrJSON.Wrap(err, msg)
	}
	return nil
}
func DisplayServerData (data []byte, sorting string) error {
	serverlist , err := ParseServerData(data)
	if err != nil {
		return ErrParse.Wrap(err, "Displaying server information failed")
	}
	if sorting != "" {
		Sort(serverlist, sorting)
	}
	fmt.Printf("NAME            DISTANCE\n")
	for _, server := range serverlist {
		fmt.Printf("%-20s%4d\n", server.Name, server.Distance)
	}
	fmt.Printf("Total number of servers: %v\n", len(serverlist))
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
func ParseServerData (data []byte) ([]servers, error) {
	var serverlist []servers
	err := JsonBytesToStruct(data, &serverlist)
	if err != nil {
		logrus.Error("Servers file contains invalid information")
		return nil, ErrParse.Wrap(err, "Servers data contains invalid information")
	}
	return serverlist, nil
}
func Sort(serverList []servers, sorting string) {
	logrus.Debug("Data sorting was called. Argument: ", sorting)
	switch sorting {
	case "best":
		sort.Slice(serverList, func(i, j int) bool {
			return serverList[i].Distance < serverList[j].Distance
		})
	case "alphabetical":
		sort.Slice(serverList, func(i, j int) bool {
			return serverList[i].Name < serverList[j].Name
		})
	default:
		logrus.Warning("Unexpected sorting parameter received. List will not be sorted")
		}
}