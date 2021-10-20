package main

import "os"

func ExitIfError(err error) {
	if err != nil {
		panic(err)
	}
}

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

//CheckEnvVarLog return the verbosity mode
func CheckEnvVarLog() bool {
	var verbosityMode bool
	if os.Getenv("VERBOSITY") == "DEBUG" {
		verbosityMode = true
	}
	return verbosityMode
}
