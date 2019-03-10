package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bitnode",
	Short: "Cluster managament application",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.Set("api.version", version)
}

func init() {
	user, err := user.Current()
	if err != nil {
		log.Println("Error getting current user info, defaulting configuration to current directory")
		rootCmd.PersistentFlags().StringP("config", "c", "bitnode_config.db", "Path to bitnode config.db file")
	} else {
		rootCmd.PersistentFlags().StringP("config", "c", user.HomeDir+"/.bitnode/config.db", "Path to bitnode config.db file")
	}

	viper.BindPFlag("bitnode.db", rootCmd.PersistentFlags().Lookup("config"))
}
