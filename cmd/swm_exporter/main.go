package main

import (
	"time"

	"github.com/Networks-it-uc3m/LPM/internal/swmintegration"
)

func main() {

	swmintegration.RunExporter(time.Minute * 5)

}
