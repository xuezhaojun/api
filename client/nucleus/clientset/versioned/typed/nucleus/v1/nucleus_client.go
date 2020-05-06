// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"github.com/open-cluster-management/api/client/nucleus/clientset/versioned/scheme"
	v1 "github.com/open-cluster-management/api/nucleus/v1"
	rest "k8s.io/client-go/rest"
)

type NucleusV1Interface interface {
	RESTClient() rest.Interface
	HubCoresGetter
	SpokeCoresGetter
}

// NucleusV1Client is used to interact with features provided by the nucleus.open-cluster-management.io group.
type NucleusV1Client struct {
	restClient rest.Interface
}

func (c *NucleusV1Client) HubCores() HubCoreInterface {
	return newHubCores(c)
}

func (c *NucleusV1Client) SpokeCores() SpokeCoreInterface {
	return newSpokeCores(c)
}

// NewForConfig creates a new NucleusV1Client for the given config.
func NewForConfig(c *rest.Config) (*NucleusV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NucleusV1Client{client}, nil
}

// NewForConfigOrDie creates a new NucleusV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NucleusV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NucleusV1Client for the given RESTClient.
func New(c rest.Interface) *NucleusV1Client {
	return &NucleusV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NucleusV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
