package cmd

import (
	"github.com/urfave/cli/v2"
)

// PodNamespaceArg is the CLI argument for the pod namespace
const PodNamespaceArg string = "pod-namespace"

// PodNamespaceFlag is the urfave/cli Flag configuration for the pod namespace
var PodNamespaceFlag = &cli.StringFlag{
	Name:    PodNamespaceArg,
	Usage:   "Specifies the namespace that current application pod is running in.",
	EnvVars: []string{"POD_NAMESPACE"},
}
