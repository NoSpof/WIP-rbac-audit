package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// RoleParser open a json file and export the file on Items array
func RoleParser(path string) Items {
	jsonFile, err := os.Open(path)
	CheckIfErrInfo(err)
	if CheckEnvVarLog() {
		log.Println("Successfully Opened " + path)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var items Items
	json.Unmarshal(byteValue, &items)
	return items
}
