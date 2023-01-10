// +build !ignore_autogenerated

/*
Copyright 2022 The OpenFunction Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/dapr/dapr/pkg/apis/components/v1alpha1"
	apiv1alpha1 "github.com/kedacore/keda/v2/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildImpl) DeepCopyInto(out *BuildImpl) {
	*out = *in
	if in.Builder != nil {
		in, out := &in.Builder, &out.Builder
		*out = new(string)
		**out = **in
	}
	if in.BuilderCredentials != nil {
		in, out := &in.BuilderCredentials, &out.BuilderCredentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Shipwright != nil {
		in, out := &in.Shipwright, &out.Shipwright
		*out = new(ShipwrightEngine)
		(*in).DeepCopyInto(*out)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SrcRepo != nil {
		in, out := &in.SrcRepo, &out.SrcRepo
		*out = new(GitRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.Dockerfile != nil {
		in, out := &in.Dockerfile, &out.Dockerfile
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.SuccessfulBuildsHistoryLimit != nil {
		in, out := &in.SuccessfulBuildsHistoryLimit, &out.SuccessfulBuildsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedBuildsHistoryLimit != nil {
		in, out := &in.FailedBuildsHistoryLimit, &out.FailedBuildsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.BuilderMaxAge != nil {
		in, out := &in.BuilderMaxAge, &out.BuilderMaxAge
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildImpl.
func (in *BuildImpl) DeepCopy() *BuildImpl {
	if in == nil {
		return nil
	}
	out := new(BuildImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Builder) DeepCopyInto(out *Builder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Builder.
func (in *Builder) DeepCopy() *Builder {
	if in == nil {
		return nil
	}
	out := new(Builder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Builder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderList) DeepCopyInto(out *BuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Builder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderList.
func (in *BuilderList) DeepCopy() *BuilderList {
	if in == nil {
		return nil
	}
	out := new(BuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderSpec) DeepCopyInto(out *BuilderSpec) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Builder != nil {
		in, out := &in.Builder, &out.Builder
		*out = new(string)
		**out = **in
	}
	if in.BuilderCredentials != nil {
		in, out := &in.BuilderCredentials, &out.BuilderCredentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Shipwright != nil {
		in, out := &in.Shipwright, &out.Shipwright
		*out = new(ShipwrightEngine)
		(*in).DeepCopyInto(*out)
	}
	if in.SrcRepo != nil {
		in, out := &in.SrcRepo, &out.SrcRepo
		*out = new(GitRepo)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageCredentials != nil {
		in, out := &in.ImageCredentials, &out.ImageCredentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Dockerfile != nil {
		in, out := &in.Dockerfile, &out.Dockerfile
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderSpec.
func (in *BuilderSpec) DeepCopy() *BuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderStatus) DeepCopyInto(out *BuilderStatus) {
	*out = *in
	if in.ResourceRef != nil {
		in, out := &in.ResourceRef, &out.ResourceRef
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(Output)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderStatus.
func (in *BuilderStatus) DeepCopy() *BuilderStatus {
	if in == nil {
		return nil
	}
	out := new(BuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSourceResult) DeepCopyInto(out *BundleSourceResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSourceResult.
func (in *BundleSourceResult) DeepCopy() *BundleSourceResult {
	if in == nil {
		return nil
	}
	out := new(BundleSourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonRouteSpec) DeepCopyInto(out *CommonRouteSpec) {
	*out = *in
	if in.GatewayRef != nil {
		in, out := &in.GatewayRef, &out.GatewayRef
		*out = new(GatewayRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonRouteSpec.
func (in *CommonRouteSpec) DeepCopy() *CommonRouteSpec {
	if in == nil {
		return nil
	}
	out := new(CommonRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaprIO) DeepCopyInto(out *DaprIO) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaprIO.
func (in *DaprIO) DeepCopy() *DaprIO {
	if in == nil {
		return nil
	}
	out := new(DaprIO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionAddress) DeepCopyInto(out *FunctionAddress) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(AddressType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionAddress.
func (in *FunctionAddress) DeepCopy() *FunctionAddress {
	if in == nil {
		return nil
	}
	out := new(FunctionAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ImageCredentials != nil {
		in, out := &in.ImageCredentials, &out.ImageCredentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(RouteImpl)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(BuildImpl)
		(*in).DeepCopyInto(*out)
	}
	if in.Serving != nil {
		in, out := &in.Serving, &out.Serving
		*out = new(ServingImpl)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceImpl)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(RouteStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Condition)
		**out = **in
	}
	if in.Serving != nil {
		in, out := &in.Serving, &out.Serving
		*out = new(Condition)
		**out = **in
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]FunctionAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(Revision)
		**out = **in
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayRef) DeepCopyInto(out *GatewayRef) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(v1alpha2.Namespace)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayRef.
func (in *GatewayRef) DeepCopy() *GatewayRef {
	if in == nil {
		return nil
	}
	out := new(GatewayRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepo) DeepCopyInto(out *GitRepo) {
	*out = *in
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	if in.SourceSubPath != nil {
		in, out := &in.SourceSubPath, &out.SourceSubPath
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepo.
func (in *GitRepo) DeepCopy() *GitRepo {
	if in == nil {
		return nil
	}
	out := new(GitRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSourceResult) DeepCopyInto(out *GitSourceResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSourceResult.
func (in *GitSourceResult) DeepCopy() *GitSourceResult {
	if in == nil {
		return nil
	}
	out := new(GitSourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaScaleOptions) DeepCopyInto(out *KedaScaleOptions) {
	*out = *in
	if in.ScaledObject != nil {
		in, out := &in.ScaledObject, &out.ScaledObject
		*out = new(KedaScaledObject)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaledJob != nil {
		in, out := &in.ScaledJob, &out.ScaledJob
		*out = new(KedaScaledJob)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaScaleOptions.
func (in *KedaScaleOptions) DeepCopy() *KedaScaleOptions {
	if in == nil {
		return nil
	}
	out := new(KedaScaleOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaScaledJob) DeepCopyInto(out *KedaScaledJob) {
	*out = *in
	if in.RestartPolicy != nil {
		in, out := &in.RestartPolicy, &out.RestartPolicy
		*out = new(v1.RestartPolicy)
		**out = **in
	}
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	in.ScalingStrategy.DeepCopyInto(&out.ScalingStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaScaledJob.
func (in *KedaScaledJob) DeepCopy() *KedaScaledJob {
	if in == nil {
		return nil
	}
	out := new(KedaScaledJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KedaScaledObject) DeepCopyInto(out *KedaScaledObject) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(apiv1alpha1.AdvancedConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KedaScaledObject.
func (in *KedaScaledObject) DeepCopy() *KedaScaledObject {
	if in == nil {
		return nil
	}
	out := new(KedaScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Output) DeepCopyInto(out *Output) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Output.
func (in *Output) DeepCopy() *Output {
	if in == nil {
		return nil
	}
	out := new(Output)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Revision) DeepCopyInto(out *Revision) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Revision.
func (in *Revision) DeepCopy() *Revision {
	if in == nil {
		return nil
	}
	out := new(Revision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteImpl) DeepCopyInto(out *RouteImpl) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]v1alpha2.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]v1alpha2.HTTPRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteImpl.
func (in *RouteImpl) DeepCopy() *RouteImpl {
	if in == nil {
		return nil
	}
	out := new(RouteImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]v1alpha2.Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]v1alpha2.HTTPPathMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleOptions) DeepCopyInto(out *ScaleOptions) {
	*out = *in
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Keda != nil {
		in, out := &in.Keda, &out.Keda
		*out = new(KedaScaleOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Knative != nil {
		in, out := &in.Knative, &out.Knative
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleOptions.
func (in *ScaleOptions) DeepCopy() *ScaleOptions {
	if in == nil {
		return nil
	}
	out := new(ScaleOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceImpl) DeepCopyInto(out *ServiceImpl) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceImpl.
func (in *ServiceImpl) DeepCopy() *ServiceImpl {
	if in == nil {
		return nil
	}
	out := new(ServiceImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Serving) DeepCopyInto(out *Serving) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Serving.
func (in *Serving) DeepCopy() *Serving {
	if in == nil {
		return nil
	}
	out := new(Serving)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Serving) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingImpl) DeepCopyInto(out *ServingImpl) {
	*out = *in
	if in.ScaleOptions != nil {
		in, out := &in.ScaleOptions, &out.ScaleOptions
		*out = new(ScaleOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]*DaprIO, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DaprIO)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]*DaprIO, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DaprIO)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make(map[string]*v1alpha1.ComponentSpec, len(*in))
		for key, val := range *in {
			var outVal *v1alpha1.ComponentSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1alpha1.ComponentSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Pubsub != nil {
		in, out := &in.Pubsub, &out.Pubsub
		*out = make(map[string]*v1alpha1.ComponentSpec, len(*in))
		for key, val := range *in {
			var outVal *v1alpha1.ComponentSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1alpha1.ComponentSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]Triggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingImpl.
func (in *ServingImpl) DeepCopy() *ServingImpl {
	if in == nil {
		return nil
	}
	out := new(ServingImpl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingList) DeepCopyInto(out *ServingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Serving, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingList.
func (in *ServingList) DeepCopy() *ServingList {
	if in == nil {
		return nil
	}
	out := new(ServingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingSpec) DeepCopyInto(out *ServingSpec) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ImageCredentials != nil {
		in, out := &in.ImageCredentials, &out.ImageCredentials
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]*DaprIO, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DaprIO)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]*DaprIO, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DaprIO)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ScaleOptions != nil {
		in, out := &in.ScaleOptions, &out.ScaleOptions
		*out = new(ScaleOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make(map[string]*v1alpha1.ComponentSpec, len(*in))
		for key, val := range *in {
			var outVal *v1alpha1.ComponentSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1alpha1.ComponentSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Pubsub != nil {
		in, out := &in.Pubsub, &out.Pubsub
		*out = make(map[string]*v1alpha1.ComponentSpec, len(*in))
		for key, val := range *in {
			var outVal *v1alpha1.ComponentSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1alpha1.ComponentSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]Triggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingSpec.
func (in *ServingSpec) DeepCopy() *ServingSpec {
	if in == nil {
		return nil
	}
	out := new(ServingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServingStatus) DeepCopyInto(out *ServingStatus) {
	*out = *in
	if in.ResourceRef != nil {
		in, out := &in.ResourceRef, &out.ResourceRef
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServingStatus.
func (in *ServingStatus) DeepCopy() *ServingStatus {
	if in == nil {
		return nil
	}
	out := new(ServingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipwrightEngine) DeepCopyInto(out *ShipwrightEngine) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(Strategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipwrightEngine.
func (in *ShipwrightEngine) DeepCopy() *ShipwrightEngine {
	if in == nil {
		return nil
	}
	out := new(ShipwrightEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResult) DeepCopyInto(out *SourceResult) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSourceResult)
		**out = **in
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleSourceResult)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResult.
func (in *SourceResult) DeepCopy() *SourceResult {
	if in == nil {
		return nil
	}
	out := new(SourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Triggers) DeepCopyInto(out *Triggers) {
	*out = *in
	in.ScaleTriggers.DeepCopyInto(&out.ScaleTriggers)
	if in.TargetKind != nil {
		in, out := &in.TargetKind, &out.TargetKind
		*out = new(ScaleTargetKind)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Triggers.
func (in *Triggers) DeepCopy() *Triggers {
	if in == nil {
		return nil
	}
	out := new(Triggers)
	in.DeepCopyInto(out)
	return out
}
