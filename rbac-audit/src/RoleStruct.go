package main

// Results array from result
type Results struct {
	Results []Result `json:"results"`
}

// Result item to array
type Result struct {
	Name      string       `json:"name"`
	Namespace string       `json:"namespace"`
	Rules     []RuleReturn `json:"rules"`
}

// RuleReturn the rule return for output
type RuleReturn struct {
	ID        int    `json:"id"`
	Objectype string `json:"object"`
	Detail    string `json:"detail"`
	Status    string `json:"status"`
}

// Items struct which contains
// an array of users
type Items struct {
	Items []Item `json:"items"`
}

//Rules array of rule
type Rules struct {
	Rules []Rule `json:"rules"`
}

// Item struct which contains a name
// a type and a list of social links
type Item struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Rules []Rule `json:"rules"`
}

//Rule item for Array
type Rule struct {
	APIGroups  []string `json:"apiGroups"`
	Ressources []string `json:"resources"`
	Verbs      []string `json:"verbs"`
}
