/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-deep-copies.sh.

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func deepCopy_unversioned_Duration(in unversioned.Duration, out *unversioned.Duration, c *conversion.Cloner) error {
	out.Duration = in.Duration
	return nil
}

func deepCopy_unversioned_TypeMeta(in unversioned.TypeMeta, out *unversioned.TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func deepCopy_v1alpha1_KubeProxyConfiguration(in KubeProxyConfiguration, out *KubeProxyConfiguration, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	if err := deepCopy_unversioned_Duration(in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, c); err != nil {
		return err
	}
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		out.OOMScoreAdj = new(int32)
		*out.OOMScoreAdj = *in.OOMScoreAdj
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = in.Mode
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	if err := deepCopy_unversioned_Duration(in.UDPIdleTimeout, &out.UDPIdleTimeout, c); err != nil {
		return err
	}
	out.ConntrackMax = in.ConntrackMax
	if err := deepCopy_unversioned_Duration(in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1alpha1_KubeSchedulerConfiguration(in KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if in.EnableProfiling != nil {
		out.EnableProfiling = new(bool)
		*out.EnableProfiling = *in.EnableProfiling
	} else {
		out.EnableProfiling = nil
	}
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	if err := deepCopy_v1alpha1_LeaderElectionConfiguration(in.LeaderElection, &out.LeaderElection, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1alpha1_LeaderElectionConfiguration(in LeaderElectionConfiguration, out *LeaderElectionConfiguration, c *conversion.Cloner) error {
	if in.LeaderElect != nil {
		out.LeaderElect = new(bool)
		*out.LeaderElect = *in.LeaderElect
	} else {
		out.LeaderElect = nil
	}
	if err := deepCopy_unversioned_Duration(in.LeaseDuration, &out.LeaseDuration, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_Duration(in.RenewDeadline, &out.RenewDeadline, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_Duration(in.RetryPeriod, &out.RetryPeriod, c); err != nil {
		return err
	}
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_unversioned_Duration,
		deepCopy_unversioned_TypeMeta,
		deepCopy_v1alpha1_KubeProxyConfiguration,
		deepCopy_v1alpha1_KubeSchedulerConfiguration,
		deepCopy_v1alpha1_LeaderElectionConfiguration,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}
