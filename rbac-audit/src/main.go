package main

import (
	"fmt"
)

func main() {
	config, err := NewConfig("./config.yml")
	CheckIfErrExit(err)
	if config.Verbosity.Format != "json" {
		fmt.Println("Welcom on audit tools")
	}
	roles := RoleParser(config.File.Role.Path)
	clusterroles := RoleParser(config.File.ClusterRole.Path)
	//clusterrolebinding := RoleBindingParser(config.File.ClusterRoleBindings.Path)
	switch config.Verbosity.Format {
	case "table":
		RoleTable(roles, config, "roles")
		RoleTable(clusterroles, config, "clusterroles")
	case "json":
		RoleJSON(roles, config)
		RoleJSON(clusterroles, config)
	default:
		RoleTable(roles, config, "roles")
		RoleTable(clusterroles, config, "clusterroles")
	}
}
