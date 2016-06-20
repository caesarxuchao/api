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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package componentconfig

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_componentconfig_IPVar,
		DeepCopy_componentconfig_KubeControllerManagerConfiguration,
		DeepCopy_componentconfig_KubeProxyConfiguration,
		DeepCopy_componentconfig_KubeSchedulerConfiguration,
		DeepCopy_componentconfig_KubeletConfiguration,
		DeepCopy_componentconfig_LeaderElectionConfiguration,
		DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration,
		DeepCopy_componentconfig_PortRangeVar,
		DeepCopy_componentconfig_VolumeConfiguration,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_componentconfig_IPVar(in IPVar, out *IPVar, c *conversion.Cloner) error {
	if in.Val != nil {
		in, out := in.Val, &out.Val
		*out = new(string)
		**out = *in
	} else {
		out.Val = nil
	}
	return nil
}

func DeepCopy_componentconfig_KubeControllerManagerConfiguration(in KubeControllerManagerConfiguration, out *KubeControllerManagerConfiguration, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
	out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
	out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
	out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
	out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
	out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
	out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
	out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
	out.LookupCacheSizeForRC = in.LookupCacheSizeForRC
	out.LookupCacheSizeForRS = in.LookupCacheSizeForRS
	out.LookupCacheSizeForDaemonSet = in.LookupCacheSizeForDaemonSet
	if err := unversioned.DeepCopy_unversioned_Duration(in.ServiceSyncPeriod, &out.ServiceSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.NodeSyncPeriod, &out.NodeSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.ResourceQuotaSyncPeriod, &out.ResourceQuotaSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.NamespaceSyncPeriod, &out.NamespaceSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.PVClaimBinderSyncPeriod, &out.PVClaimBinderSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.MinResyncPeriod, &out.MinResyncPeriod, c); err != nil {
		return err
	}
	out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
	if err := unversioned.DeepCopy_unversioned_Duration(in.HorizontalPodAutoscalerSyncPeriod, &out.HorizontalPodAutoscalerSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.DeploymentControllerSyncPeriod, &out.DeploymentControllerSyncPeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.PodEvictionTimeout, &out.PodEvictionTimeout, c); err != nil {
		return err
	}
	out.DeletingPodsQps = in.DeletingPodsQps
	out.DeletingPodsBurst = in.DeletingPodsBurst
	if err := unversioned.DeepCopy_unversioned_Duration(in.NodeMonitorGracePeriod, &out.NodeMonitorGracePeriod, c); err != nil {
		return err
	}
	out.RegisterRetryCount = in.RegisterRetryCount
	if err := unversioned.DeepCopy_unversioned_Duration(in.NodeStartupGracePeriod, &out.NodeStartupGracePeriod, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.NodeMonitorPeriod, &out.NodeMonitorPeriod, c); err != nil {
		return err
	}
	out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
	out.EnableProfiling = in.EnableProfiling
	out.ClusterName = in.ClusterName
	out.ClusterCIDR = in.ClusterCIDR
	out.ServiceCIDR = in.ServiceCIDR
	out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
	out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
	out.ConfigureCloudRoutes = in.ConfigureCloudRoutes
	out.RootCAFile = in.RootCAFile
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := DeepCopy_componentconfig_LeaderElectionConfiguration(in.LeaderElection, &out.LeaderElection, c); err != nil {
		return err
	}
	if err := DeepCopy_componentconfig_VolumeConfiguration(in.VolumeConfiguration, &out.VolumeConfiguration, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.ControllerStartInterval, &out.ControllerStartInterval, c); err != nil {
		return err
	}
	out.EnableGarbageCollector = in.EnableGarbageCollector
	return nil
}

func DeepCopy_componentconfig_KubeProxyConfiguration(in KubeProxyConfiguration, out *KubeProxyConfiguration, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	if in.IPTablesMasqueradeBit != nil {
		in, out := in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
		*out = new(int32)
		**out = *in
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.IPTablesSyncPeriod, &out.IPTablesSyncPeriod, c); err != nil {
		return err
	}
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		in, out := in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = *in
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = in.Mode
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	if err := unversioned.DeepCopy_unversioned_Duration(in.UDPIdleTimeout, &out.UDPIdleTimeout, c); err != nil {
		return err
	}
	out.ConntrackMax = in.ConntrackMax
	if err := unversioned.DeepCopy_unversioned_Duration(in.ConntrackTCPEstablishedTimeout, &out.ConntrackTCPEstablishedTimeout, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_componentconfig_KubeSchedulerConfiguration(in KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	out.EnableProfiling = in.EnableProfiling
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := DeepCopy_componentconfig_LeaderElectionConfiguration(in.LeaderElection, &out.LeaderElection, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_componentconfig_KubeletConfiguration(in KubeletConfiguration, out *KubeletConfiguration, c *conversion.Cloner) error {
	out.Config = in.Config
	if err := unversioned.DeepCopy_unversioned_Duration(in.SyncFrequency, &out.SyncFrequency, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.FileCheckFrequency, &out.FileCheckFrequency, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.HTTPCheckFrequency, &out.HTTPCheckFrequency, c); err != nil {
		return err
	}
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	out.EnableServer = in.EnableServer
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.CertDirectory = in.CertDirectory
	out.HostnameOverride = in.HostnameOverride
	out.PodInfraContainerImage = in.PodInfraContainerImage
	out.DockerEndpoint = in.DockerEndpoint
	out.RootDirectory = in.RootDirectory
	out.SeccompProfileRoot = in.SeccompProfileRoot
	out.AllowPrivileged = in.AllowPrivileged
	out.HostNetworkSources = in.HostNetworkSources
	out.HostPIDSources = in.HostPIDSources
	out.HostIPCSources = in.HostIPCSources
	out.RegistryPullQPS = in.RegistryPullQPS
	out.RegistryBurst = in.RegistryBurst
	out.EventRecordQPS = in.EventRecordQPS
	out.EventBurst = in.EventBurst
	out.EnableDebuggingHandlers = in.EnableDebuggingHandlers
	if err := unversioned.DeepCopy_unversioned_Duration(in.MinimumGCAge, &out.MinimumGCAge, c); err != nil {
		return err
	}
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	out.MaxContainerCount = in.MaxContainerCount
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	out.OOMScoreAdj = in.OOMScoreAdj
	out.RegisterNode = in.RegisterNode
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = in.ClusterDNS
	if err := unversioned.DeepCopy_unversioned_Duration(in.StreamingConnectionIdleTimeout, &out.StreamingConnectionIdleTimeout, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.NodeStatusUpdateFrequency, &out.NodeStatusUpdateFrequency, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.ImageMinimumGCAge, &out.ImageMinimumGCAge, c); err != nil {
		return err
	}
	out.ImageGCHighThresholdPercent = in.ImageGCHighThresholdPercent
	out.ImageGCLowThresholdPercent = in.ImageGCLowThresholdPercent
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	if err := unversioned.DeepCopy_unversioned_Duration(in.VolumeStatsAggPeriod, &out.VolumeStatsAggPeriod, c); err != nil {
		return err
	}
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginDir = in.NetworkPluginDir
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	out.ContainerRuntime = in.ContainerRuntime
	out.RktPath = in.RktPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	out.LockFilePath = in.LockFilePath
	out.ExitOnLockContention = in.ExitOnLockContention
	out.ConfigureCBR0 = in.ConfigureCBR0
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.NvidiaGPUs = in.NvidiaGPUs
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	out.CPUCFSQuota = in.CPUCFSQuota
	out.Containerized = in.Containerized
	out.MaxOpenFiles = in.MaxOpenFiles
	out.ReconcileCIDR = in.ReconcileCIDR
	out.RegisterSchedulable = in.RegisterSchedulable
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SerializeImagePulls = in.SerializeImagePulls
	out.ExperimentalFlannelOverlay = in.ExperimentalFlannelOverlay
	if err := unversioned.DeepCopy_unversioned_Duration(in.OutOfDiskTransitionFrequency, &out.OutOfDiskTransitionFrequency, c); err != nil {
		return err
	}
	out.NodeIP = in.NodeIP
	if in.NodeLabels != nil {
		in, out := in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.NodeLabels = nil
	}
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	out.EvictionHard = in.EvictionHard
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	if err := unversioned.DeepCopy_unversioned_Duration(in.EvictionPressureTransitionPeriod, &out.EvictionPressureTransitionPeriod, c); err != nil {
		return err
	}
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.PodsPerCore = in.PodsPerCore
	out.EnableControllerAttachDetach = in.EnableControllerAttachDetach
	return nil
}

func DeepCopy_componentconfig_LeaderElectionConfiguration(in LeaderElectionConfiguration, out *LeaderElectionConfiguration, c *conversion.Cloner) error {
	out.LeaderElect = in.LeaderElect
	if err := unversioned.DeepCopy_unversioned_Duration(in.LeaseDuration, &out.LeaseDuration, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.RenewDeadline, &out.RenewDeadline, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Duration(in.RetryPeriod, &out.RetryPeriod, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration(in PersistentVolumeRecyclerConfiguration, out *PersistentVolumeRecyclerConfiguration, c *conversion.Cloner) error {
	out.MaximumRetry = in.MaximumRetry
	out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
	out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
	out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
	out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
	out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
	out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
	return nil
}

func DeepCopy_componentconfig_PortRangeVar(in PortRangeVar, out *PortRangeVar, c *conversion.Cloner) error {
	if in.Val != nil {
		in, out := in.Val, &out.Val
		*out = new(string)
		**out = *in
	} else {
		out.Val = nil
	}
	return nil
}

func DeepCopy_componentconfig_VolumeConfiguration(in VolumeConfiguration, out *VolumeConfiguration, c *conversion.Cloner) error {
	out.EnableHostPathProvisioning = in.EnableHostPathProvisioning
	out.EnableDynamicProvisioning = in.EnableDynamicProvisioning
	if err := DeepCopy_componentconfig_PersistentVolumeRecyclerConfiguration(in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, c); err != nil {
		return err
	}
	out.FlexVolumePluginDir = in.FlexVolumePluginDir
	return nil
}
