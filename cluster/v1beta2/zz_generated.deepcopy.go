//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright Contributors to the Open Cluster Management project
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSelector) DeepCopyInto(out *ManagedClusterSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSelector.
func (in *ManagedClusterSelector) DeepCopy() *ManagedClusterSelector {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSet) DeepCopyInto(out *ManagedClusterSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSet.
func (in *ManagedClusterSet) DeepCopy() *ManagedClusterSet {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetBinding) DeepCopyInto(out *ManagedClusterSetBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetBinding.
func (in *ManagedClusterSetBinding) DeepCopy() *ManagedClusterSetBinding {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterSetBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetBindingList) DeepCopyInto(out *ManagedClusterSetBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedClusterSetBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetBindingList.
func (in *ManagedClusterSetBindingList) DeepCopy() *ManagedClusterSetBindingList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterSetBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetBindingSpec) DeepCopyInto(out *ManagedClusterSetBindingSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetBindingSpec.
func (in *ManagedClusterSetBindingSpec) DeepCopy() *ManagedClusterSetBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetBindingStatus) DeepCopyInto(out *ManagedClusterSetBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetBindingStatus.
func (in *ManagedClusterSetBindingStatus) DeepCopy() *ManagedClusterSetBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetList) DeepCopyInto(out *ManagedClusterSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedClusterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetList.
func (in *ManagedClusterSetList) DeepCopy() *ManagedClusterSetList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetSpec) DeepCopyInto(out *ManagedClusterSetSpec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetSpec.
func (in *ManagedClusterSetSpec) DeepCopy() *ManagedClusterSetSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSetStatus) DeepCopyInto(out *ManagedClusterSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSetStatus.
func (in *ManagedClusterSetStatus) DeepCopy() *ManagedClusterSetStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSetStatus)
	in.DeepCopyInto(out)
	return out
}
