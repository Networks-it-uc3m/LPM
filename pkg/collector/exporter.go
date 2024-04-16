package collector

import (
	"time"

	"github.com/Networks-it-uc3m/LPM/internal/swmintegration"
	//"github.com/Networks-it-uc3m/L2S-M/src/database/pkg/databaseClient"
)

func RunExporter(duration time.Duration) {

	for {

		swmClient := swmintegration.SWMClient{}

		swmClient.NewClient()

		//networkTopology := databaseClient.Get("topology")
		networkTopology := swmintegration.HardcodeTopology()

		//networkTopology.FillTopologyWithMetrics()

		swmClient.ExportCRD("he-codeco-swm", networkTopology)

		time.Sleep(duration)
	}
}
