package collector

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func StartCollector() {

	for index := range lpmInstance.Metrics {
		go func(lpmDataIndex int) {
			lpmInstance.Metrics[lpmDataIndex].RunPeriodicTests()
		}(index)
	}

	for index := range lpmInstance.Servers {
		go func(lpmDataIndex int) {
			lpmInstance.Servers[lpmDataIndex]()
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

func (metric *Metric) RunPeriodicTests() {

	log.Infof("Testing %s in network link between node %s and node %s", metric.Name, metric.SourceNodeName, metric.TargetNodeName)

	if metric.TestTimeInterval != -1 {
		for true {

			metric.value = metric.method(metric.TargetNodeIp)
			log.Infof(" %s between node %s and node %s is %f.", metric.Name, metric.SourceNodeName, metric.TargetNodeName, metric.value)

			time.Sleep(time.Duration(metric.TestTimeInterval) * time.Minute)
		}
	}

}