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

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_AdmissionConfiguration, InType: reflect.TypeOf(&AdmissionConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_AdmissionPluginConfiguration, InType: reflect.TypeOf(&AdmissionPluginConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeProxyConfiguration, InType: reflect.TypeOf(&KubeProxyConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeSchedulerConfiguration, InType: reflect.TypeOf(&KubeSchedulerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletAnonymousAuthentication, InType: reflect.TypeOf(&KubeletAnonymousAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletAuthentication, InType: reflect.TypeOf(&KubeletAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletAuthorization, InType: reflect.TypeOf(&KubeletAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletConfiguration, InType: reflect.TypeOf(&KubeletConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletWebhookAuthentication, InType: reflect.TypeOf(&KubeletWebhookAuthentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletWebhookAuthorization, InType: reflect.TypeOf(&KubeletWebhookAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_KubeletX509Authentication, InType: reflect.TypeOf(&KubeletX509Authentication{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_LeaderElectionConfiguration, InType: reflect.TypeOf(&LeaderElectionConfiguration{})},
	)
}

func DeepCopy_v1alpha1_AdmissionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionConfiguration)
		out := out.(*AdmissionConfiguration)
		*out = *in
		if in.Plugins != nil {
			in, out := &in.Plugins, &out.Plugins
			*out = make([]AdmissionPluginConfiguration, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_AdmissionPluginConfiguration(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_AdmissionPluginConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AdmissionPluginConfiguration)
		out := out.(*AdmissionPluginConfiguration)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Configuration); err != nil {
			return err
		} else {
			out.Configuration = *newVal.(*runtime.RawExtension)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeProxyConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeProxyConfiguration)
		out := out.(*KubeProxyConfiguration)
		*out = *in
		if in.IPTablesMasqueradeBit != nil {
			in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
			*out = new(int32)
			**out = **in
		}
		if in.OOMScoreAdj != nil {
			in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeSchedulerConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeSchedulerConfiguration)
		out := out.(*KubeSchedulerConfiguration)
		*out = *in
		if in.EnableProfiling != nil {
			in, out := &in.EnableProfiling, &out.EnableProfiling
			*out = new(bool)
			**out = **in
		}
		if err := DeepCopy_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletAnonymousAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAnonymousAuthentication)
		out := out.(*KubeletAnonymousAuthentication)
		*out = *in
		if in.Enabled != nil {
			in, out := &in.Enabled, &out.Enabled
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthentication)
		out := out.(*KubeletAuthentication)
		*out = *in
		if err := DeepCopy_v1alpha1_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletAuthorization)
		out := out.(*KubeletAuthorization)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletConfiguration)
		out := out.(*KubeletConfiguration)
		*out = *in
		if in.EnableServer != nil {
			in, out := &in.EnableServer, &out.EnableServer
			*out = new(bool)
			**out = **in
		}
		if err := DeepCopy_v1alpha1_KubeletAuthentication(&in.Authentication, &out.Authentication, c); err != nil {
			return err
		}
		if in.AllowPrivileged != nil {
			in, out := &in.AllowPrivileged, &out.AllowPrivileged
			*out = new(bool)
			**out = **in
		}
		if in.HostNetworkSources != nil {
			in, out := &in.HostNetworkSources, &out.HostNetworkSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.HostPIDSources != nil {
			in, out := &in.HostPIDSources, &out.HostPIDSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.HostIPCSources != nil {
			in, out := &in.HostIPCSources, &out.HostIPCSources
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.RegistryPullQPS != nil {
			in, out := &in.RegistryPullQPS, &out.RegistryPullQPS
			*out = new(int32)
			**out = **in
		}
		if in.EventRecordQPS != nil {
			in, out := &in.EventRecordQPS, &out.EventRecordQPS
			*out = new(int32)
			**out = **in
		}
		if in.EnableDebuggingHandlers != nil {
			in, out := &in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers
			*out = new(bool)
			**out = **in
		}
		if in.MaxContainerCount != nil {
			in, out := &in.MaxContainerCount, &out.MaxContainerCount
			*out = new(int32)
			**out = **in
		}
		if in.OOMScoreAdj != nil {
			in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
			*out = new(int32)
			**out = **in
		}
		if in.RegisterNode != nil {
			in, out := &in.RegisterNode, &out.RegisterNode
			*out = new(bool)
			**out = **in
		}
		if in.ImageGCHighThresholdPercent != nil {
			in, out := &in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent
			*out = new(int32)
			**out = **in
		}
		if in.ImageGCLowThresholdPercent != nil {
			in, out := &in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent
			*out = new(int32)
			**out = **in
		}
		if in.ExperimentalCgroupsPerQOS != nil {
			in, out := &in.ExperimentalCgroupsPerQOS, &out.ExperimentalCgroupsPerQOS
			*out = new(bool)
			**out = **in
		}
		if in.LockFilePath != nil {
			in, out := &in.LockFilePath, &out.LockFilePath
			*out = new(string)
			**out = **in
		}
		if in.CPUCFSQuota != nil {
			in, out := &in.CPUCFSQuota, &out.CPUCFSQuota
			*out = new(bool)
			**out = **in
		}
		if in.Containerized != nil {
			in, out := &in.Containerized, &out.Containerized
			*out = new(bool)
			**out = **in
		}
		if in.RegisterSchedulable != nil {
			in, out := &in.RegisterSchedulable, &out.RegisterSchedulable
			*out = new(bool)
			**out = **in
		}
		if in.RegisterWithTaints != nil {
			in, out := &in.RegisterWithTaints, &out.RegisterWithTaints
			*out = make([]v1.Taint, len(*in))
			for i := range *in {
				if err := v1.DeepCopy_v1_Taint(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.KubeAPIQPS != nil {
			in, out := &in.KubeAPIQPS, &out.KubeAPIQPS
			*out = new(int32)
			**out = **in
		}
		if in.SerializeImagePulls != nil {
			in, out := &in.SerializeImagePulls, &out.SerializeImagePulls
			*out = new(bool)
			**out = **in
		}
		if in.NodeLabels != nil {
			in, out := &in.NodeLabels, &out.NodeLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.EvictionHard != nil {
			in, out := &in.EvictionHard, &out.EvictionHard
			*out = new(string)
			**out = **in
		}
		if in.ExperimentalKernelMemcgNotification != nil {
			in, out := &in.ExperimentalKernelMemcgNotification, &out.ExperimentalKernelMemcgNotification
			*out = new(bool)
			**out = **in
		}
		if in.EnableControllerAttachDetach != nil {
			in, out := &in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach
			*out = new(bool)
			**out = **in
		}
		if in.SystemReserved != nil {
			in, out := &in.SystemReserved, &out.SystemReserved
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.KubeReserved != nil {
			in, out := &in.KubeReserved, &out.KubeReserved
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.MakeIPTablesUtilChains != nil {
			in, out := &in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains
			*out = new(bool)
			**out = **in
		}
		if in.IPTablesMasqueradeBit != nil {
			in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
			*out = new(int32)
			**out = **in
		}
		if in.IPTablesDropBit != nil {
			in, out := &in.IPTablesDropBit, &out.IPTablesDropBit
			*out = new(int32)
			**out = **in
		}
		if in.AllowedUnsafeSysctls != nil {
			in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletWebhookAuthentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthentication)
		out := out.(*KubeletWebhookAuthentication)
		*out = *in
		if in.Enabled != nil {
			in, out := &in.Enabled, &out.Enabled
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletWebhookAuthorization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletWebhookAuthorization)
		out := out.(*KubeletWebhookAuthorization)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_KubeletX509Authentication(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*KubeletX509Authentication)
		out := out.(*KubeletX509Authentication)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_LeaderElectionConfiguration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LeaderElectionConfiguration)
		out := out.(*LeaderElectionConfiguration)
		*out = *in
		if in.LeaderElect != nil {
			in, out := &in.LeaderElect, &out.LeaderElect
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}
