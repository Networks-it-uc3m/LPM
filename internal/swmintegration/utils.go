package swmintegration

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func HardcodeTopology() NetworkTopology {
	networkTopology := NetworkTopology{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "qos-scheduler.siemens.com/v1alpha1",
			Kind:       "NetworkTopology",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "l2sm-sample-cluster",
		},
		Spec: NetworkTopologySpec{
			NetworkImplementation: "l2sm-network",
			PhysicalBase:          "logical-network",
			Nodes: []TopologyNodeSpec{
				{Name: "node-a", Type: COMPUTE_NODE},
				{Name: "node-b", Type: COMPUTE_NODE},
				{Name: "node-c", Type: COMPUTE_NODE},
				{Name: "node-d", Type: COMPUTE_NODE},
				{Name: "node-e", Type: COMPUTE_NODE},
			},
			Links: []TopologyLinkSpec{
				{Source: "node-a", Target: "node-b", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1G", LatencyNanos: "2e6"}},
				{Source: "node-a", Target: "node-c", Capabilities: TopologyLinkCapabilities{BandWidthBits: "500M", LatencyNanos: "3e6"}},
				{Source: "node-b", Target: "node-a", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1G", LatencyNanos: "2e6"}},
				{Source: "node-b", Target: "node-c", Capabilities: TopologyLinkCapabilities{BandWidthBits: "2G", LatencyNanos: "1e6"}},
				{Source: "node-b", Target: "node-d", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1.5G", LatencyNanos: "2.5e6"}},
				{Source: "node-c", Target: "node-a", Capabilities: TopologyLinkCapabilities{BandWidthBits: "500M", LatencyNanos: "3e6"}},
				{Source: "node-c", Target: "node-b", Capabilities: TopologyLinkCapabilities{BandWidthBits: "2G", LatencyNanos: "1e6"}},
				{Source: "node-c", Target: "node-d", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1G", LatencyNanos: "2e6"}},
				{Source: "node-d", Target: "node-b", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1G", LatencyNanos: "2e6"}},
				{Source: "node-d", Target: "node-c", Capabilities: TopologyLinkCapabilities{BandWidthBits: "1G", LatencyNanos: "2e6"}},
				{Source: "node-d", Target: "node-e", Capabilities: TopologyLinkCapabilities{BandWidthBits: "2G", LatencyNanos: "2.5e6"}},
				{Source: "node-e", Target: "node-d", Capabilities: TopologyLinkCapabilities{BandWidthBits: "2G", LatencyNanos: "2.5e6"}},
			},
		},
	}
	return networkTopology
}
