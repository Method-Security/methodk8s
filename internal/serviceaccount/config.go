package serviceaccount

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"
)

func createClusterRoleBinding(clientset *kubernetes.Clientset, namespace, serviceAccountName string, apply bool) (*rbacv1.ClusterRoleBinding, error) {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRoleBinding",
		},
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

	ctx := context.Background()
	if apply {
		_, err := clientset.RbacV1().ClusterRoleBindings().Create(ctx, clusterRoleBinding, metav1.CreateOptions{})
		if err != nil {
			if errors.IsAlreadyExists(err) {
				fmt.Println("ClusterRoleBinding 'method-read-only-binding' already exists, continuing...")
				return clusterRoleBinding, nil
			}
			return nil, err
		}
	}
	return clusterRoleBinding, nil
}

func createSecret(clientset *kubernetes.Clientset, namespace, serviceAccountName string, apply bool) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "method-service-account-secret",
			Namespace: namespace,
			Annotations: map[string]string{
				"kubernetes.io/service-account.name": serviceAccountName,
			},
		},
		Type: corev1.SecretTypeServiceAccountToken,
	}

	ctx := context.Background()
	if apply {
		_, err := clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
		if err != nil {
			if errors.IsAlreadyExists(err) {
				fmt.Println("Secret 'method-service-account-secret' already exists, continuing...")
				return secret, nil
			}
			return nil, err
		}
	}
	return secret, nil
}

func createServiceAccount(clientset *kubernetes.Clientset, namespace, serviceAccountName string, apply bool) (*corev1.ServiceAccount, error) {
	serviceAccount := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceAccountName,
			Namespace: namespace,
		},
	}

	ctx := context.Background()
	if apply {
		_, err := clientset.CoreV1().ServiceAccounts(namespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
		if err != nil {
			if errors.IsAlreadyExists(err) {
				fmt.Println("ServiceAccount 'method-service-account' already exists, continuing...")
				return serviceAccount, nil
			}
			return nil, err
		}
	}
	return serviceAccount, nil
}

func createClusterRole(clientset *kubernetes.Clientset, apply bool) (*rbacv1.ClusterRole, error) {
	clusterRole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
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

	ctx := context.Background()
	if apply {
		_, err := clientset.RbacV1().ClusterRoles().Create(ctx, clusterRole, metav1.CreateOptions{})
		if err != nil {
			if errors.IsAlreadyExists(err) {
				fmt.Println("ClusterRole 'method-read-only-clusterrole' already exists, continuing...")
				return clusterRole, nil
			}
			return nil, err
		}
	}
	return clusterRole, nil
}

func Config(ctx context.Context, k8config *rest.Config, apply bool, namespace string) error {
	config := k8config

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	serviceAccountName := "method-service-account"

	clusterRole, err := createClusterRole(clientset, apply)
	if err != nil {
		return err
	}
	if apply {
		fmt.Println("- Cluster Role configured (1/4)")
	} else {
		fmt.Println("---")
		prettyPrintYAML(clusterRole)
	}

	serviceAccount, err := createServiceAccount(clientset, namespace, serviceAccountName, apply)
	if err != nil {
		return err
	}
	if apply {
		fmt.Println("- Service Account configured (2/4)")
	} else {
		fmt.Println("---")
		prettyPrintYAML(serviceAccount)
	}

	secret, err := createSecret(clientset, namespace, serviceAccountName, apply)
	if err != nil {
		return err
	}
	if apply {
		fmt.Println("- Secret configured (3/4)")
	} else {
		fmt.Println("---")
		prettyPrintYAML(secret)
	}

	clusterRoleBinding, err := createClusterRoleBinding(clientset, namespace, serviceAccountName, apply)
	if err != nil {
		return err
	}
	if apply {
		fmt.Println("- Cluster Role Binding configured(4/4)")
		fmt.Println("You are all set! Run 'methodk8 service-account creds' to see your tokens")
	} else {
		fmt.Println("---")
		prettyPrintYAML(clusterRoleBinding)
	}

	return nil
}

// prettyPrintYAML prints a Kubernetes object as pretty YAML to the console
func prettyPrintYAML(obj interface{}) {
	// Set creationTimestamp to nil
	switch o := obj.(type) {
	case *rbacv1.ClusterRole:
		o.ObjectMeta.CreationTimestamp = metav1.Time{}
	case *rbacv1.ClusterRoleBinding:
		o.ObjectMeta.CreationTimestamp = metav1.Time{}
	case *corev1.Secret:
		o.ObjectMeta.CreationTimestamp = metav1.Time{}
	case *corev1.ServiceAccount:
		o.ObjectMeta.CreationTimestamp = metav1.Time{}
	}

	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Printf("Error marshalling to YAML: %v\n", err)
		return
	}
	fmt.Printf("%s\n", string(yamlData))
}
