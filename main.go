package main

import (
	"flag"
	"os"

	"github.com/method-security/methodk8s/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	methodk8s := cmd.NewMethodK8s(version)
	methodk8s.InitRootCommand()

	methodk8s.InitIngressCommand()
	methodk8s.InitNodeCommand()
	methodk8s.InitPodCommand()
	methodk8s.InitServiceCommand()

	if err := methodk8s.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
