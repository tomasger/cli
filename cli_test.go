package main

import (
	"testing"
)
func TestArgParserWithLogin(t *testing.T) {
	args :=[]string {"path", "login", "--username", "tesonet", "--password", "partyanimal"}
	result := argParser(args)
	if result != Login {
		t.Error()
	}
}
func TestArgParserWithServers(t *testing.T) {
	args :=[]string {"path", "servers", "--local"}
	result := argParser(args)
	if result != Servers {
		t.Error()
	}
}
func TestArgParserWithInvalidCommand(t *testing.T) {
	args :=[]string {"path", "reconnect"}
	result := argParser(args)
	if result != Help {
		t.Error()
	}
}
func TestLoginWithValidInput(t *testing.T) {
	args :=[]string {"path", "login", "--username", "tesonet", "--password", "partyanimal"}
	result := argParser(args)
	if result != Login {
		t.Error()
	}
}
//func TestLoginWithInvalidUsername(t *testing.T) {
//	args :=[]string {"path", "login", "--username", "", "--password", "partyanimal"}
//	result := argParser(args)
//	fmt.Println("Result is: ", result)
//	if result == Login {
//		t.Error()
//	}
//}
func TestServersWithoutArguments(t *testing.T) {
	args :=[]string {"path", "servers"}
	result := parseServerParameters(args)
	if result {
		t.Error()
	}
}
func TestServersWithLocal(t *testing.T) {
	args :=[]string {"path", "servers", "--local"}
	result := parseServerParameters(args)
	if !result {
		t.Error()
	}
}