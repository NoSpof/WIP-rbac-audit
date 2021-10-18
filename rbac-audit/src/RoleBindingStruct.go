package main

//RoleBindings struct which contains an array of RoleBinding
type RoleBindings struct {
	RoleBindings []RoleBinding `json:"items"`
}

//RoleBinding struct which contain One rolebinding
type RoleBinding struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	RoleRef struct {
		APIGroups string `json:"apiGroups"`
		Kind      string `json:"kind"`
		Name      string `json:"name"`
	} `json:"roleRef"`
	Subjects []Subject
}

//Subject One user or Sa on kubernetes
type Subject struct {
	APIGroups string `json:"apiGroups"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
}
