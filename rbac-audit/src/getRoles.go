package main

import (
	"context"

	"k8s.io/api/rbac/v1beta1"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getRoles(kate *kubernetes.Clientset) *v1beta1.RoleList {
	roles, err := kate.RbacV1beta1().Roles("").List(context.TODO(), metav1.ListOptions{})
	ExitIfError(err)
	return roles
}

func getClusterRole(kate *kubernetes.Clientset) *v1beta1.ClusterRoleList {
	clusterRoles, err := kate.RbacV1beta1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	ExitIfError(err)
	return clusterRoles
}

func getRolesBinding(kate *kubernetes.Clientset) *v1beta1.RoleBindingList {
	rolesBinding, err := kate.RbacV1beta1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
	ExitIfError(err)
	return rolesBinding
}

func getClusterRoleBinding(kate *kubernetes.Clientset) *v1beta1.ClusterRoleBindingList {
	clusterRolesBinding, err := kate.RbacV1beta1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	ExitIfError(err)
	return clusterRolesBinding
}
