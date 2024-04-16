package exporterclient

import (
	"context"
	"encoding/json"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

type ExporterClient interface {
	NewClient(schemaGVR schema.GroupVersionResource)
	ExportCRD(namespace string, crd StructCustomResourceDefinition)
	GetSchemaGVR() schema.GroupVersionResource
	GetDynamicClient() dynamic.DynamicClient
}

type BasicClient struct {
	SchemaGVR     schema.GroupVersionResource
	DynamicClient *dynamic.DynamicClient
}

func (basicClient *BasicClient) NewClient() error {

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

	schemaGVR := schema.GroupVersionResource{Group: "group", Version: "version", Resource: "resource"}

	basicClient = &BasicClient{SchemaGVR: schemaGVR, DynamicClient: clientset}

	return nil
}

func (basicClient *BasicClient) ExportCRD(namespace string, structCRD StructCustomResourceDefinition) {
	jsonData, err := json.Marshal(structCRD)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON into a map[string]interface{} to prepare for unstructured conversion
	var objMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &objMap); err != nil {
		panic(err)
	}

	// Create an unstructured.Unstructured object from the map
	unstructuredObj := &unstructured.Unstructured{Object: objMap}
	unstructuredObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "group",
		Version: "version",
		Kind:    "Resource",
	})

	basicClient.DynamicClient.Resource(basicClient.SchemaGVR).Namespace(namespace).Create(context.Background(), unstructuredObj, metav1.CreateOptions{})

}
