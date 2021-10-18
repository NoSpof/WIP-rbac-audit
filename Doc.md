<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# app

```go
import "fnacdarty.com/k8s-rbac-audit"
```

## Index

- [func CheckEnvVarLog() bool](<#func-checkenvvarlog>)
- [func CheckIfErrExit(err error)](<#func-checkiferrexit>)
- [func CheckIfErrInfo(err error)](<#func-checkiferrinfo>)
- [func InArray(str string, list []string) bool](<#func-inarray>)
- [func RoleBindingTable(rolebinding RoleBindings, config *Config)](<#func-rolebindingtable>)
- [func RoleJSON(items Items, config *Config)](<#func-rolejson>)
- [func RoleTable(items Items, config *Config, kind string)](<#func-roletable>)
- [func ValidateConfigPath(path string) error](<#func-validateconfigpath>)
- [type Config](<#type-config>)
  - [func GetConfig() (*Config, error)](<#func-getconfig>)
  - [func NewConfig(configPath string) (*Config, error)](<#func-newconfig>)
- [type Item](<#type-item>)
- [type Items](<#type-items>)
  - [func RoleParser(path string) Items](<#func-roleparser>)
- [type Result](<#type-result>)
- [type Results](<#type-results>)
- [type RoleBinding](<#type-rolebinding>)
- [type RoleBindings](<#type-rolebindings>)
  - [func RoleBindingParser(path string) RoleBindings](<#func-rolebindingparser>)
- [type Rule](<#type-rule>)
- [type RuleReturn](<#type-rulereturn>)
- [type Rules](<#type-rules>)
- [type Subject](<#type-subject>)


## func CheckEnvVarLog

```go
func CheckEnvVarLog() bool
```

CheckEnvVarLog return the verbosity mode

## func CheckIfErrExit

```go
func CheckIfErrExit(err error)
```

CheckIfErrExit return exit 1 if the err is checked\. The programm will stopped

## func CheckIfErrInfo

```go
func CheckIfErrInfo(err error)
```

CheckIfErrInfo return exit 0 if the err is checked

## func InArray

```go
func InArray(str string, list []string) bool
```

InArray : Check if value are in array string Return true if the string are inside array from parameters : string and and array of string\.

## func RoleBindingTable

```go
func RoleBindingTable(rolebinding RoleBindings, config *Config)
```

RoleBindingTable print the rbac rolebinding

## func RoleJSON

```go
func RoleJSON(items Items, config *Config)
```

RoleJSON print the result on json form

## func RoleTable

```go
func RoleTable(items Items, config *Config, kind string)
```

RoleTable export the output to table from role

## func ValidateConfigPath

```go
func ValidateConfigPath(path string) error
```

ValidateConfigPath check if the config path is a file

## type Config

Config Global file to configure

```go
type Config struct {
    File struct {
        Role struct {
            Path    string   `yaml:"path"`
            Exclude []string `yaml:"exclude"`
            Deny    struct {
                Verbs      []string `yaml:"verbs"`
                APIGroups  []string `yaml:"apiGroups"`
                Ressources []string `yaml:"resources"`
            }   `yaml:"deny"`
        }   `yaml:"role"`
        RoleBindings struct {
            Path    string   `yaml:"path"`
            Exclude []string `yaml:"exclude"`
            Deny    struct {
                Verbs      []string `yaml:"verbs"`
                APIGroups  []string `yaml:"apiGroups"`
                Ressources []string `yaml:"resources"`
            }   `yaml:"deny"`
        }   `yaml:"roleBindings"`
        ClusterRoleBindings struct {
            Path    string   `yaml:"path"`
            Exclude []string `yaml:"exclude"`
            Deny    struct {
                Verbs      []string `yaml:"verbs"`
                APIGroups  []string `yaml:"apiGroups"`
                Ressources []string `yaml:"resources"`
            }   `yaml:"deny"`
        }   `yaml:"clusterRoleBindings"`
        ClusterRole struct {
            Path    string   `yaml:"path"`
            Exclude []string `yaml:"exclude"`
            Deny    struct {
                Verbs      []string `yaml:"verbs"`
                APIGroups  []string `yaml:"apiGroups"`
                Ressources []string `yaml:"resources"`
            }   `yaml:"deny"`
        }   `yaml:"clusterRole"`
    }   `yaml:"file"`
    Verbosity struct {
        Level  string `yaml:"level"`
        Format string `yaml:"format"`
    }   `yaml:"verbosity"`
}
```

### func GetConfig

```go
func GetConfig() (*Config, error)
```

GetConfig get configuration file and export it to struct

### func NewConfig

```go
func NewConfig(configPath string) (*Config, error)
```

NewConfig generate a new configuration from file

## type Item

Item struct which contains a name a type and a list of social links

```go
type Item struct {
    APIVersion string `json:"apiVersion"`
    Kind       string `json:"kind"`
    Metadata   struct {
        Name      string `json:"name"`
        Namespace string `json:"namespace"`
    }   `json:"metadata"`
    Rules []Rule `json:"rules"`
}
```

## type Items

Items struct which contains an array of users

```go
type Items struct {
    Items []Item `json:"items"`
}
```

### func RoleParser

```go
func RoleParser(path string) Items
```

RoleParser open a json file and export the file on Items array

## type Result

Result item to array

```go
type Result struct {
    Name      string       `json:"name"`
    Namespace string       `json:"namespace"`
    Rules     []RuleReturn `json:"rules"`
}
```

## type Results

Results array from result

```go
type Results struct {
    Results []Result `json:"results"`
}
```

## type RoleBinding

RoleBinding struct which contain One rolebinding

```go
type RoleBinding struct {
    APIVersion string `json:"apiVersion"`
    Kind       string `json:"kind"`
    Metadata   struct {
        Name      string `json:"name"`
        Namespace string `json:"namespace"`
    }   `json:"metadata"`
    RoleRef struct {
        APIGroups string `json:"apiGroups"`
        Kind      string `json:"kind"`
        Name      string `json:"name"`
    }   `json:"roleRef"`
    Subjects []Subject
}
```

## type RoleBindings

RoleBindings struct which contains an array of RoleBinding

```go
type RoleBindings struct {
    RoleBindings []RoleBinding `json:"items"`
}
```

### func RoleBindingParser

```go
func RoleBindingParser(path string) RoleBindings
```

RoleBindingParser open a json file and export it on array

## type Rule

Rule item for Array

```go
type Rule struct {
    APIGroups  []string `json:"apiGroups"`
    Ressources []string `json:"resources"`
    Verbs      []string `json:"verbs"`
}
```

## type RuleReturn

RuleReturn the rule return for output

```go
type RuleReturn struct {
    ID        int    `json:"id"`
    Objectype string `json:"object"`
    Detail    string `json:"detail"`
    Status    string `json:"status"`
}
```

## type Rules

Rules array of rule

```go
type Rules struct {
    Rules []Rule `json:"rules"`
}
```

## type Subject

Subject One user or Sa on kubernetes

```go
type Subject struct {
    APIGroups string `json:"apiGroups"`
    Kind      string `json:"kind"`
    Name      string `json:"name"`
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)