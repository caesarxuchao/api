// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_Job_To_extensions_Job,
		Convert_extensions_Job_To_v1_Job,
		Convert_v1_JobCondition_To_extensions_JobCondition,
		Convert_extensions_JobCondition_To_v1_JobCondition,
		Convert_v1_JobList_To_extensions_JobList,
		Convert_extensions_JobList_To_v1_JobList,
		Convert_v1_JobSpec_To_extensions_JobSpec,
		Convert_extensions_JobSpec_To_v1_JobSpec,
		Convert_v1_JobStatus_To_extensions_JobStatus,
		Convert_extensions_JobStatus_To_v1_JobStatus,
		Convert_v1_LabelSelector_To_unversioned_LabelSelector,
		Convert_unversioned_LabelSelector_To_v1_LabelSelector,
		Convert_v1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement,
		Convert_unversioned_LabelSelectorRequirement_To_v1_LabelSelectorRequirement,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_Job_To_extensions_Job(in *Job, out *extensions.Job, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*Job))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1_JobSpec_To_extensions_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_JobStatus_To_extensions_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Job_To_extensions_Job(in *Job, out *extensions.Job, s conversion.Scope) error {
	return autoConvert_v1_Job_To_extensions_Job(in, out, s)
}

func autoConvert_extensions_Job_To_v1_Job(in *extensions.Job, out *Job, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.Job))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_extensions_JobSpec_To_v1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_extensions_JobStatus_To_v1_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_Job_To_v1_Job(in *extensions.Job, out *Job, s conversion.Scope) error {
	return autoConvert_extensions_Job_To_v1_Job(in, out, s)
}

func autoConvert_v1_JobCondition_To_extensions_JobCondition(in *JobCondition, out *extensions.JobCondition, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*JobCondition))(in)
	}
	out.Type = extensions.JobConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastProbeTime, &out.LastProbeTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastTransitionTime, &out.LastTransitionTime, s); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1_JobCondition_To_extensions_JobCondition(in *JobCondition, out *extensions.JobCondition, s conversion.Scope) error {
	return autoConvert_v1_JobCondition_To_extensions_JobCondition(in, out, s)
}

func autoConvert_extensions_JobCondition_To_v1_JobCondition(in *extensions.JobCondition, out *JobCondition, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.JobCondition))(in)
	}
	out.Type = JobConditionType(in.Type)
	out.Status = api_v1.ConditionStatus(in.Status)
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastProbeTime, &out.LastProbeTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastTransitionTime, &out.LastTransitionTime, s); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_extensions_JobCondition_To_v1_JobCondition(in *extensions.JobCondition, out *JobCondition, s conversion.Scope) error {
	return autoConvert_extensions_JobCondition_To_v1_JobCondition(in, out, s)
}

func autoConvert_v1_JobList_To_extensions_JobList(in *JobList, out *extensions.JobList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*JobList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]extensions.Job, len(*in))
		for i := range *in {
			if err := Convert_v1_Job_To_extensions_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_JobList_To_extensions_JobList(in *JobList, out *extensions.JobList, s conversion.Scope) error {
	return autoConvert_v1_JobList_To_extensions_JobList(in, out, s)
}

func autoConvert_extensions_JobList_To_v1_JobList(in *extensions.JobList, out *JobList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.JobList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Job, len(*in))
		for i := range *in {
			if err := Convert_extensions_Job_To_v1_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_extensions_JobList_To_v1_JobList(in *extensions.JobList, out *JobList, s conversion.Scope) error {
	return autoConvert_extensions_JobList_To_v1_JobList(in, out, s)
}

func autoConvert_v1_JobSpec_To_extensions_JobSpec(in *JobSpec, out *extensions.JobSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*JobSpec))(in)
	}
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int)
		**out = int(**in)
	} else {
		out.Parallelism = nil
	}
	if in.Completions != nil {
		in, out := &in.Completions, &out.Completions
		*out = new(int)
		**out = int(**in)
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
		*out = new(unversioned.LabelSelector)
		if err := Convert_v1_LabelSelector_To_unversioned_LabelSelector(*in, *out, s); err != nil {
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
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Template, &out.Template, 0); err != nil {
		return err
	}
	return nil
}

func Convert_v1_JobSpec_To_extensions_JobSpec(in *JobSpec, out *extensions.JobSpec, s conversion.Scope) error {
	return autoConvert_v1_JobSpec_To_extensions_JobSpec(in, out, s)
}

func autoConvert_extensions_JobSpec_To_v1_JobSpec(in *extensions.JobSpec, out *JobSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.JobSpec))(in)
	}
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = int32(**in)
	} else {
		out.Parallelism = nil
	}
	if in.Completions != nil {
		in, out := &in.Completions, &out.Completions
		*out = new(int32)
		**out = int32(**in)
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
		if err := Convert_unversioned_LabelSelector_To_v1_LabelSelector(*in, *out, s); err != nil {
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
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.Template, &out.Template, 0); err != nil {
		return err
	}
	return nil
}

func Convert_extensions_JobSpec_To_v1_JobSpec(in *extensions.JobSpec, out *JobSpec, s conversion.Scope) error {
	return autoConvert_extensions_JobSpec_To_v1_JobSpec(in, out, s)
}

func autoConvert_v1_JobStatus_To_extensions_JobStatus(in *JobStatus, out *extensions.JobStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*JobStatus))(in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]extensions.JobCondition, len(*in))
		for i := range *in {
			if err := Convert_v1_JobCondition_To_extensions_JobCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.StartTime = nil
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CompletionTime = nil
	}
	out.Active = int(in.Active)
	out.Succeeded = int(in.Succeeded)
	out.Failed = int(in.Failed)
	return nil
}

func Convert_v1_JobStatus_To_extensions_JobStatus(in *JobStatus, out *extensions.JobStatus, s conversion.Scope) error {
	return autoConvert_v1_JobStatus_To_extensions_JobStatus(in, out, s)
}

func autoConvert_extensions_JobStatus_To_v1_JobStatus(in *extensions.JobStatus, out *JobStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.JobStatus))(in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(*in))
		for i := range *in {
			if err := Convert_extensions_JobCondition_To_v1_JobCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.StartTime = nil
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CompletionTime = nil
	}
	out.Active = int32(in.Active)
	out.Succeeded = int32(in.Succeeded)
	out.Failed = int32(in.Failed)
	return nil
}

func Convert_extensions_JobStatus_To_v1_JobStatus(in *extensions.JobStatus, out *JobStatus, s conversion.Scope) error {
	return autoConvert_extensions_JobStatus_To_v1_JobStatus(in, out, s)
}

func autoConvert_v1_LabelSelector_To_unversioned_LabelSelector(in *LabelSelector, out *unversioned.LabelSelector, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*LabelSelector))(in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.MatchLabels = nil
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]unversioned.LabelSelectorRequirement, len(*in))
		for i := range *in {
			if err := Convert_v1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func Convert_v1_LabelSelector_To_unversioned_LabelSelector(in *LabelSelector, out *unversioned.LabelSelector, s conversion.Scope) error {
	return autoConvert_v1_LabelSelector_To_unversioned_LabelSelector(in, out, s)
}

func autoConvert_unversioned_LabelSelector_To_v1_LabelSelector(in *unversioned.LabelSelector, out *LabelSelector, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*unversioned.LabelSelector))(in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
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
			if err := Convert_unversioned_LabelSelectorRequirement_To_v1_LabelSelectorRequirement(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func Convert_unversioned_LabelSelector_To_v1_LabelSelector(in *unversioned.LabelSelector, out *LabelSelector, s conversion.Scope) error {
	return autoConvert_unversioned_LabelSelector_To_v1_LabelSelector(in, out, s)
}

func autoConvert_v1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in *LabelSelectorRequirement, out *unversioned.LabelSelectorRequirement, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*LabelSelectorRequirement))(in)
	}
	out.Key = in.Key
	out.Operator = unversioned.LabelSelectorOperator(in.Operator)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Values = nil
	}
	return nil
}

func Convert_v1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in *LabelSelectorRequirement, out *unversioned.LabelSelectorRequirement, s conversion.Scope) error {
	return autoConvert_v1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in, out, s)
}

func autoConvert_unversioned_LabelSelectorRequirement_To_v1_LabelSelectorRequirement(in *unversioned.LabelSelectorRequirement, out *LabelSelectorRequirement, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*unversioned.LabelSelectorRequirement))(in)
	}
	out.Key = in.Key
	out.Operator = LabelSelectorOperator(in.Operator)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Values = nil
	}
	return nil
}

func Convert_unversioned_LabelSelectorRequirement_To_v1_LabelSelectorRequirement(in *unversioned.LabelSelectorRequirement, out *LabelSelectorRequirement, s conversion.Scope) error {
	return autoConvert_unversioned_LabelSelectorRequirement_To_v1_LabelSelectorRequirement(in, out, s)
}
