package main

import (
	"encoding/json"
	"fmt"
)

// RoleJSON print the result on json form
func RoleJSON(items Items, config *Config) {
	var ret = []Result{}
	var description string
	var status string
	var kind string
	for i := 0; i < len(items.Items); i++ {
		if InArray(items.Items[i].Metadata.Name, config.File.Role.Exclude) {
			if CheckEnvVarLog() {
				fmt.Println("Value : " + items.Items[i].Metadata.Name + " excluded")
			}
		} else {
			arrayResultReturns := []RuleReturn{}
			for j := 0; j < len(items.Items[i].Rules); j++ {

				//For ApiGroups
				for k := 0; k < len(items.Items[i].Rules[j].APIGroups); k++ {
					if InArray(items.Items[i].Rules[j].APIGroups[k], config.File.Role.Deny.APIGroups) {
						description = "ApiGroups  Deny"
						status = "Deny"
						kind = "ApiGroups"
					} else {
						description = ""
						status = "Approved"
						kind = "ApiGroups"
					}
					resultRules := RuleReturn{j, kind, description, status}
					if status != "Approved" || config.Verbosity.Level == "debug" {
						arrayResultReturns = append(arrayResultReturns, resultRules)
					}

				} //End For ApiGroups

				//For Verbs
				for l := 0; l < len(items.Items[i].Rules[j].Verbs); l++ {
					if InArray(items.Items[i].Rules[j].Verbs[l], config.File.Role.Deny.Verbs) || items.Items[i].Rules[j].Verbs[l] == "" || items.Items[i].Rules[j].Verbs[l] == "*" {
						description = items.Items[i].Rules[j].Verbs[l]
						status = "Deny"
						kind = "Verbs"
					} else {
						description = items.Items[i].Rules[j].Verbs[l]
						status = "Approved"
						kind = "Verbs"
					}
					resultRules := RuleReturn{j, kind, description, status}
					if status != "Approved" || config.Verbosity.Level == "debug" {
						arrayResultReturns = append(arrayResultReturns, resultRules)
					}
				}
				//End For Verbs
				//For Resources
				for p := 0; p < len(items.Items[i].Rules[j].Ressources); p++ {
					if InArray(items.Items[i].Rules[j].Ressources[p], config.File.Role.Deny.Ressources) || items.Items[i].Rules[j].Ressources[p] == "" || items.Items[i].Rules[j].Ressources[p] == "*" {
						description = items.Items[i].Rules[j].Ressources[p]
						status = "Deny"
						kind = "Resources"
					} else {
						description = items.Items[i].Rules[j].Ressources[p]
						status = "Approved"
						kind = "Resources"
					}
					resultRules := RuleReturn{j, kind, description, status}
					if status != "Approved" || config.Verbosity.Level == "debug" {
						arrayResultReturns = append(arrayResultReturns, resultRules)
					}
				}
				//End For Resources

			}
			result := Result{items.Items[i].Metadata.Name, items.Items[i].Metadata.Namespace, arrayResultReturns}
			ret = append(ret, result)
		}
	}
	data, err := json.Marshal(ret)
	CheckIfErrExit(err)
	fmt.Printf("%s\n", data)
}
