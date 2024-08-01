package cmd

import (
	"context"
	"fmt"

	serivceAccount "github.com/method-security/methodk8s/internal/serviceaccount"
	"github.com/spf13/cobra"
)

func (a *MethodK8s) InitServiceAccountCommand() {
	serviceAccountCmd := &cobra.Command{

		Use:   "service-account",
		Short: "Service Account setup and config",
		Long:  `Service Account setup and config`,
	}

	var namespace string
	var apply bool
	credentialsCmd := &cobra.Command{
		Use:   "creds",
		Short: "Service account credentials",
		Long:  `Use this command to print to console the service account credentials`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := serivceAccount.PrintCredentials(ctx, a.K8Config, namespace)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	configureCmd := &cobra.Command{
		Use:   "config",
		Short: "Set up service account in k8s cluster",
		Long:  `Set up service account in k8s cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := serivceAccount.Config(ctx, a.K8Config, apply, namespace)
			if err != nil {
				fmt.Println(err)
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	credentialsCmd.Flags().StringVar(&namespace, "namespace", "default", "Set the namespace for the Service Account and Secret")
	configureCmd.Flags().BoolVar(&apply, "apply", false, "Apply the Service Account yamls")
	configureCmd.Flags().StringVar(&namespace, "namespace", "default", "Set the namespace for the Service Account and Secret")

	serviceAccountCmd.AddCommand(credentialsCmd)
	serviceAccountCmd.AddCommand(configureCmd)
	a.RootCmd.AddCommand(serviceAccountCmd)
}
