package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)
var realStoragePath string = storagePath
var realFileNames FileNames = fileNames

func SetTestStoragePath() {
	storagePath = "test_data/"
}
func SetTestCredentialsName(newName string) {
	fileNames.credentials = newName
}
func SetTestServersName(newName string) {
	fileNames.servers = newName
}
func RevertTestStorageData() {
	storagePath = realStoragePath
	fileNames = realFileNames
}
func BackupStorageFile(fileName string) bool{
	if _, err := os.Stat(storagePath + fileNames.credentials); os.IsNotExist(err) {
		return false // file did not exist
	}
	os.Rename(storagePath+fileName, storagePath+fileName+"_testbackup")
	return true

}
func RestoreStorageFile(fileName string) {
	os.Remove(storagePath+fileName)
	os.Rename(storagePath+fileName+"_testbackup", storagePath+fileName)
}
func TestSaveLoginData(t *testing.T) {
	SetTestStoragePath()
	SetTestCredentialsName("cred_saveLoginData")
	username, password := "test", "something"
	SaveLoginData(username, password)
	dat, err := ioutil.ReadFile(storagePath + fileNames.credentials)
	if err != nil {
		t.Errorf("Save login data failed. Expected: file %v%v to be created. Got: %v", storagePath, fileNames, err)
	}
	RevertTestStorageData()
	expectedData := []byte(username+"\n"+password)
	if !bytes.Equal(dat,expectedData) {
		t.Errorf("Save login data failed. Expected file contents: %s. Got: %s", expectedData, dat)
	}
}

func TestSaveServersData(t *testing.T) {
	SetTestStoragePath()
	SetTestServersName("servers_saveServersData")
	expectedList := []byte("[{\"name\":\"United States #93\",\"distance\":1634},{\"name\":\"Germany #81\",\"distance\":26},{\"name\":\"Latvia #7\",\"distance\":1581}]")
	SaveServerData(expectedList)
	dat, err := ioutil.ReadFile(storagePath + fileNames.servers)
	if err != nil {
		t.Errorf("Save server data failed. Expected: file %v%v to be created. Got: %v", storagePath, fileNames, err)
	}
	RevertTestStorageData()
	if !bytes.Equal(dat,expectedList) {
		t.Errorf("Save login data failed. Expected file contents: %s. Got: %s", expectedList, dat)
	}
}
func TestLoadLoginDataValid(t *testing.T) {
	SetTestStoragePath()
	SetTestCredentialsName("cred_valid")
	uname, pass, err := LoadLoginData()
	expected_uname := "validusername1"
	expected_pass := "validPassword!"
	RevertTestStorageData()
	if err != nil || uname != expected_uname || pass != expected_pass {
		t.Errorf("Test Load login data failed. Expected: %s, %s; error: %v. Got: %s, %s; error: %v",
			expected_uname, expected_pass, nil, uname, pass, err)
	}
}
func TestLoadLoginDataInvalid(t *testing.T) {
	SetTestStoragePath()
	SetTestCredentialsName("cred_invalid")
	_, _, err := LoadLoginData()
	RevertTestStorageData()
	if err, ok := err.(CliError); !(ok && err.errorType == ErrParse) {
		t.Errorf("Test Load login data failed. Expected: ErrParse. Got: %v", err)
	}
}
func TestLoadServerDataValid(t *testing.T) {
	SetTestStoragePath()
	SetTestServersName("servers_valid")
	data, err := LoadServerData()
	expectedData := []byte("[{\"name\":\"United States #93\",\"distance\":1634},{\"name\":\"Germany #81\",\"distance\":26},{\"name\":\"Latvia #7\",\"distance\":1581}]\n")
	RevertTestStorageData()
	if err != nil || !bytes.Equal(expectedData,data) {
		t.Errorf("Test Load server data failed. Expected: %s; error: %v. Got: %s; error: %v",
			expectedData, nil, data, err)
	}
}
func TestLoadServerDataInvalid(t *testing.T) {
	SetTestStoragePath()
	SetTestServersName("servers_invalid")
	data, err := LoadServerData()
	RevertTestStorageData()
	_, err = ParseServerData(data)
	if err, ok := err.(CliError); !(ok && err.errorType == ErrParse) {
		t.Errorf("Test Load login data failed. Expected: ErrParse. Got: %v", err)
	}
}