// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	rest "k8s.io/client-go/rest"
	scheme "open-cluster-management.io/api/client/cluster/clientset/versioned/scheme"
	clusterv1 "open-cluster-management.io/api/cluster/v1"
)

type ClusterV1Interface interface {
	RESTClient() rest.Interface
	ManagedClustersGetter
}

// ClusterV1Client is used to interact with features provided by the cluster.open-cluster-management.io group.
type ClusterV1Client struct {
	restClient rest.Interface
}

func (c *ClusterV1Client) ManagedClusters() ManagedClusterInterface {
	return newManagedClusters(c)
}

// NewForConfig creates a new ClusterV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ClusterV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ClusterV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ClusterV1Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterV1Client for the given RESTClient.
func New(c rest.Interface) *ClusterV1Client {
	return &ClusterV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := clusterv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ClusterV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
