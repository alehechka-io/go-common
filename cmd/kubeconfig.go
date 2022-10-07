package cmd

import (
	"path/filepath"

	"github.com/urfave/cli/v2"
	"k8s.io/client-go/util/homedir"
)

// OutOfClusterArg is the CLI argument to specify out-of-cluster usage
const OutOfClusterArg string = "out-of-cluster"

// OutOfClusterFlag is the urfave/cli Flag configuration for out-of-cluster usage
var OutOfClusterFlag = &cli.BoolFlag{
	Name:    OutOfClusterArg,
	Usage:   "Will use the default ~/.kube/config file on the local machine to connect to the cluster externally.",
	Aliases: []string{"local"},
}

// KubeConfigArg is the CLI argument for the location of the kube config file
const KubeConfigArg string = "kubeconfig"

// KubeConfig prepares a urfave/cli Flag configuration for the kube config file
func KubeConfig() *cli.StringFlag {
	kubeconfig := &cli.StringFlag{Name: KubeConfigArg}
	if home := homedir.HomeDir(); home != "" {
		kubeconfig.Value = filepath.Join(home, ".kube", "config")
		kubeconfig.Usage = "(optional) absolute path to the kubeconfig file"
	} else {
		kubeconfig.Usage = "absolute path to the kubeconfig file (required if running OutOfCluster)"
	}
	return kubeconfig
}
