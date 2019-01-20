package main

import (
	"encoding/json"
	"fmt"
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
func DisplayServerData (data []byte) {
	var serverlist []servers
	JsonBytesToStruct(data, &serverlist) // needs error handling
	fmt.Println("NAME")
	for _, s := range serverlist {
		fmt.Printf("%s\n", s.Name)
	}
	fmt.Printf("Total: %v\n", len(serverlist))
}