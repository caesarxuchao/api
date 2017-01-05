// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_v1_APIGroup, InType: reflect.TypeOf(&APIGroup{})},
		{Fn: DeepCopy_v1_APIGroupList, InType: reflect.TypeOf(&APIGroupList{})},
		{Fn: DeepCopy_v1_APIResource, InType: reflect.TypeOf(&APIResource{})},
		{Fn: DeepCopy_v1_APIResourceList, InType: reflect.TypeOf(&APIResourceList{})},
		{Fn: DeepCopy_v1_APIVersions, InType: reflect.TypeOf(&APIVersions{})},
		{Fn: DeepCopy_v1_Duration, InType: reflect.TypeOf(&Duration{})},
		{Fn: DeepCopy_v1_ExportOptions, InType: reflect.TypeOf(&ExportOptions{})},
		{Fn: DeepCopy_v1_GetOptions, InType: reflect.TypeOf(&GetOptions{})},
		{Fn: DeepCopy_v1_GroupKind, InType: reflect.TypeOf(&GroupKind{})},
		{Fn: DeepCopy_v1_GroupResource, InType: reflect.TypeOf(&GroupResource{})},
		{Fn: DeepCopy_v1_GroupVersion, InType: reflect.TypeOf(&GroupVersion{})},
		{Fn: DeepCopy_v1_GroupVersionForDiscovery, InType: reflect.TypeOf(&GroupVersionForDiscovery{})},
		{Fn: DeepCopy_v1_GroupVersionKind, InType: reflect.TypeOf(&GroupVersionKind{})},
		{Fn: DeepCopy_v1_GroupVersionResource, InType: reflect.TypeOf(&GroupVersionResource{})},
		{Fn: DeepCopy_v1_LabelSelector, InType: reflect.TypeOf(&LabelSelector{})},
		{Fn: DeepCopy_v1_LabelSelectorRequirement, InType: reflect.TypeOf(&LabelSelectorRequirement{})},
		{Fn: DeepCopy_v1_ListMeta, InType: reflect.TypeOf(&ListMeta{})},
		{Fn: DeepCopy_v1_OwnerReference, InType: reflect.TypeOf(&OwnerReference{})},
		{Fn: DeepCopy_v1_Patch, InType: reflect.TypeOf(&Patch{})},
		{Fn: DeepCopy_v1_RootPaths, InType: reflect.TypeOf(&RootPaths{})},
		{Fn: DeepCopy_v1_ServerAddressByClientCIDR, InType: reflect.TypeOf(&ServerAddressByClientCIDR{})},
		{Fn: DeepCopy_v1_Status, InType: reflect.TypeOf(&Status{})},
		{Fn: DeepCopy_v1_StatusCause, InType: reflect.TypeOf(&StatusCause{})},
		{Fn: DeepCopy_v1_StatusDetails, InType: reflect.TypeOf(&StatusDetails{})},
		{Fn: DeepCopy_v1_Time, InType: reflect.TypeOf(&Time{})},
		{Fn: DeepCopy_v1_Timestamp, InType: reflect.TypeOf(&Timestamp{})},
		{Fn: DeepCopy_v1_TypeMeta, InType: reflect.TypeOf(&TypeMeta{})},
	}
}

func DeepCopy_v1_APIGroup(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIGroup)
		out := out.(*APIGroup)
		*out = *in
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]GroupVersionForDiscovery, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		}
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		}
		return nil
	}
}

func DeepCopy_v1_APIGroupList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIGroupList)
		out := out.(*APIGroupList)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]APIGroup, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_APIGroup(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_APIResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIResource)
		out := out.(*APIResource)
		*out = *in
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make(Verbs, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_APIResourceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIResourceList)
		out := out.(*APIResourceList)
		*out = *in
		if in.APIResources != nil {
			in, out := &in.APIResources, &out.APIResources
			*out = make([]APIResource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_APIResource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_APIVersions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIVersions)
		out := out.(*APIVersions)
		*out = *in
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		}
		return nil
	}
}

func DeepCopy_v1_Duration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Duration)
		out := out.(*Duration)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_ExportOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExportOptions)
		out := out.(*ExportOptions)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GetOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GetOptions)
		out := out.(*GetOptions)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupKind(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupKind)
		out := out.(*GroupKind)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupResource)
		out := out.(*GroupResource)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupVersion(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersion)
		out := out.(*GroupVersion)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupVersionForDiscovery(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionForDiscovery)
		out := out.(*GroupVersionForDiscovery)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupVersionKind(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionKind)
		out := out.(*GroupVersionKind)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_GroupVersionResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionResource)
		out := out.(*GroupVersionResource)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_LabelSelector(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelector)
		out := out.(*LabelSelector)
		*out = *in
		if in.MatchLabels != nil {
			in, out := &in.MatchLabels, &out.MatchLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.MatchExpressions != nil {
			in, out := &in.MatchExpressions, &out.MatchExpressions
			*out = make([]LabelSelectorRequirement, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_LabelSelectorRequirement(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1_LabelSelectorRequirement(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelectorRequirement)
		out := out.(*LabelSelectorRequirement)
		*out = *in
		if in.Values != nil {
			in, out := &in.Values, &out.Values
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_ListMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ListMeta)
		out := out.(*ListMeta)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_OwnerReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*OwnerReference)
		out := out.(*OwnerReference)
		*out = *in
		if in.Controller != nil {
			in, out := &in.Controller, &out.Controller
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1_Patch(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Patch)
		out := out.(*Patch)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_RootPaths(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RootPaths)
		out := out.(*RootPaths)
		*out = *in
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1_ServerAddressByClientCIDR(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServerAddressByClientCIDR)
		out := out.(*ServerAddressByClientCIDR)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_Status(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Status)
		out := out.(*Status)
		*out = *in
		if in.Details != nil {
			in, out := &in.Details, &out.Details
			*out = new(StatusDetails)
			if err := DeepCopy_v1_StatusDetails(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1_StatusCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatusCause)
		out := out.(*StatusCause)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_StatusDetails(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatusDetails)
		out := out.(*StatusDetails)
		*out = *in
		if in.Causes != nil {
			in, out := &in.Causes, &out.Causes
			*out = make([]StatusCause, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		}
		return nil
	}
}

func DeepCopy_v1_Time(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Time)
		out := out.(*Time)
		*out = in.DeepCopy()
		return nil
	}
}

func DeepCopy_v1_Timestamp(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Timestamp)
		out := out.(*Timestamp)
		*out = *in
		return nil
	}
}

func DeepCopy_v1_TypeMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TypeMeta)
		out := out.(*TypeMeta)
		*out = *in
		return nil
	}
}
