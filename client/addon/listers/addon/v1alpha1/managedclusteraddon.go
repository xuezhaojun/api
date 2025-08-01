// Copyright Contributors to the Open Cluster Management project
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	addonv1alpha1 "open-cluster-management.io/api/addon/v1alpha1"
)

// ManagedClusterAddOnLister helps list ManagedClusterAddOns.
// All objects returned here must be treated as read-only.
type ManagedClusterAddOnLister interface {
	// List lists all ManagedClusterAddOns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*addonv1alpha1.ManagedClusterAddOn, err error)
	// ManagedClusterAddOns returns an object that can list and get ManagedClusterAddOns.
	ManagedClusterAddOns(namespace string) ManagedClusterAddOnNamespaceLister
	ManagedClusterAddOnListerExpansion
}

// managedClusterAddOnLister implements the ManagedClusterAddOnLister interface.
type managedClusterAddOnLister struct {
	listers.ResourceIndexer[*addonv1alpha1.ManagedClusterAddOn]
}

// NewManagedClusterAddOnLister returns a new ManagedClusterAddOnLister.
func NewManagedClusterAddOnLister(indexer cache.Indexer) ManagedClusterAddOnLister {
	return &managedClusterAddOnLister{listers.New[*addonv1alpha1.ManagedClusterAddOn](indexer, addonv1alpha1.Resource("managedclusteraddon"))}
}

// ManagedClusterAddOns returns an object that can list and get ManagedClusterAddOns.
func (s *managedClusterAddOnLister) ManagedClusterAddOns(namespace string) ManagedClusterAddOnNamespaceLister {
	return managedClusterAddOnNamespaceLister{listers.NewNamespaced[*addonv1alpha1.ManagedClusterAddOn](s.ResourceIndexer, namespace)}
}

// ManagedClusterAddOnNamespaceLister helps list and get ManagedClusterAddOns.
// All objects returned here must be treated as read-only.
type ManagedClusterAddOnNamespaceLister interface {
	// List lists all ManagedClusterAddOns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*addonv1alpha1.ManagedClusterAddOn, err error)
	// Get retrieves the ManagedClusterAddOn from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*addonv1alpha1.ManagedClusterAddOn, error)
	ManagedClusterAddOnNamespaceListerExpansion
}

// managedClusterAddOnNamespaceLister implements the ManagedClusterAddOnNamespaceLister
// interface.
type managedClusterAddOnNamespaceLister struct {
	listers.ResourceIndexer[*addonv1alpha1.ManagedClusterAddOn]
}
