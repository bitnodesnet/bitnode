package cmd

import (
	"github.com/bitnodesnet/bitnode/cmd/node"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Commands for managing nodes withing the bitnode",
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.AddCommand(node.DeleteCmd())
	nodeCmd.AddCommand(node.DeployCmd())
	nodeCmd.AddCommand(node.ListCmd())
	nodeCmd.AddCommand(node.NameCmd())
	nodeCmd.AddCommand(node.ScanCmd())
	nodeCmd.AddCommand(node.StatsCmd())

}
