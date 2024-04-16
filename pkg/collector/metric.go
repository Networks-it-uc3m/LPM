package collector

import (
	"fmt"
	"strings"
)

type MeasureMethod func(targetNodeIP string) float64

type Metric struct {
	Name             string
	SourceNodeName   string
	TargetNodeName   string
	TargetNodeIp     string
	TestTimeInterval int
	MetricId         string
	method           MeasureMethod
	value            float64
}

func GenerateMetricId(metricName, sourceNode, targetNode string) string {
	return fmt.Sprintf("%s_%s:%s", metricName, sourceNode, targetNode)
}

// DecomposeMetricId takes a metric ID and returns the original metric name, source node, and target node.
func DecomposeMetricId(metricId string) (string, string, string) {
	// Split the string at the first underscore to isolate the metric name.
	parts := strings.SplitN(metricId, "_", 2)
	if len(parts) < 2 {
		return "", "", "" // Return empty strings if the format is incorrect.
	}
	metricName := parts[0]

	// Split the second part at the colon to separate the source and target nodes.
	nodes := strings.Split(parts[1], ":")
	if len(nodes) < 2 {
		return metricName, "", "" // Return the metric name and empty strings for nodes if format is incorrect.
	}

	sourceNode, targetNode := nodes[0], nodes[1]
	return metricName, sourceNode, targetNode
}
