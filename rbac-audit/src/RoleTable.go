package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// RoleTable export the output to table from role
func RoleTable(items Items, config *Config, kind string) {
	var configExclude []string
	var denyAPIGroups []string
	var denyVerbs []string
	var denyResources []string
	data := [][]string{}
	var description string
	var result string
	switch kind {
	case "roles":
		configExclude = config.File.Role.Exclude
		denyAPIGroups = config.File.Role.Deny.APIGroups
		denyVerbs = config.File.Role.Deny.Verbs
		denyResources = config.File.Role.Deny.Ressources
	case "clusterroles":
		configExclude = config.File.ClusterRole.Exclude
		denyAPIGroups = config.File.ClusterRole.Deny.APIGroups
		denyVerbs = config.File.ClusterRole.Deny.Verbs
		denyResources = config.File.ClusterRole.Deny.Ressources
	default:
		configExclude = config.File.Role.Exclude
		denyAPIGroups = config.File.Role.Deny.APIGroups
		denyVerbs = config.File.Role.Deny.Verbs
		denyResources = config.File.Role.Deny.Ressources
	}

	for i := 0; i < len(items.Items); i++ {
		if InArray(items.Items[i].Metadata.Name, configExclude) {
			if CheckEnvVarLog() {
				fmt.Println("Value : " + items.Items[i].Metadata.Name + " excluded")
			}
		} else {
			fmt.Println("Audit " + items.Items[i].Metadata.Name + " : ")
			for j := 0; j < len(items.Items[i].Rules); j++ {
				fmt.Println("Rules #" + strconv.Itoa(j))
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Object", "Details", "Result"})
				// Check APiGroups
				for k := 0; k < len(items.Items[i].Rules[j].APIGroups); k++ {
					if InArray(items.Items[i].Rules[j].APIGroups[k], denyAPIGroups) {
						description = "ApiGroups  Deny"
						result = "Deny"
					} else {
						description = ""
						result = "Approved"
					}
					apigroupsResult := []string{"ApiGroups", description, result}
					if result != "Approved" || config.Verbosity.Level == "debug" {
						data = append(data, apigroupsResult)
					}
				}
				//Check Verbs
				for l := 0; l < len(items.Items[i].Rules[j].Verbs); l++ {
					if InArray(items.Items[i].Rules[j].Verbs[l], denyVerbs) || items.Items[i].Rules[j].Verbs[l] == "" || items.Items[i].Rules[j].Verbs[l] == "*" {
						description = items.Items[i].Rules[j].Verbs[l]
						result = "Deny"
					} else {
						description = items.Items[i].Rules[j].Verbs[l]
						result = "Approved"
					}
					verbsResult := []string{"Verbs", description, result}
					if result != "Approved" || config.Verbosity.Level == "debug" {
						data = append(data, verbsResult)
					}
				}
				// Check ressources
				for p := 0; p < len(items.Items[i].Rules[j].Ressources); p++ {
					if InArray(items.Items[i].Rules[j].Ressources[p], denyResources) || items.Items[i].Rules[j].Ressources[p] == "" || items.Items[i].Rules[j].Ressources[p] == "*" {
						description = items.Items[i].Rules[j].Ressources[p]
						result = "Deny"
					} else {
						description = items.Items[i].Rules[j].Ressources[p]
						result = "Approved"
					}
					ressourcesResult := []string{"Ressources", description, result}
					if result != "Approved" || config.Verbosity.Level == "debug" {
						data = append(data, ressourcesResult)
					}
				}
				for _, v := range data {
					table.Append(v)
				}
				table.Render()
			}

		}

	}

}
