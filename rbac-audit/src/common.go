package main

import (
	"log"
	"os"
)

// InArray : Check if value are in array string
// Return true if the string are inside array from parameters : string and and array of string.
func InArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// CheckIfErrInfo return exit 0 if the err is checked
func CheckIfErrInfo(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(0)
	}
}

// CheckIfErrExit return exit 1 if the err is checked. The programm will stopped
func CheckIfErrExit(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}

//CheckEnvVarLog return the verbosity mode
func CheckEnvVarLog() bool {
	var verbosityMode bool
	if os.Getenv("VERBOSITY") == "DEBUG" {
		verbosityMode = true
	}
	return verbosityMode
}
