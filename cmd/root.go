package cmd

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Method-Security/pkg/signal"
	"github.com/Method-Security/pkg/writer"
	"github.com/method-security/methodk8s/internal/config"
	"github.com/palantir/pkg/datetime"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type MethodK8s struct {
	version      string
	RootFlags    config.RootFlags
	OutputConfig writer.OutputConfig
	OutputSignal signal.Signal
	K8Config     *rest.Config
	RootCmd      *cobra.Command
}

func NewMethodK8s(version string) *MethodK8s {
	methodK8s := MethodK8s{
		version: version,
		RootFlags: config.RootFlags{
			Quiet:   false,
			Verbose: false,
			ServiceAccountConfig: config.ServiceAccountConfig{
				ServiceAccount: false,
				Token:          "",
				CACert:         "",
				APIServerURL:   "",
			},
			OtherConfig: config.OtherConfig{
				Context: "",
				Path:    "",
				URL:     "",
			},
		},
		OutputConfig: writer.NewOutputConfig(nil, writer.NewFormat(writer.SIGNAL)),
		OutputSignal: signal.NewSignal(nil, datetime.DateTime(time.Now()), nil, 0, nil),
		K8Config:     nil,
	}
	return &methodK8s
}

func (a *MethodK8s) InitRootCommand() {
	var outputFormat string
	var outputFile string
	a.RootCmd = &cobra.Command{
		Use:   "methodk8s",
		Short: "Audit K8 resources",
		Long: `The K8s config is defined in order of:
		1. The '--service-account' flag which creates a config via a token, CA-Cert, and URL set as ENV vars or flags
		2. The '--path' flag which passes the path to a .kube/config file
		3. The $KUBECONFIG var which holds the path to a .kube/config file
		4. The '--url' flag which passes the URL of a potential cluster`,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			var err error

			format, err := validateOutputFormat(outputFormat)
			if err != nil {
				return err
			}
			var outputFilePointer *string
			if outputFile != "" {
				outputFilePointer = &outputFile
			} else {
				outputFilePointer = nil
			}
			a.OutputConfig = writer.NewOutputConfig(outputFilePointer, format)
			cmd.SetContext(svc1log.WithLogger(cmd.Context(), config.InitializeLogging(cmd, &a.RootFlags)))

			k8Config, err := GetK8Config(a)
			if err != nil {
				return err
			}
			a.K8Config = k8Config

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			completedAt := datetime.DateTime(time.Now())
			a.OutputSignal.CompletedAt = &completedAt
			return writer.Write(
				a.OutputSignal.Content,
				a.OutputConfig,
				a.OutputSignal.StartedAt,
				a.OutputSignal.CompletedAt,
				a.OutputSignal.Status,
				a.OutputSignal.ErrorMessage,
			)
		},
	}

	// Standard flags
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Quiet, "quiet", "q", false, "Suppress output")
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Verbose, "verbose", "v", false, "Verbose output")
	a.RootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "Path to output file. If blank, will output to STDOUT")
	a.RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "signal", "Output format (signal, json, yaml). Default value is signal")

	// ServiceAccountConfig flags
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.ServiceAccountConfig.ServiceAccount, "service-account", "s", false, "Set to true if using service account workflow")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.ServiceAccountConfig.Token, "token", "t", "", "Service account token")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.ServiceAccountConfig.CACert, "cert", "a", "", "Base64 encoded ca certificate")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.ServiceAccountConfig.APIServerURL, "server-url", "e", "", "Cluster server url")

	// OtherConfig flags
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.OtherConfig.Context, "context", "c", "", "Cluster context (ie. minikube)")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.OtherConfig.Path, "path", "p", "", "Absolute or relative path to the config file (ie. ~/.kube/config)")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.OtherConfig.URL, "url", "u", "", "Cluster url (ie. mycluster.com)")

	// Flag rules
	a.RootCmd.MarkFlagsMutuallyExclusive("context", "service-account")
	a.RootCmd.MarkFlagsMutuallyExclusive("context", "token")
	a.RootCmd.MarkFlagsMutuallyExclusive("context", "cert")
	a.RootCmd.MarkFlagsMutuallyExclusive("context", "server-url")
	a.RootCmd.MarkFlagsMutuallyExclusive("path", "service-account")
	a.RootCmd.MarkFlagsMutuallyExclusive("path", "token")
	a.RootCmd.MarkFlagsMutuallyExclusive("path", "cert")
	a.RootCmd.MarkFlagsMutuallyExclusive("path", "server-url")
	a.RootCmd.MarkFlagsMutuallyExclusive("url", "service-account")
	a.RootCmd.MarkFlagsMutuallyExclusive("url", "token")
	a.RootCmd.MarkFlagsMutuallyExclusive("url", "cert")
	a.RootCmd.MarkFlagsMutuallyExclusive("url", "server-url")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of methodk8s",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(a.version)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	a.RootCmd.AddCommand(versionCmd)
}

// A utility function to validate that the provided output format is one of the supported formats: json, yaml, signal.
func validateOutputFormat(output string) (writer.Format, error) {
	var format writer.FormatValue
	switch strings.ToLower(output) {
	case "json":
		format = writer.JSON
	case "yaml":
		return writer.Format{}, errors.New("yaml output format is not supported for methodk8s")
	case "signal":
		format = writer.SIGNAL
	default:
		return writer.Format{}, errors.New("invalid output format. Valid formats are: json, yaml, signal")
	}
	return writer.NewFormat(format), nil
}

// GetK8Config gets the k8s config object from the various auth mechanisms
func GetK8Config(a *MethodK8s) (*rest.Config, error) {
	if a.RootFlags.ServiceAccountConfig.ServiceAccount {
		var err error

		// Get Token
		var token string
		if a.RootFlags.ServiceAccountConfig.Token != "" {
			token = a.RootFlags.ServiceAccountConfig.Token
		} else {
			token = os.Getenv("TOKEN")
		}

		// Get CA Certificate
		var caCert []byte
		if a.RootFlags.ServiceAccountConfig.CACert != "" {
			caCert, err = base64.StdEncoding.DecodeString(a.RootFlags.ServiceAccountConfig.CACert)
			if err != nil {
				return nil, err
			}
		} else {
			caCert, err = base64.StdEncoding.DecodeString(os.Getenv("CA_CERT"))
			if err != nil {
				return nil, err
			}
		}

		// Get API Server URL
		var apiServerURL string
		if a.RootFlags.ServiceAccountConfig.APIServerURL != "" {
			apiServerURL = a.RootFlags.ServiceAccountConfig.APIServerURL
		} else {
			apiServerURL = os.Getenv("API_SERVER")
			if err != nil {
				return nil, err
			}
		}

		k8Config, err := MakeConfigFromSecret(token, caCert, apiServerURL)
		if err != nil {
			return nil, err
		}
		return k8Config, nil
	} else if a.RootFlags.OtherConfig.Path != "" {
		k8ConfigPath := a.RootFlags.OtherConfig.Path
		k8Config, err := MakeConfigFromPath(k8ConfigPath, a.RootFlags.OtherConfig.Context)
		if err != nil {
			return nil, err
		}
		return k8Config, nil

	} else if kubeEnv, exists := os.LookupEnv("KUBECONFIG"); exists && kubeEnv != "" {
		k8ConfigPath := os.Getenv("KUBECONFIG")
		k8Config, err := MakeConfigFromPath(k8ConfigPath, a.RootFlags.OtherConfig.Context)
		if err != nil {
			return nil, err
		}
		return k8Config, nil

	} else if a.RootFlags.OtherConfig.URL != "" {
		k8ConfigURL := a.RootFlags.OtherConfig.URL
		k8Config := MakeConfigFromURL(k8ConfigURL)
		return k8Config, nil

	}
	err := errors.New("please provide either: " +
		"Service account creds," +
		"a path to a config file, " +
		"assign $KUBECONFIG to a path to the config file, " +
		"or provide a URL to the cluster")
	return nil, err

}

// MakeConfigFromSecret generates the k8s config object from a secret, ca file (if present), and api server
func MakeConfigFromSecret(token string, caCert []byte, apiServerURL string) (*rest.Config, error) {
	if caCert != nil {
		return &rest.Config{
			Host: apiServerURL,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: caCert,
			},
			BearerToken: token,
		}, nil
	}
	return &rest.Config{
		Host: apiServerURL,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // Disable TLS verification
		},
		BearerToken: token,
	}, nil
}

// MakeConfigFromURL generates the k8s config object from a k8s cluster URL
func MakeConfigFromURL(clusterURL string) *rest.Config {
	return &rest.Config{
		Host: clusterURL,
	}
}

// MakeConfigFromPath generates the k8s config object from a path to a config file
func MakeConfigFromPath(configPath string, context string) (*rest.Config, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath}
	configOverrides := &clientcmd.ConfigOverrides{}

	if context != "" {
		configOverrides.CurrentContext = context
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).ClientConfig()
}
