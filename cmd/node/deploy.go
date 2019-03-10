package node

import (
	"github.com/bitnodesnet/bitnode/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func DeployCmd() *cobra.Command {
	nodeDeployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Start node worker for the bitnode",
		Long:  "Adds a worker node into the bitnode that will listen for commands from the controller",
		Run: func(cmd *cobra.Command, args []string) {
			node := node.Node{
				Version: viper.GetString("api.version"),
			}

			node.Start()
		},
	}

	nodeDeployCmd.Flags().StringP("port", "p", "10000", "port for grpc service")
	nodeDeployCmd.Flags().String("mdns_service", "_bitnode._tcp", "Service name for mdns service discovery")
	viper.BindPFlag("rpc.port", nodeDeployCmd.Flags().Lookup("port"))
	viper.BindPFlag("mdns.service", nodeDeployCmd.Flags().Lookup("mdns_service"))

	return nodeDeployCmd
}
