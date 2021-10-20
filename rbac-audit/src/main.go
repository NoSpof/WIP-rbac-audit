package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	config, err := NewConfig("./config.yml")
	ExitIfError(err)
	if config.Verbosity.Format != "json" {
		fmt.Println("Welcom on audit tools")
	}
	kate := connectk8s("./config.yml")
	roles := getRoles(kate)
	//clusterRoles := getClusterRole(kate)
	//rolesBinding := getRolesBinding(kate)
	//clusterRolesBinding := getClusterRoleBinding(kate)
	//switch config.Verbosity.Format {
	//case "table":
	//	RoleTable(roles, config, "roles")
	//	RoleTable(clusterRoles, config, "clusterroles")
	//case "json":
	//	RoleJSON(roles, config)
	//	RoleJSON(clusterRoles, config)
	//default:
	//	RoleTable(roles, config, "roles")
	//	RoleTable(clusterRoles, config, "clusterroles")
	//}
	data := [][]string{}
	var description string
	var result string
	var configExclude []string
	//var denyAPIGroups []string
	var denyVerbs []string
	//var denyResources []string
	configExclude = config.File.Role.Exclude
	//denyAPIGroups = config.File.Role.Deny.APIGroups
	denyVerbs = config.File.Role.Deny.Verbs
	//denyResources = config.File.Role.Deny.Ressources
	fmt.Println(roles.Items)
	for _, item := range roles.Items {
		if InArray(item.Name, configExclude) {
			log.Println(item.Name + "is excluded")
		} else {
			fmt.Println("#######")
			if config.Verbosity.Level == "debug" {
				log.Println("Parsing : " + item.Name + " Roles")
			}
			for _, rule := range item.Rules {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Object", "Details", "Result"})
				for v := 0; v < len(rule.Verbs); v++ {
					if InArray(rule.Verbs[v], denyVerbs) {
						description = rule.Verbs[v]
						result = "Deny"
					} else {
						description = rule.Verbs[v]
						result = "Approved"
					}
					verbsResult := []string{"Verbs", description, result}
					if result != "Approved" || config.Verbosity.Level == "debug" {
						data = append(data, verbsResult)
					}
				}

				for _, line := range data {
					table.Append(line)
				}
				table.Render()
			}
			fmt.Println("#######")

		}

	}
	//fmt.Println(clusterRoles)
	//fmt.Println(rolesBinding)
	//fmt.Println(clusterRolesBinding)
}
