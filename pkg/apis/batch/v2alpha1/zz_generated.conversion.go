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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v2alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v2alpha1_Job_To_batch_Job,
		Convert_batch_Job_To_v2alpha1_Job,
		Convert_v2alpha1_JobCondition_To_batch_JobCondition,
		Convert_batch_JobCondition_To_v2alpha1_JobCondition,
		Convert_v2alpha1_JobList_To_batch_JobList,
		Convert_batch_JobList_To_v2alpha1_JobList,
		Convert_v2alpha1_JobSpec_To_batch_JobSpec,
		Convert_batch_JobSpec_To_v2alpha1_JobSpec,
		Convert_v2alpha1_JobStatus_To_batch_JobStatus,
		Convert_batch_JobStatus_To_v2alpha1_JobStatus,
		Convert_v2alpha1_JobTemplate_To_batch_JobTemplate,
		Convert_batch_JobTemplate_To_v2alpha1_JobTemplate,
		Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec,
		Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec,
		Convert_v2alpha1_LabelSelector_To_unversioned_LabelSelector,
		Convert_unversioned_LabelSelector_To_v2alpha1_LabelSelector,
		Convert_v2alpha1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement,
		Convert_unversioned_LabelSelectorRequirement_To_v2alpha1_LabelSelectorRequirement,
		Convert_v2alpha1_ScheduledJob_To_batch_ScheduledJob,
		Convert_batch_ScheduledJob_To_v2alpha1_ScheduledJob,
		Convert_v2alpha1_ScheduledJobList_To_batch_ScheduledJobList,
		Convert_batch_ScheduledJobList_To_v2alpha1_ScheduledJobList,
		Convert_v2alpha1_ScheduledJobSpec_To_batch_ScheduledJobSpec,
		Convert_batch_ScheduledJobSpec_To_v2alpha1_ScheduledJobSpec,
		Convert_v2alpha1_ScheduledJobStatus_To_batch_ScheduledJobStatus,
		Convert_batch_ScheduledJobStatus_To_v2alpha1_ScheduledJobStatus,
	)
}

func autoConvert_v2alpha1_Job_To_batch_Job(in *Job, out *batch.Job, s conversion.Scope) error {
	SetDefaults_Job(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobStatus_To_batch_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_Job_To_batch_Job(in *Job, out *batch.Job, s conversion.Scope) error {
	return autoConvert_v2alpha1_Job_To_batch_Job(in, out, s)
}

func autoConvert_batch_Job_To_v2alpha1_Job(in *batch.Job, out *Job, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobSpec_To_v2alpha1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_JobStatus_To_v2alpha1_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_Job_To_v2alpha1_Job(in *batch.Job, out *Job, s conversion.Scope) error {
	return autoConvert_batch_Job_To_v2alpha1_Job(in, out, s)
}

func autoConvert_v2alpha1_JobCondition_To_batch_JobCondition(in *JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	out.Type = batch.JobConditionType(in.Type)
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

func Convert_v2alpha1_JobCondition_To_batch_JobCondition(in *JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobCondition_To_batch_JobCondition(in, out, s)
}

func autoConvert_batch_JobCondition_To_v2alpha1_JobCondition(in *batch.JobCondition, out *JobCondition, s conversion.Scope) error {
	out.Type = JobConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
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

func Convert_batch_JobCondition_To_v2alpha1_JobCondition(in *batch.JobCondition, out *JobCondition, s conversion.Scope) error {
	return autoConvert_batch_JobCondition_To_v2alpha1_JobCondition(in, out, s)
}

func autoConvert_v2alpha1_JobList_To_batch_JobList(in *JobList, out *batch.JobList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.Job, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_Job_To_batch_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v2alpha1_JobList_To_batch_JobList(in *JobList, out *batch.JobList, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobList_To_batch_JobList(in, out, s)
}

func autoConvert_batch_JobList_To_v2alpha1_JobList(in *batch.JobList, out *JobList, s conversion.Scope) error {
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
			if err := Convert_batch_Job_To_v2alpha1_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_batch_JobList_To_v2alpha1_JobList(in *batch.JobList, out *JobList, s conversion.Scope) error {
	return autoConvert_batch_JobList_To_v2alpha1_JobList(in, out, s)
}

func autoConvert_v2alpha1_JobSpec_To_batch_JobSpec(in *JobSpec, out *batch.JobSpec, s conversion.Scope) error {
	out.Parallelism = in.Parallelism
	out.Completions = in.Completions
	out.ActiveDeadlineSeconds = in.ActiveDeadlineSeconds
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(unversioned.LabelSelector)
		if err := Convert_v2alpha1_LabelSelector_To_unversioned_LabelSelector(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	out.ManualSelector = in.ManualSelector
	if err := v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_batch_JobSpec_To_v2alpha1_JobSpec(in *batch.JobSpec, out *JobSpec, s conversion.Scope) error {
	out.Parallelism = in.Parallelism
	out.Completions = in.Completions
	out.ActiveDeadlineSeconds = in.ActiveDeadlineSeconds
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(LabelSelector)
		if err := Convert_unversioned_LabelSelector_To_v2alpha1_LabelSelector(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	out.ManualSelector = in.ManualSelector
	if err := v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2alpha1_JobStatus_To_batch_JobStatus(in *JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]batch.JobCondition, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_JobCondition_To_batch_JobCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	out.StartTime = in.StartTime
	out.CompletionTime = in.CompletionTime
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func Convert_v2alpha1_JobStatus_To_batch_JobStatus(in *JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobStatus_To_batch_JobStatus(in, out, s)
}

func autoConvert_batch_JobStatus_To_v2alpha1_JobStatus(in *batch.JobStatus, out *JobStatus, s conversion.Scope) error {
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(*in))
		for i := range *in {
			if err := Convert_batch_JobCondition_To_v2alpha1_JobCondition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	out.StartTime = in.StartTime
	out.CompletionTime = in.CompletionTime
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func Convert_batch_JobStatus_To_v2alpha1_JobStatus(in *batch.JobStatus, out *JobStatus, s conversion.Scope) error {
	return autoConvert_batch_JobStatus_To_v2alpha1_JobStatus(in, out, s)
}

func autoConvert_v2alpha1_JobTemplate_To_batch_JobTemplate(in *JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_JobTemplate_To_batch_JobTemplate(in *JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplate_To_batch_JobTemplate(in, out, s)
}

func autoConvert_batch_JobTemplate_To_v2alpha1_JobTemplate(in *batch.JobTemplate, out *JobTemplate, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_JobTemplate_To_v2alpha1_JobTemplate(in *batch.JobTemplate, out *JobTemplate, s conversion.Scope) error {
	return autoConvert_batch_JobTemplate_To_v2alpha1_JobTemplate(in, out, s)
}

func autoConvert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in *JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in *JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in, out, s)
}

func autoConvert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *batch.JobTemplateSpec, out *JobTemplateSpec, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobSpec_To_v2alpha1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *batch.JobTemplateSpec, out *JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in, out, s)
}

func autoConvert_v2alpha1_LabelSelector_To_unversioned_LabelSelector(in *LabelSelector, out *unversioned.LabelSelector, s conversion.Scope) error {
	out.MatchLabels = in.MatchLabels
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]unversioned.LabelSelectorRequirement, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func Convert_v2alpha1_LabelSelector_To_unversioned_LabelSelector(in *LabelSelector, out *unversioned.LabelSelector, s conversion.Scope) error {
	return autoConvert_v2alpha1_LabelSelector_To_unversioned_LabelSelector(in, out, s)
}

func autoConvert_unversioned_LabelSelector_To_v2alpha1_LabelSelector(in *unversioned.LabelSelector, out *LabelSelector, s conversion.Scope) error {
	out.MatchLabels = in.MatchLabels
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]LabelSelectorRequirement, len(*in))
		for i := range *in {
			if err := Convert_unversioned_LabelSelectorRequirement_To_v2alpha1_LabelSelectorRequirement(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func Convert_unversioned_LabelSelector_To_v2alpha1_LabelSelector(in *unversioned.LabelSelector, out *LabelSelector, s conversion.Scope) error {
	return autoConvert_unversioned_LabelSelector_To_v2alpha1_LabelSelector(in, out, s)
}

func autoConvert_v2alpha1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in *LabelSelectorRequirement, out *unversioned.LabelSelectorRequirement, s conversion.Scope) error {
	out.Key = in.Key
	out.Operator = unversioned.LabelSelectorOperator(in.Operator)
	out.Values = in.Values
	return nil
}

func Convert_v2alpha1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in *LabelSelectorRequirement, out *unversioned.LabelSelectorRequirement, s conversion.Scope) error {
	return autoConvert_v2alpha1_LabelSelectorRequirement_To_unversioned_LabelSelectorRequirement(in, out, s)
}

func autoConvert_unversioned_LabelSelectorRequirement_To_v2alpha1_LabelSelectorRequirement(in *unversioned.LabelSelectorRequirement, out *LabelSelectorRequirement, s conversion.Scope) error {
	out.Key = in.Key
	out.Operator = LabelSelectorOperator(in.Operator)
	out.Values = in.Values
	return nil
}

func Convert_unversioned_LabelSelectorRequirement_To_v2alpha1_LabelSelectorRequirement(in *unversioned.LabelSelectorRequirement, out *LabelSelectorRequirement, s conversion.Scope) error {
	return autoConvert_unversioned_LabelSelectorRequirement_To_v2alpha1_LabelSelectorRequirement(in, out, s)
}

func autoConvert_v2alpha1_ScheduledJob_To_batch_ScheduledJob(in *ScheduledJob, out *batch.ScheduledJob, s conversion.Scope) error {
	SetDefaults_ScheduledJob(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_ScheduledJobSpec_To_batch_ScheduledJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_ScheduledJobStatus_To_batch_ScheduledJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_ScheduledJob_To_batch_ScheduledJob(in *ScheduledJob, out *batch.ScheduledJob, s conversion.Scope) error {
	return autoConvert_v2alpha1_ScheduledJob_To_batch_ScheduledJob(in, out, s)
}

func autoConvert_batch_ScheduledJob_To_v2alpha1_ScheduledJob(in *batch.ScheduledJob, out *ScheduledJob, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_ScheduledJobSpec_To_v2alpha1_ScheduledJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_ScheduledJobStatus_To_v2alpha1_ScheduledJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_ScheduledJob_To_v2alpha1_ScheduledJob(in *batch.ScheduledJob, out *ScheduledJob, s conversion.Scope) error {
	return autoConvert_batch_ScheduledJob_To_v2alpha1_ScheduledJob(in, out, s)
}

func autoConvert_v2alpha1_ScheduledJobList_To_batch_ScheduledJobList(in *ScheduledJobList, out *batch.ScheduledJobList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.ScheduledJob, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_ScheduledJob_To_batch_ScheduledJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v2alpha1_ScheduledJobList_To_batch_ScheduledJobList(in *ScheduledJobList, out *batch.ScheduledJobList, s conversion.Scope) error {
	return autoConvert_v2alpha1_ScheduledJobList_To_batch_ScheduledJobList(in, out, s)
}

func autoConvert_batch_ScheduledJobList_To_v2alpha1_ScheduledJobList(in *batch.ScheduledJobList, out *ScheduledJobList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledJob, len(*in))
		for i := range *in {
			if err := Convert_batch_ScheduledJob_To_v2alpha1_ScheduledJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_batch_ScheduledJobList_To_v2alpha1_ScheduledJobList(in *batch.ScheduledJobList, out *ScheduledJobList, s conversion.Scope) error {
	return autoConvert_batch_ScheduledJobList_To_v2alpha1_ScheduledJobList(in, out, s)
}

func autoConvert_v2alpha1_ScheduledJobSpec_To_batch_ScheduledJobSpec(in *ScheduledJobSpec, out *batch.ScheduledJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = in.StartingDeadlineSeconds
	out.ConcurrencyPolicy = batch.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = in.Suspend
	if err := Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_ScheduledJobSpec_To_batch_ScheduledJobSpec(in *ScheduledJobSpec, out *batch.ScheduledJobSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_ScheduledJobSpec_To_batch_ScheduledJobSpec(in, out, s)
}

func autoConvert_batch_ScheduledJobSpec_To_v2alpha1_ScheduledJobSpec(in *batch.ScheduledJobSpec, out *ScheduledJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = in.StartingDeadlineSeconds
	out.ConcurrencyPolicy = ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = in.Suspend
	if err := Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_ScheduledJobSpec_To_v2alpha1_ScheduledJobSpec(in *batch.ScheduledJobSpec, out *ScheduledJobSpec, s conversion.Scope) error {
	return autoConvert_batch_ScheduledJobSpec_To_v2alpha1_ScheduledJobSpec(in, out, s)
}

func autoConvert_v2alpha1_ScheduledJobStatus_To_batch_ScheduledJobStatus(in *ScheduledJobStatus, out *batch.ScheduledJobStatus, s conversion.Scope) error {
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]api.ObjectReference, len(*in))
		for i := range *in {
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Active = nil
	}
	out.LastScheduleTime = in.LastScheduleTime
	return nil
}

func Convert_v2alpha1_ScheduledJobStatus_To_batch_ScheduledJobStatus(in *ScheduledJobStatus, out *batch.ScheduledJobStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_ScheduledJobStatus_To_batch_ScheduledJobStatus(in, out, s)
}

func autoConvert_batch_ScheduledJobStatus_To_v2alpha1_ScheduledJobStatus(in *batch.ScheduledJobStatus, out *ScheduledJobStatus, s conversion.Scope) error {
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]v1.ObjectReference, len(*in))
		for i := range *in {
			// TODO: Inefficient conversion - can we improve it?
			if err := s.Convert(&(*in)[i], &(*out)[i], 0); err != nil {
				return err
			}
		}
	} else {
		out.Active = nil
	}
	out.LastScheduleTime = in.LastScheduleTime
	return nil
}

func Convert_batch_ScheduledJobStatus_To_v2alpha1_ScheduledJobStatus(in *batch.ScheduledJobStatus, out *ScheduledJobStatus, s conversion.Scope) error {
	return autoConvert_batch_ScheduledJobStatus_To_v2alpha1_ScheduledJobStatus(in, out, s)
}
