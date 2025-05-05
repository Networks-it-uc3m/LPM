package collector

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const SF_DEFAULT_VALUE = 0.2

type ServerMethod func()

type LPMInstance struct {
	NodeName     string
	promReg      *prometheus.Registry
	Metrics      []Metric
	Servers      []ServerMethod
	SpreadFactor float64
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

func (lpmInstance *LPMInstance) SetSpreadFactor(sf float64) {

	if sf > 0.0 && sf <= 1.0 {
		lpmInstance.SpreadFactor = sf

	} else {
		lpmInstance.SpreadFactor = SF_DEFAULT_VALUE

	}
}

func (lpmInstance *LPMInstance) AddMetric(metricName string, targetNodeName string, metricInterval int, targetNodeIP string, measureMethod MeasureMethod) {

	metricId := MetricId{Name: metricName, SourceNodeName: lpmInstance.NodeName, TargetNodeName: targetNodeName}
	metricId.GenerateMetricId()
	lpmInstance.Metrics = append(lpmInstance.Metrics,
		Metric{MetricData: MetricData{metricId, 0.0},
			TargetNodeIp:     targetNodeIP,
			TestTimeInterval: metricInterval,
			method:           measureMethod,
			SpreadFactor:     lpmInstance.SpreadFactor})
	lpmCollector := lpmExporterCollector(metricId)
	lpmInstance.promReg.MustRegister(lpmCollector)
}

func (lpmInstance *LPMInstance) AddServer(serverMethod ServerMethod) {
	lpmInstance.Servers = append(lpmInstance.Servers, serverMethod)
}

func GetMetricValue(metricName, sourceNode, targetNode string) float64 {
	for _, metric := range lpmInstance.Metrics {
		if metric.Name == metricName && metric.SourceNodeName == sourceNode && metric.TargetNodeName == targetNode {
			return metric.Value
		}
	}
	return 0.0
}
