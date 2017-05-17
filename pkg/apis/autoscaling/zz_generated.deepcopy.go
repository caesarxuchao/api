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

package autoscaling

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_CrossVersionObjectReference, InType: reflect.TypeOf(&CrossVersionObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_HorizontalPodAutoscaler, InType: reflect.TypeOf(&HorizontalPodAutoscaler{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_HorizontalPodAutoscalerList, InType: reflect.TypeOf(&HorizontalPodAutoscalerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_HorizontalPodAutoscalerSpec, InType: reflect.TypeOf(&HorizontalPodAutoscalerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_HorizontalPodAutoscalerStatus, InType: reflect.TypeOf(&HorizontalPodAutoscalerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_MetricSpec, InType: reflect.TypeOf(&MetricSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_MetricStatus, InType: reflect.TypeOf(&MetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ObjectMetricSource, InType: reflect.TypeOf(&ObjectMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ObjectMetricStatus, InType: reflect.TypeOf(&ObjectMetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_PodsMetricSource, InType: reflect.TypeOf(&PodsMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_PodsMetricStatus, InType: reflect.TypeOf(&PodsMetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ResourceMetricSource, InType: reflect.TypeOf(&ResourceMetricSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ResourceMetricStatus, InType: reflect.TypeOf(&ResourceMetricStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_Scale, InType: reflect.TypeOf(&Scale{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ScaleSpec, InType: reflect.TypeOf(&ScaleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_autoscaling_ScaleStatus, InType: reflect.TypeOf(&ScaleStatus{})},
	)
}

// DeepCopy_autoscaling_CrossVersionObjectReference is an autogenerated deepcopy function.
func DeepCopy_autoscaling_CrossVersionObjectReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CrossVersionObjectReference)
		out := out.(*CrossVersionObjectReference)
		*out = *in
		return nil
	}
}

// DeepCopy_autoscaling_HorizontalPodAutoscaler is an autogenerated deepcopy function.
func DeepCopy_autoscaling_HorizontalPodAutoscaler(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscaler)
		out := out.(*HorizontalPodAutoscaler)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_autoscaling_HorizontalPodAutoscalerList is an autogenerated deepcopy function.
func DeepCopy_autoscaling_HorizontalPodAutoscalerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerList)
		out := out.(*HorizontalPodAutoscalerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]HorizontalPodAutoscaler, len(*in))
			for i := range *in {
				if err := DeepCopy_autoscaling_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_autoscaling_HorizontalPodAutoscalerSpec is an autogenerated deepcopy function.
func DeepCopy_autoscaling_HorizontalPodAutoscalerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerSpec)
		out := out.(*HorizontalPodAutoscalerSpec)
		*out = *in
		if in.MinReplicas != nil {
			in, out := &in.MinReplicas, &out.MinReplicas
			*out = new(int32)
			**out = **in
		}
		if in.Metrics != nil {
			in, out := &in.Metrics, &out.Metrics
			*out = make([]MetricSpec, len(*in))
			for i := range *in {
				if err := DeepCopy_autoscaling_MetricSpec(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_autoscaling_HorizontalPodAutoscalerStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_HorizontalPodAutoscalerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerStatus)
		out := out.(*HorizontalPodAutoscalerStatus)
		*out = *in
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		}
		if in.LastScaleTime != nil {
			in, out := &in.LastScaleTime, &out.LastScaleTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.CurrentMetrics != nil {
			in, out := &in.CurrentMetrics, &out.CurrentMetrics
			*out = make([]MetricStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_autoscaling_MetricStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_autoscaling_MetricSpec is an autogenerated deepcopy function.
func DeepCopy_autoscaling_MetricSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*MetricSpec)
		out := out.(*MetricSpec)
		*out = *in
		if in.Object != nil {
			in, out := &in.Object, &out.Object
			*out = new(ObjectMetricSource)
			if err := DeepCopy_autoscaling_ObjectMetricSource(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Pods != nil {
			in, out := &in.Pods, &out.Pods
			*out = new(PodsMetricSource)
			if err := DeepCopy_autoscaling_PodsMetricSource(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Resource != nil {
			in, out := &in.Resource, &out.Resource
			*out = new(ResourceMetricSource)
			if err := DeepCopy_autoscaling_ResourceMetricSource(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_autoscaling_MetricStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_MetricStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*MetricStatus)
		out := out.(*MetricStatus)
		*out = *in
		if in.Object != nil {
			in, out := &in.Object, &out.Object
			*out = new(ObjectMetricStatus)
			if err := DeepCopy_autoscaling_ObjectMetricStatus(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Pods != nil {
			in, out := &in.Pods, &out.Pods
			*out = new(PodsMetricStatus)
			if err := DeepCopy_autoscaling_PodsMetricStatus(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Resource != nil {
			in, out := &in.Resource, &out.Resource
			*out = new(ResourceMetricStatus)
			if err := DeepCopy_autoscaling_ResourceMetricStatus(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_autoscaling_ObjectMetricSource is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ObjectMetricSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectMetricSource)
		out := out.(*ObjectMetricSource)
		*out = *in
		out.TargetValue = in.TargetValue.DeepCopy()
		return nil
	}
}

// DeepCopy_autoscaling_ObjectMetricStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ObjectMetricStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectMetricStatus)
		out := out.(*ObjectMetricStatus)
		*out = *in
		out.CurrentValue = in.CurrentValue.DeepCopy()
		return nil
	}
}

// DeepCopy_autoscaling_PodsMetricSource is an autogenerated deepcopy function.
func DeepCopy_autoscaling_PodsMetricSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodsMetricSource)
		out := out.(*PodsMetricSource)
		*out = *in
		out.TargetAverageValue = in.TargetAverageValue.DeepCopy()
		return nil
	}
}

// DeepCopy_autoscaling_PodsMetricStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_PodsMetricStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodsMetricStatus)
		out := out.(*PodsMetricStatus)
		*out = *in
		out.CurrentAverageValue = in.CurrentAverageValue.DeepCopy()
		return nil
	}
}

// DeepCopy_autoscaling_ResourceMetricSource is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ResourceMetricSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceMetricSource)
		out := out.(*ResourceMetricSource)
		*out = *in
		if in.TargetAverageUtilization != nil {
			in, out := &in.TargetAverageUtilization, &out.TargetAverageUtilization
			*out = new(int32)
			**out = **in
		}
		if in.TargetAverageValue != nil {
			in, out := &in.TargetAverageValue, &out.TargetAverageValue
			*out = new(resource.Quantity)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

// DeepCopy_autoscaling_ResourceMetricStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ResourceMetricStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceMetricStatus)
		out := out.(*ResourceMetricStatus)
		*out = *in
		if in.CurrentAverageUtilization != nil {
			in, out := &in.CurrentAverageUtilization, &out.CurrentAverageUtilization
			*out = new(int32)
			**out = **in
		}
		out.CurrentAverageValue = in.CurrentAverageValue.DeepCopy()
		return nil
	}
}

// DeepCopy_autoscaling_Scale is an autogenerated deepcopy function.
func DeepCopy_autoscaling_Scale(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Scale)
		out := out.(*Scale)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_autoscaling_ScaleSpec is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ScaleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleSpec)
		out := out.(*ScaleSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_autoscaling_ScaleStatus is an autogenerated deepcopy function.
func DeepCopy_autoscaling_ScaleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleStatus)
		out := out.(*ScaleStatus)
		*out = *in
		return nil
	}
}
