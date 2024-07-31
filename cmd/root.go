package cmd

import (
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
			Context: "",
			Path:    "",
			Url:     "",
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
		Long: `The K8s context is defined in order of:
		1. The '--path' flag representing the path to a .kube/config file
		2. $KUBECONFIG which holds the path to a .kube/config file
		3. The '--url' flag which holds the URL of a potential cluster
		The '--context' flag can also be used to specfiy the context working with a .kube/config file
		If nothing is provided an error will be thrown`,
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

			context := a.RootFlags.Context

			var k8Config *rest.Config = nil
			if a.RootFlags.Path != "" {
				k8ConfigPath := a.RootFlags.Path
				k8Config, err = MakeConfigFromPath(k8ConfigPath, context)
				if err != nil {
					return err
				}
			} else if kubeEnv, exists := os.LookupEnv("KUBECONFIG"); exists && kubeEnv != "" {
				k8ConfigPath := os.Getenv("KUBECONFIG")
				k8Config, err = MakeConfigFromPath(k8ConfigPath, context)
				if err != nil {
					return err
				}
			} else if a.RootFlags.Url != "" {
				k8ConfigUrl := a.RootFlags.Url
				k8Config = MakeConfigFromUrl(k8ConfigUrl)
			} else {
				err := errors.New("please provide either a path to a config file, " +
					"assign $KUBECONFIG to the config file path, " +
					"or provide a URL to the cluster")
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

	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Quiet, "quiet", "q", false, "Suppress output")
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Verbose, "verbose", "v", false, "Verbose output")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.Context, "context", "c", "", "Context name")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.Path, "path", "p", "", "Config path")
	a.RootCmd.PersistentFlags().StringVarP(&a.RootFlags.Url, "url", "u", "", "Cluster url")
	a.RootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "Path to output file. If blank, will output to STDOUT")
	a.RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "signal", "Output format (signal, json, yaml). Default value is signal")

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

// Generates the k8s config object from a k8s cluster Url
func MakeConfigFromUrl(clusterUrl string) *rest.Config {
	return &rest.Config{
		Host: clusterUrl,
	}
}

// Generates the k8s config object from a path to a config file
func MakeConfigFromPath(configPath string, context string) (*rest.Config, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath}
	configOverrides := &clientcmd.ConfigOverrides{}

	if context != "" {
		configOverrides.CurrentContext = context
	}

	k8Config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).ClientConfig()
	if err != nil {
		return nil, err
	}
	return k8Config, nil
}
