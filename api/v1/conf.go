package v1

import (
	"encoding/json"
)

const (
	DEFAULT_CONFIG_FILE = "/etc/lpm/lpm-config.json"
)

type MetricConfiguration struct {
	Name       string `json:"name"`
	IP         string `json:"ip"`
	RTT        int    `json:"rttInterval,omitempty"`
	Jitter     int    `json:"jitterInterval,omitempty"`
	Throughput int    `json:"throughputInterval,omitempty"`
	OTD        int    `json:"otdInterval,omitempty"`
}

type NodeConfig struct {
	NodeName              string                `json:"Nodename"`
	SpreadFactor          float64               `json:"spreadFactor,omitempty"`
	IpAddress             string                `json:"ipAddress,omitempty"`
	ProbeInterface        string                `json:"probeInterface"`
	MetricsNeighbourNodes []MetricConfiguration `json:"MetricsNeighbourNodes"`
}

func (conf *MetricConfiguration) UnmarshalJSON(data []byte) error {
	type confAlias MetricConfiguration
	defaultConf := &confAlias{RTT: -1, Jitter: -1, Throughput: -1}

	err := json.Unmarshal(data, defaultConf)
	if err != nil {
		return err
	}

	*conf = MetricConfiguration(*defaultConf)
	return nil
}
