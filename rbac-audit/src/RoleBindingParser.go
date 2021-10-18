package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// RoleBindingParser open a json file and export it on array
func RoleBindingParser(path string) RoleBindings {
	jsonFile, err := os.Open(path)
	CheckIfErrInfo(err)
	if CheckEnvVarLog() {
		log.Println("Successfully Opened " + path)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var rolebindings RoleBindings
	json.Unmarshal(byteValue, &rolebindings)
	return rolebindings
}
