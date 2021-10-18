package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

//Config Global file to configure
type Config struct {
	File struct {
		Role struct {
			Path    string   `yaml:"path"`
			Exclude []string `yaml:"exclude"`
			Deny    struct {
				Verbs      []string `yaml:"verbs"`
				APIGroups  []string `yaml:"apiGroups"`
				Ressources []string `yaml:"resources"`
			} `yaml:"deny"`
		} `yaml:"role"`
		//RoleBindings struct {
		//	Path    string   `yaml:"path"`
		//	Exclude []string `yaml:"exclude"`
		//	Deny    struct {
		//		Verbs      []string `yaml:"verbs"`
		//		APIGroups  []string `yaml:"apiGroups"`
		//		Ressources []string `yaml:"resources"`
		//	} `yaml:"deny"`
		//} `yaml:"roleBindings"`
		//ClusterRoleBindings struct {
		//	Path    string   `yaml:"path"`
		//	Exclude []string `yaml:"exclude"`
		//	Deny    struct {
		//		Verbs      []string `yaml:"verbs"`
		//		APIGroups  []string `yaml:"apiGroups"`
		//		Ressources []string `yaml:"resources"`
		//	} `yaml:"deny"`
		//} `yaml:"clusterRoleBindings"`
		ClusterRole struct {
			Path    string   `yaml:"path"`
			Exclude []string `yaml:"exclude"`
			Deny    struct {
				Verbs      []string `yaml:"verbs"`
				APIGroups  []string `yaml:"apiGroups"`
				Ressources []string `yaml:"resources"`
			} `yaml:"deny"`
		} `yaml:"clusterRole"`
	} `yaml:"file"`
	Verbosity struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"verbosity"`
}

// ValidateConfigPath check if the config path is a file
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	CheckIfErrExit(err)
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

//NewConfig generate a new configuration from file
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}
	// Open config file
	file, err := os.Open(configPath)
	CheckIfErrExit(err)
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

//GetConfig get configuration file and export it to struct
func GetConfig() (*Config, error) {
	cfgPath := "./config.yml"
	err := ValidateConfigPath("./config.yml")
	CheckIfErrExit(err)
	cfg, err := NewConfig(cfgPath)
	CheckIfErrExit(err)
	return cfg, err
}
