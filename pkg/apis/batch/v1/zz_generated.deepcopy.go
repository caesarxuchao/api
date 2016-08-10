// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Job, InType: reflect.TypeOf(&Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobCondition, InType: reflect.TypeOf(&JobCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobList, InType: reflect.TypeOf(&JobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobSpec, InType: reflect.TypeOf(&JobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JobStatus, InType: reflect.TypeOf(&JobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LabelSelector, InType: reflect.TypeOf(&LabelSelector{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LabelSelectorRequirement, InType: reflect.TypeOf(&LabelSelectorRequirement{})},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1_Job(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Job)
		out := out.(*Job)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_JobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_JobCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobCondition)
		out := out.(*JobCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1_JobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobList)
		out := out.(*JobList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Job, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Job(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_JobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobSpec)
		out := out.(*JobSpec)
		if in.Parallelism != nil {
			in, out := &in.Parallelism, &out.Parallelism
			*out = new(int32)
			**out = **in
		} else {
			out.Parallelism = nil
		}
		if in.Completions != nil {
			in, out := &in.Completions, &out.Completions
			*out = new(int32)
			**out = **in
		} else {
			out.Completions = nil
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.ActiveDeadlineSeconds = nil
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(LabelSelector)
			if err := DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if in.ManualSelector != nil {
			in, out := &in.ManualSelector, &out.ManualSelector
			*out = new(bool)
			**out = **in
		} else {
			out.ManualSelector = nil
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_JobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobStatus)
		out := out.(*JobStatus)
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]JobCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_JobCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if in.StartTime != nil {
			in, out := &in.StartTime, &out.StartTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.StartTime = nil
		}
		if in.CompletionTime != nil {
			in, out := &in.CompletionTime, &out.CompletionTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.CompletionTime = nil
		}
		out.Active = in.Active
		out.Succeeded = in.Succeeded
		out.Failed = in.Failed
		return nil
	}
}

func DeepCopy_v1_LabelSelector(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelector)
		out := out.(*LabelSelector)
		if in.MatchLabels != nil {
			in, out := &in.MatchLabels, &out.MatchLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.MatchLabels = nil
		}
		if in.MatchExpressions != nil {
			in, out := &in.MatchExpressions, &out.MatchExpressions
			*out = make([]LabelSelectorRequirement, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_LabelSelectorRequirement(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.MatchExpressions = nil
		}
		return nil
	}
}

func DeepCopy_v1_LabelSelectorRequirement(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelectorRequirement)
		out := out.(*LabelSelectorRequirement)
		out.Key = in.Key
		out.Operator = in.Operator
		if in.Values != nil {
			in, out := &in.Values, &out.Values
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Values = nil
		}
		return nil
	}
}
