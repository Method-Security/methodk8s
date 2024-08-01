package serviceaccount

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func createClusterRoleBinding(clientset *kubernetes.Clientset, namespace, serviceAccountName string) error {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "method-read-only-binding",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "method-read-only-clusterrole",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      serviceAccountName,
				Namespace: namespace,
			},
		},
	}

	_, err := clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Println("ClusterRoleBinding 'method-read-only-binding' already exists, continuing...")
			return nil
		}
		return err
	}
	return nil
}

func createSecret(clientset *kubernetes.Clientset, namespace, serviceAccountName string) error {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "method-service-account-secret",
			Namespace: namespace,
			Annotations: map[string]string{
				"kubernetes.io/service-account.name": serviceAccountName,
			},
		},
		Type: v1.SecretTypeServiceAccountToken,
	}

	_, err := clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Println("Secret 'method-service-account-secret' already exists, continuing...")
			return nil
		}
		return err
	}
	return nil
}

func createServiceAccount(clientset *kubernetes.Clientset, namespace, serviceAccountName string) error {
	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: namespace,
		},
	}

	_, err := clientset.CoreV1().ServiceAccounts(namespace).Create(context.TODO(), serviceAccount, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Println("ServiceAccount 'method-service-account' already exists, continuing...")
			return nil
		}
		return err
	}
	return nil
}

func createClusterRole(clientset *kubernetes.Clientset) error {
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "method-read-only-clusterrole",
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"pods", "services", "endpoints", "nodes"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"apps"},
				Resources: []string{"deployments", "daemonsets", "replicasets", "statefulsets"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"batch"},
				Resources: []string{"jobs", "cronjobs"},
				Verbs:     []string{"get", "list", "watch"},
			},
		},
	}

	_, err := clientset.RbacV1().ClusterRoles().Create(context.TODO(), clusterRole, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Println("ClusterRole 'method-read-only-clusterrole' already exists, continuing...")
			return nil
		}
		return err
	}
	return nil
}

func ApplyServiceAccountConfig(k8config *rest.Config) error {
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	namespace := "default"
	serviceAccountName := "method-service-account"

	if err := createClusterRole(clientset); err != nil {
		return err
	}
	fmt.Println("- Set up Cluster Role (1/4)")

	if err := createServiceAccount(clientset, namespace, serviceAccountName); err != nil {
		return err
	}
	fmt.Println("- Set up Service Account (2/4)")

	if err := createSecret(clientset, namespace, serviceAccountName); err != nil {
		return err
	}
	fmt.Println("- Set up Secret (3/4)")

	if err := createClusterRoleBinding(clientset, namespace, serviceAccountName); err != nil {
		return err
	}
	fmt.Println("- Set up Cluster Role Binding (4/4)")
	fmt.Println("You are all set! Run 'methodk8 service-account creds' to see your tokens")

	return nil
}
