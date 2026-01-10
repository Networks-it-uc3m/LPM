package cmd

import (
	"os"

	lpmv1 "github.com/Networks-it-uc3m/LPM/api/v1"
	"github.com/spf13/cobra"
)

var cf string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lpm",
	Short: "L2S-M Performance Measurements (LPM) CLI",
	Long: `L2S-M Performance Measurements (LPM) is a module that measures and exposes
network performance metrics (e.g., RTT, jitter, throughput) over the L2S-M overlay
within a single Kubernetes cluster.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cf, "config_file", lpmv1.DEFAULT_CONFIG_FILE, "configuration path where config.json and topology.json are going to be placed.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("grpc_server", "", false, "Help message for toggle")
}
