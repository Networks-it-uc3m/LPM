package swmintegration

import (
	"context"

	"github.com/Networks-it-uc3m/LPM/pkg/exporterclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// Uses ExporterClient interface
type SWMClient struct {
	SchemaGVR     schema.GroupVersionResource
	DynamicClient *dynamic.DynamicClient
}

func (swmClient *SWMClient) GetSchemaGVR() schema.GroupVersionResource {
	return swmClient.SchemaGVR
}
func (swmClient *SWMClient) GetDynamicClient() dynamic.DynamicClient {
	return *swmClient.DynamicClient
}

func (swmClient *SWMClient) NewClient() error {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := dynamic.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	schemaGVR := schema.GroupVersionResource{Group: "qos-scheduler.siemens.com", Version: "v1alpha1", Resource: "networktopologies"}

	swmClient = &SWMClient{SchemaGVR: schemaGVR, DynamicClient: clientset}

	return nil
}

func (swmClient *SWMClient) ExportCRD(namespace string, networkTopology exporterclient.StructCustomResourceDefinition) {

	unstructuredObj := networkTopology.GetUnstructuredData()

	swmClient.DynamicClient.Resource(swmClient.SchemaGVR).Namespace(namespace).Create(context.Background(), unstructuredObj, metav1.CreateOptions{})

}
