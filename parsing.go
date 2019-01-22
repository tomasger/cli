package main

import (
	"encoding/json"
	"fmt"
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
func DisplayServerData (data []byte) error {
	serverlist , err := ParseServerData(data)
	if err != nil {
		return ErrParse.Wrap(err, "Displaying server information failed")
	}
	fmt.Println("NAME")
	for _, s := range serverlist {
		fmt.Printf("%s\n", s.Name)
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
		return nil, ErrParse.New("Credentials file contained invalid information")
	}
	return credentials, nil
}
func ParseServerData (data []byte) ([]servers, error) {
	var serverlist []servers
	err := JsonBytesToStruct(data, &serverlist)
	if err != nil {
		return nil, ErrParse.Wrap(err, "Servers data contained invalid information")
	}
	return serverlist, nil
}