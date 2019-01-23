package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func JsonBytesToStruct(data []byte, s interface{}) error {
	dec_error := json.Unmarshal(data, &s)
	if dec_error != nil {
		msg := ""
		switch t := s.(type) {
		case *[]servers:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed" , t)
		case *map[string]string:
			msg = fmt.Sprintf("JSON Bytes to %T conversion failed" , t)
		default:
			msg = fmt.Sprintf("JSON Bytes to %T conversion is not implemented" , t)
		}
		return ErrJSON.Wrap(dec_error, msg)
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
	for _, s := range serverlist {
		fmt.Printf("%-20s%4d\n", s.Name, s.Distance)
	}
	fmt.Printf("Total: %v\n", len(serverlist))
	return nil
}
func ParseLoginData(data []byte) ([]string, error) {
	//test if encrypted
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
func Sort(s []servers, sorting string) {
	logrus.Debug("Data sorting was called. Argument: ", sorting)
	switch sorting {
	case "best":
		sort.Slice(s, func(i, j int) bool {
			return s[i].Distance < s[j].Distance
		})
	case "alphabetical":
		sort.Slice(s, func(i, j int) bool {
			return s[i].Name < s[j].Name
		})
	default:
		logrus.Warning("Unexpected sorting parameter received. List will not be sorted")
		}
}