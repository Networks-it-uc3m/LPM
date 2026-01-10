/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Networks-it-uc3m/LPM/internal/l2smexporter"
	"github.com/Networks-it-uc3m/LPM/internal/swmintegration"

	"github.com/spf13/cobra"
	// Adjust the import path based on your module path
)

// nedCmd represents the ned command
var exporterCmd = &cobra.Command{
	Use:   "exporter",
	Short: "Export collected metrics to external integrations",
	Long: `Run the LPM exporter process.

The exporter periodically reads the locally collected metrics and publishes them to
an external target. Supported targets:
  - l2sm (default): updates the L2S-M overlay / related custom resources
  - swm: exports to the SWM integration (requires TOPOLOGY_NAMESPACE env var)

Examples:
  lpm exporter
  lpm exporter --target l2sm
  lpm exporter --target swm`,
	Run: func(cmd *cobra.Command, args []string) {

		tg, err := cmd.Flags().GetString("target")
		if err != nil {
			fmt.Println("Error with the target variable. Error:", err)
			return
		}
		switch tg {
		case "swm":
			swmintegration.RunExporter(time.Minute*5, os.Getenv("TOPOLOGY_NAMESPACE"))
		default:
			l2smexporter.RunExporter(time.Minute * 5)
		}
	},
}

func init() {
	rootCmd.AddCommand(exporterCmd)
	exporterCmd.Flags().String(
		"target",
		"l2sm",
		"export target: l2sm (default) or swm",
	)

}
