package cmd

import (
	serivceAccount "github.com/method-security/methodk8s/internal/serviceaccount"
	"github.com/spf13/cobra"
)

func (a *MethodK8s) InitServiceAccountCommand() {
	serviceAccountCmd := &cobra.Command{

		Use:   "service-account",
		Short: "Service Account setup and config",
		Long:  `Service Account setup and config`,
	}

	credentialsCmd := &cobra.Command{
		Use:   "creds",
		Short: "Service account credentials",
		Long:  `Use this command to print to console the service account credentials`,
		Run: func(cmd *cobra.Command, args []string) {
			err := serivceAccount.PrintCredentials(a.K8Config)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = nil
		},
	}
	applyCmd := &cobra.Command{
		Use:   "apply",
		Short: "Set up service account in k8s cluster",
		Long:  `Set up service account in k8s cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			err := serivceAccount.ApplyServiceAccountConfig(a.K8Config)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = nil
		},
	}

	serviceAccountCmd.AddCommand(credentialsCmd)
	serviceAccountCmd.AddCommand(applyCmd)
	a.RootCmd.AddCommand(serviceAccountCmd)
}
