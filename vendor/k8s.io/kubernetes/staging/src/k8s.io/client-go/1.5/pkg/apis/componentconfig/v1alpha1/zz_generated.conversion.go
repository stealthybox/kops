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

package v1alpha1

import (
	api "k8s.io/client-go/1.5/pkg/api"
	componentconfig "k8s.io/client-go/1.5/pkg/apis/componentconfig"
	conversion "k8s.io/client-go/1.5/pkg/conversion"
	runtime "k8s.io/client-go/1.5/pkg/runtime"
	config "k8s.io/client-go/1.5/pkg/util/config"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration,
		Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration,
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration,
		Convert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
	)
}

func autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	SetDefaults_KubeProxyConfiguration(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesMasqueradeBit = in.IPTablesMasqueradeBit
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	out.OOMScoreAdj = in.OOMScoreAdj
	out.Mode = componentconfig.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	out.ConntrackMax = in.ConntrackMax
	out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
	out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
	return nil
}

func Convert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in *KubeProxyConfiguration, out *componentconfig.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeProxyConfiguration_To_componentconfig_KubeProxyConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	out.IPTablesMasqueradeBit = in.IPTablesMasqueradeBit
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	out.OOMScoreAdj = in.OOMScoreAdj
	out.Mode = ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	out.ConntrackMax = in.ConntrackMax
	out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
	out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
	return nil
}

func Convert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in *componentconfig.KubeProxyConfiguration, out *KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeProxyConfiguration_To_v1alpha1_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	SetDefaults_KubeSchedulerConfiguration(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = int32(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int32(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Port = int(in.Port)
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableProfiling, &out.EnableProfiling, s); err != nil {
		return err
	}
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = int(in.KubeAPIBurst)
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	return nil
}

func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in *KubeletConfiguration, out *componentconfig.KubeletConfiguration, s conversion.Scope) error {
	SetDefaults_KubeletConfiguration(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.PodManifestPath = in.PodManifestPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableServer, &out.EnableServer, s); err != nil {
		return err
	}
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
	if err := api.Convert_Pointer_bool_To_bool(&in.AllowPrivileged, &out.AllowPrivileged, s); err != nil {
		return err
	}
	out.HostNetworkSources = in.HostNetworkSources
	out.HostPIDSources = in.HostPIDSources
	out.HostIPCSources = in.HostIPCSources
	if err := api.Convert_Pointer_int32_To_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := api.Convert_Pointer_int32_To_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.MinimumGCAge = in.MinimumGCAge
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	if err := api.Convert_Pointer_int32_To_int32(&in.MaxContainerCount, &out.MaxContainerCount, s); err != nil {
		return err
	}
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := api.Convert_Pointer_int32_To_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_bool_To_bool(&in.RegisterNode, &out.RegisterNode, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = in.ClusterDNS
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := api.Convert_Pointer_int32_To_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_int32_To_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginDir = in.NetworkPluginDir
	out.CNIConfDir = in.CNIConfDir
	out.CNIBinDir = in.CNIBinDir
	out.NetworkPluginMTU = in.NetworkPluginMTU
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := api.Convert_Pointer_bool_To_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.ContainerRuntime = in.ContainerRuntime
	out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
	out.RemoteImageEndpoint = in.RemoteImageEndpoint
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.RktPath = in.RktPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	if err := api.Convert_Pointer_string_To_string(&in.LockFilePath, &out.LockFilePath, s); err != nil {
		return err
	}
	out.ExitOnLockContention = in.ExitOnLockContention
	if err := api.Convert_Pointer_bool_To_bool(&in.ConfigureCBR0, &out.ConfigureCBR0, s); err != nil {
		return err
	}
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.NvidiaGPUs = in.NvidiaGPUs
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	if err := api.Convert_Pointer_bool_To_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_bool_To_bool(&in.Containerized, &out.Containerized, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	if err := api.Convert_Pointer_bool_To_bool(&in.ReconcileCIDR, &out.ReconcileCIDR, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_bool_To_bool(&in.RegisterSchedulable, &out.RegisterSchedulable, s); err != nil {
		return err
	}
	out.ContentType = in.ContentType
	if err := api.Convert_Pointer_int32_To_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := api.Convert_Pointer_bool_To_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.ExperimentalFlannelOverlay = in.ExperimentalFlannelOverlay
	out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
	out.NodeIP = in.NodeIP
	out.NodeLabels = in.NodeLabels
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	if err := api.Convert_Pointer_string_To_string(&in.EvictionHard, &out.EvictionHard, s); err != nil {
		return err
	}
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
	out.PodsPerCore = in.PodsPerCore
	if err := api.Convert_Pointer_bool_To_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(config.ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.SystemReserved = nil
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(config.ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.KubeReserved = nil
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := api.Convert_Pointer_bool_To_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_int32_To_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := api.Convert_Pointer_int32_To_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.AllowedUnsafeSysctls = in.AllowedUnsafeSysctls
	return nil
}

func Convert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in *KubeletConfiguration, out *componentconfig.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeletConfiguration_To_componentconfig_KubeletConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *componentconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.PodManifestPath = in.PodManifestPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableServer, &out.EnableServer, s); err != nil {
		return err
	}
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
	if err := api.Convert_bool_To_Pointer_bool(&in.AllowPrivileged, &out.AllowPrivileged, s); err != nil {
		return err
	}
	out.HostNetworkSources = in.HostNetworkSources
	out.HostPIDSources = in.HostPIDSources
	out.HostIPCSources = in.HostIPCSources
	if err := api.Convert_int32_To_Pointer_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := api.Convert_int32_To_Pointer_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.MinimumGCAge = in.MinimumGCAge
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	if err := api.Convert_int32_To_Pointer_int32(&in.MaxContainerCount, &out.MaxContainerCount, s); err != nil {
		return err
	}
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := api.Convert_int32_To_Pointer_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	if err := api.Convert_bool_To_Pointer_bool(&in.RegisterNode, &out.RegisterNode, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = in.ClusterDNS
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := api.Convert_int32_To_Pointer_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := api.Convert_int32_To_Pointer_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginMTU = in.NetworkPluginMTU
	out.NetworkPluginDir = in.NetworkPluginDir
	out.CNIConfDir = in.CNIConfDir
	out.CNIBinDir = in.CNIBinDir
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	if err := api.Convert_bool_To_Pointer_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	out.ContainerRuntime = in.ContainerRuntime
	out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
	out.RemoteImageEndpoint = in.RemoteImageEndpoint
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.RktPath = in.RktPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	if err := api.Convert_string_To_Pointer_string(&in.LockFilePath, &out.LockFilePath, s); err != nil {
		return err
	}
	out.ExitOnLockContention = in.ExitOnLockContention
	if err := api.Convert_bool_To_Pointer_bool(&in.ConfigureCBR0, &out.ConfigureCBR0, s); err != nil {
		return err
	}
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.NvidiaGPUs = in.NvidiaGPUs
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	if err := api.Convert_bool_To_Pointer_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	if err := api.Convert_bool_To_Pointer_bool(&in.Containerized, &out.Containerized, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	if err := api.Convert_bool_To_Pointer_bool(&in.ReconcileCIDR, &out.ReconcileCIDR, s); err != nil {
		return err
	}
	if err := api.Convert_bool_To_Pointer_bool(&in.RegisterSchedulable, &out.RegisterSchedulable, s); err != nil {
		return err
	}
	out.ContentType = in.ContentType
	if err := api.Convert_int32_To_Pointer_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := api.Convert_bool_To_Pointer_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.ExperimentalFlannelOverlay = in.ExperimentalFlannelOverlay
	out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
	out.NodeIP = in.NodeIP
	out.NodeLabels = in.NodeLabels
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	if err := api.Convert_string_To_Pointer_string(&in.EvictionHard, &out.EvictionHard, s); err != nil {
		return err
	}
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
	out.PodsPerCore = in.PodsPerCore
	if err := api.Convert_bool_To_Pointer_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.SystemReserved = nil
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.KubeReserved = nil
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := api.Convert_bool_To_Pointer_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := api.Convert_int32_To_Pointer_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := api.Convert_int32_To_Pointer_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.AllowedUnsafeSysctls = in.AllowedUnsafeSysctls
	return nil
}

func Convert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in *componentconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeletConfiguration_To_v1alpha1_KubeletConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	SetDefaults_LeaderElectionConfiguration(in)
	if err := api.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return nil
}

func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := api.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return nil
}

func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}
