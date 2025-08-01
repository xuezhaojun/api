// Copyright Contributors to the Open Cluster Management project
// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	fmt "fmt"
	http "net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	clusterv1 "open-cluster-management.io/api/client/cluster/clientset/versioned/typed/cluster/v1"
	clusterv1alpha1 "open-cluster-management.io/api/client/cluster/clientset/versioned/typed/cluster/v1alpha1"
	clusterv1beta1 "open-cluster-management.io/api/client/cluster/clientset/versioned/typed/cluster/v1beta1"
	clusterv1beta2 "open-cluster-management.io/api/client/cluster/clientset/versioned/typed/cluster/v1beta2"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ClusterV1() clusterv1.ClusterV1Interface
	ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface
	ClusterV1beta1() clusterv1beta1.ClusterV1beta1Interface
	ClusterV1beta2() clusterv1beta2.ClusterV1beta2Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	clusterV1       *clusterv1.ClusterV1Client
	clusterV1alpha1 *clusterv1alpha1.ClusterV1alpha1Client
	clusterV1beta1  *clusterv1beta1.ClusterV1beta1Client
	clusterV1beta2  *clusterv1beta2.ClusterV1beta2Client
}

// ClusterV1 retrieves the ClusterV1Client
func (c *Clientset) ClusterV1() clusterv1.ClusterV1Interface {
	return c.clusterV1
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return c.clusterV1alpha1
}

// ClusterV1beta1 retrieves the ClusterV1beta1Client
func (c *Clientset) ClusterV1beta1() clusterv1beta1.ClusterV1beta1Interface {
	return c.clusterV1beta1
}

// ClusterV1beta2 retrieves the ClusterV1beta2Client
func (c *Clientset) ClusterV1beta2() clusterv1beta2.ClusterV1beta2Interface {
	return c.clusterV1beta2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.clusterV1, err = clusterv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.clusterV1alpha1, err = clusterv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.clusterV1beta1, err = clusterv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.clusterV1beta2, err = clusterv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.clusterV1 = clusterv1.New(c)
	cs.clusterV1alpha1 = clusterv1alpha1.New(c)
	cs.clusterV1beta1 = clusterv1beta1.New(c)
	cs.clusterV1beta2 = clusterv1beta2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
