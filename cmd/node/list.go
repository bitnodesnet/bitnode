package node

import (
	"github.com/bitnodesnet/bitnode/controller"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	nodeListCmd := &cobra.Command{
		Use:   "list",
		Short: "List known nodes",
		Run: func(cmd *cobra.Command, args []string) {
			controller := controller.GetController()
			controller.ListNodes()
			defer controller.CleanExit()
		},
	}

	return nodeListCmd
}
