package node

import (
	"time"

	"github.com/bitnodesnet/bitnode/controller"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ScanCmd() *cobra.Command {
	nodeScanCmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan for new nodes on the local network",
		Run: func(cmd *cobra.Command, args []string) {
			controller := controller.GetController()
			controller.FindNodes()
			defer controller.CleanExit()
		},
	}

	nodeScanCmd.Flags().DurationP("duration", "d", 5*time.Second, "Duration to scan for nodes")
	nodeScanCmd.Flags().String("mdns_service", "_bitnode._tcp", "Service name for mdns service discovery")

	viper.BindPFlag("scan.duration", nodeScanCmd.Flags().Lookup("duration"))
	viper.BindPFlag("mdns.service", nodeScanCmd.Flags().Lookup("mdns_service"))

	return nodeScanCmd
}
