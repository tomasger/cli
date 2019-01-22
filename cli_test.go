package main

import (
	"github.com/jessevdk/go-flags"
	"testing"
)

func TestLoginWithGoodData(t *testing.T) {
	credFileExisted := backupStorageFile(name.credentials)
	l := LoginCommand{"someaccount", "somepassword"}
	err := l.Execute(nil)
	if credFileExisted {
		restoreStorageFile(name.credentials)
	}
	if err != nil {
		t.Errorf("Login command with valid parameters failed. Error: %v", err.Error())
	}
}
func TestLoginWithInvalidData(t *testing.T) {
	credFileExisted := backupStorageFile(name.credentials)
	tooLongUsername := "i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt" +
		"i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt" +
		"i5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjti5iJcdyVU6UsI0DYbfjt"
	l := LoginCommand{tooLongUsername, "somepassword"}
	err := l.Execute(nil)
	if credFileExisted {
		restoreStorageFile(name.credentials)
	}
	if e, ok := err.(*flags.Error); !(ok && e.Type == flags.ErrInvalidChoice) {
		t.Errorf("Login command with invalid parameters failed. Expected: %v. Got: %v", flags.ErrInvalidChoice, err.Error())
	}
}