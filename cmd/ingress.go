package cmd

import (
	"github.com/method-security/methodk8s/internal/ingress"
	"github.com/spf13/cobra"
)

func (a *MethodK8s) InitIngressCommand() {
	ingressCmd := &cobra.Command{
		Use:   "ingress",
		Short: "Audit and command Ingresses",
		Long:  `Audit and command Ingresses`,
	}

	enumerateCmd := &cobra.Command{
		Use:   "enumerate",
		Short: "Enumerate Ingresses",
		Long:  `Enumerate Ingresses`,
		Run: func(cmd *cobra.Command, args []string) {

			onlyGateway, err := cmd.Flags().GetBool("gateways")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			report, err := ingress.EnumerateIngresses(a.K8Config, onlyGateway)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	enumerateCmd.Flags().Bool("gateways", false, "Only include gateway objects")

	ingressCmd.AddCommand(enumerateCmd)
	a.RootCmd.AddCommand(ingressCmd)
}
