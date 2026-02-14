package collector

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func StartCollector() {

	for index := range lpmInstance.Metrics {
		go func(lpmDataIndex int) {
			lpmInstance.Metrics[lpmDataIndex].RunPeriodicTests(lpmInstance.ProbeInterface)
		}(index)
	}

	for index := range lpmInstance.Servers {
		go func(lpmDataIndex int) {
			lpmInstance.Servers[lpmDataIndex](lpmInstance.ProbeInterface)
		}(index)
	}
	handler := promhttp.HandlerFor(
		lpmInstance.promReg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		})

	http.Handle("/metrics", handler)

	http.ListenAndServe(":8090", nil)
}

func (metric *Metric) RunPeriodicTests(probeInterface string) {

	if metric.TestTimeInterval == -1 {
		return
	}

	log.Infof("Testing %s in network link between node %s and node %s", metric.Name, metric.SourceNodeName, metric.TargetNodeName)

	for {
		for i := 1; i < 4; i++ {
			maxDelay := (time.Duration(metric.TestTimeInterval) * time.Minute) * time.Duration(metric.SpreadFactor*10) / 10
			minDelay := time.Second
			randomFactor := minDelay + time.Duration(rand.Float64()*10)*(maxDelay-minDelay)/10
			jitter := time.Duration(1<<i) * time.Duration(10*(rand.Float64()*1.0+0.5)) / 10 * time.Second
			randomDelay := jitter + randomFactor
			time.Sleep(randomDelay)

			metric.Value = metric.method(metric.TargetNodeIp, probeInterface)

			if metric.Value != 0 {
				break
			}
			log.Infof("Couldn't measure %s between node %s and node %s. Trying again.", metric.Name, metric.SourceNodeName, metric.TargetNodeName)
		}

		log.Infof(" %s between node %s and node %s is %f.", metric.Name, metric.SourceNodeName, metric.TargetNodeName, metric.Value)

		time.Sleep(time.Duration(metric.TestTimeInterval) * time.Minute)
	}

}
