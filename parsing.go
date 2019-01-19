package main

import (
	"encoding/json"
	"fmt"
)

func JsonBytesToStruct(data []byte, s interface{}) error {
	switch t := s.(type) {
	case *[]servers:
		dec_error := json.Unmarshal(data, &s)
		if dec_error != nil {
			fmt.Printf("FATAL: JSON Bytes to %T conversion failed" , t)
			return dec_error
		}
	case *map[string]string:
		dec_error := json.Unmarshal(data, &s)
		if dec_error != nil {
			fmt.Printf("FATAL: JSON Bytes to %T conversion failed" , t)
		}
	default:
		fmt.Printf("FATAL: type %T is not implemented for conversion", t)
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