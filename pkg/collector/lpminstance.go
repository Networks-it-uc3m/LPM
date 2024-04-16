package collector

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type ServerMethod func()

type LPMInstance struct {
	NodeName string
	promReg  *prometheus.Registry
	Metrics  []Metric
	Servers  []ServerMethod
}

var lock = &sync.Mutex{}

var lpmInstance *LPMInstance

func GetInstance() *LPMInstance {
	if lpmInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if lpmInstance == nil {
			fmt.Println("Creating single instance now.")
			lpmInstance = &LPMInstance{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	lpmInstance.promReg = prometheus.NewRegistry()
	return lpmInstance
}

func (lpmInstance *LPMInstance) SetNodeName(name string) {
	lpmInstance.NodeName = name
}

func (lpmInstance *LPMInstance) AddMetric(metricName string, targetNodeName string, metricInterval int, targetNodeIP string, measureMethod MeasureMethod) {
	metricId := GenerateMetricId(metricName, lpmInstance.NodeName, targetNodeName)
	lpmInstance.Metrics = append(lpmInstance.Metrics, Metric{Name: metricName, SourceNodeName: lpmInstance.NodeName, TargetNodeName: targetNodeName, TestTimeInterval: metricInterval, MetricId: metricId, method: measureMethod, value: 0.0})
	lpmCollector := lpmExporterCollector(metricId)
	lpmInstance.promReg.MustRegister(lpmCollector)
}

func (lpmInstance *LPMInstance) AddServer(serverMethod ServerMethod) {
	lpmInstance.Servers = append(lpmInstance.Servers, serverMethod)
}

func GetMetricValueFromId(metricId string) float64 {
	for _, metric := range lpmInstance.Metrics {
		if metric.MetricId == metricId {
			return metric.value
		}
	}
	return 0.0
}
