package swmintegration

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type NodeType string

const (
	COMPUTE_NODE NodeType = "COMPUTE"
	NETWORK_NODE NodeType = "NETWORK"
)

// TopologyNodeSpec describes one node in the topology,
// can be a compute node or a network node.
type TopologyNodeSpec struct {
	// The name of this node. Must be unique within the topology.
	// If this node is also a Kubernetes node, use the name
	// of the Kubernetes node here.
	Name string `json:"name,omitempty"`

	// Values are COMPUTE or NETWORK.
	// +kubebuilder:default:=COMPUTE
	Type NodeType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=NodeType"`
}

// TopologyLinkCapabilities are the QoS capabilities of a Link.
type TopologyLinkCapabilities struct {
	// Bandwidth capacity of a link.
	// It is specified in bit/s, e.g. 5M means 5Mbit/s.
	// +kubebuilder:validation:Pattern:=^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
	BandWidthBits string `json:"bandWidthBits,omitempty"`

	// Worst-case delay in nanoseconds.
	// You can use scientific notation, so 10e6 is 1ms.
	// +kubebuilder:validation:Pattern:=^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
	LatencyNanos string `json:"latencyNanos,omitempty"`

	// Network type (implementation)-dependent
	// qos information.
	// +optional
	OtherCapabilities map[string]string `json:"otherCapabilities"`
}

// TopologyLinkSpec is one link in the topology. Links are directed.
type TopologyLinkSpec struct {
	// The source node of this link.
	Source string `json:"source"`

	// The target node of this link
	Target string `json:"target"`

	// The link's QoS Capabilities.
	// These are required for physical links but optional
	// for logical links. Empty capabilities on a logical
	// link means the link will use the capabilities of the
	// underlying physical link.
	// If a logical link specifies capabilities in excess of
	// what the underlying physical link supports, it is up
	// to the corresponding network operator to decide what to do.
	// +optional
	Capabilities TopologyLinkCapabilities `json:"capabilities"`
}

// TopologyPathSpec is a list of nodes connected by a path.
type TopologyPathSpec struct {

	// The list of nodes that this path traverses.
	// For every consecutive pair of nodes on this path,
	// a corresponding TopologyLink must exist.
	Nodes []string `json:"nodes"`
}

// NetworkTopologySpec describes the topology and
// provides the name of the network type (network implementation).
// The default Topology operator will not delete a network when
// the topology spec is deleted.
type NetworkTopologySpec struct {

	// This is the value of the network-implementation tag that
	// will be attached to NetworkLinks and NetworkPaths created from
	// this topology.
	// It is used to distinguish between different network types
	// with different QoS capabilities.
	NetworkImplementation string `json:"networkImplementation"`

	// The network-implementation tag of the physical network that
	// this network's links are based on.
	// When this is empty, this topology will be considered to declare
	// a physical network.
	// We might need to have a list of bases here and then let each link
	// declare which of those bases it uses.
	PhysicalBase string `json:"physicalBase"`

	// The nodes in your topology.
	Nodes []TopologyNodeSpec `json:"nodes,omitempty"`

	// All the links in your topology. Links are directed, so if you
	// want both a->b and b->a to exist, you need to specify both.
	// Loopback links (from each node to itself) will be inserted
	// automatically.
	Links []TopologyLinkSpec `json:"links,omitempty"`
}

// NetworkTopology specifies how the nodes are connected in a network.
type NetworkTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkTopologySpec `json:"spec,omitempty"`
}

// NetworkTopologyList contains a list of NetworkTopology
type NetworkTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkTopology `json:"items"`
}

func (networkTopology NetworkTopology) GetUnstructuredData() *unstructured.Unstructured {

	jsonData, err := json.Marshal(networkTopology)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON into a map[string]interface{} to prepare for unstructured conversion
	var objMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &objMap); err != nil {
		panic(err)
	}
	unstructuredObj := &unstructured.Unstructured{Object: objMap}

	unstructuredObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "qos-scheduler.siemens.com",
		Version: "v1alpha1",
		Kind:    "NetworkTopology",
	})

	// Create an unstructured.Unstructured object from the map
	return unstructuredObj

}
