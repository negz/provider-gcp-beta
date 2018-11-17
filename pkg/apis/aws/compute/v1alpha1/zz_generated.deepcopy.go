// +build !ignore_autogenerated

/*
Copyright 2018 The Conductor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (e *EKSCluster) DeepCopyInto(out *EKSCluster) {
	*out = *e
	out.TypeMeta = e.TypeMeta
	e.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	e.Spec.DeepCopyInto(&out.Spec)
	e.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCluster.
func (e *EKSCluster) DeepCopy() *EKSCluster {
	if e == nil {
		return nil
	}
	out := new(EKSCluster)
	e.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (e *EKSCluster) DeepCopyObject() runtime.Object {
	if c := e.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterList) DeepCopyInto(out *EKSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterList.
func (in *EKSClusterList) DeepCopy() *EKSClusterList {
	if in == nil {
		return nil
	}
	out := new(EKSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterSpec) DeepCopyInto(out *EKSClusterSpec) {
	*out = *in
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupsIds != nil {
		in, out := &in.SecurityGroupsIds, &out.SecurityGroupsIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ClassRef != nil {
		in, out := &in.ClassRef, &out.ClassRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ConnectionSecretRef != nil {
		in, out := &in.ConnectionSecretRef, &out.ConnectionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterSpec.
func (in *EKSClusterSpec) DeepCopy() *EKSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EKSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSClusterStatus) DeepCopyInto(out *EKSClusterStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.BindingStatusPhase = in.BindingStatusPhase
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSClusterStatus.
func (in *EKSClusterStatus) DeepCopy() *EKSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EKSClusterStatus)
	in.DeepCopyInto(out)
	return out
}
