// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "open-cluster-management.io/api/client/work/clientset/versioned/scheme"
	workv1 "open-cluster-management.io/api/work/v1"
)

// ManifestWorksGetter has a method to return a ManifestWorkInterface.
// A group's client should implement this interface.
type ManifestWorksGetter interface {
	ManifestWorks(namespace string) ManifestWorkInterface
}

// ManifestWorkInterface has methods to work with ManifestWork resources.
type ManifestWorkInterface interface {
	Create(ctx context.Context, manifestWork *workv1.ManifestWork, opts metav1.CreateOptions) (*workv1.ManifestWork, error)
	Update(ctx context.Context, manifestWork *workv1.ManifestWork, opts metav1.UpdateOptions) (*workv1.ManifestWork, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, manifestWork *workv1.ManifestWork, opts metav1.UpdateOptions) (*workv1.ManifestWork, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*workv1.ManifestWork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*workv1.ManifestWorkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *workv1.ManifestWork, err error)
	ManifestWorkExpansion
}

// manifestWorks implements ManifestWorkInterface
type manifestWorks struct {
	*gentype.ClientWithList[*workv1.ManifestWork, *workv1.ManifestWorkList]
}

// newManifestWorks returns a ManifestWorks
func newManifestWorks(c *WorkV1Client, namespace string) *manifestWorks {
	return &manifestWorks{
		gentype.NewClientWithList[*workv1.ManifestWork, *workv1.ManifestWorkList](
			"manifestworks",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *workv1.ManifestWork { return &workv1.ManifestWork{} },
			func() *workv1.ManifestWorkList { return &workv1.ManifestWorkList{} },
		),
	}
}
