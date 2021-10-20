package main

import (
	"fmt"
	"log"
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
	fmt.Println(roles.Items)
	for _, item := range roles.Items {
		fmt.Println("#######")
		if config.Verbosity.Level == "debug" {
			log.Println("Parsing : " + item.Name + " Roles")
		}
		fmt.Println(item.Rules)
		fmt.Println("#######")
	}
	//fmt.Println(clusterRoles)
	//fmt.Println(rolesBinding)
	//fmt.Println(clusterRolesBinding)
}
