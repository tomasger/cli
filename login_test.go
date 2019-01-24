package main

import (
	"github.com/jessevdk/go-flags"
	"testing"
)

func TestLoginWithGoodData(t *testing.T) {
	SetTestStoragePath()
	SetTestCredentialsName("cred_loginCommandValid")
	l := LoginCommand{"someaccount", "somepassword"}
	err := l.Execute(nil)
	RevertTestStorageData()
	if err != nil {
		t.Errorf("Login command with valid parameters failed. Error: %v", err.Error())
	}
}
func TestLoginWithInvalidData(t *testing.T) {
	SetTestStoragePath()
	SetTestCredentialsName("cred_loginCommandInvalid")
	tooLongUsername := "i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt" +
		"i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt" +
		"i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt"
	l := LoginCommand{tooLongUsername, "somepassword"}
	err := l.Execute(nil)
	RevertTestStorageData()
	if e, ok := err.(*flags.Error); !(ok && e.Type == flags.ErrInvalidChoice) {
		t.Errorf("Login command with invalid parameters failed. Expected: %v. Got: %v", flags.ErrInvalidChoice, err.Error())
	}
}