//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	"github.com/kcp-dev/client-go/informers/internalinterfaces"
)

type ClusterInterface interface {
	// PersistentVolumes returns a PersistentVolumeClusterInformer
	PersistentVolumes() PersistentVolumeClusterInformer
	// PersistentVolumeClaims returns a PersistentVolumeClaimClusterInformer
	PersistentVolumeClaims() PersistentVolumeClaimClusterInformer
	// Pods returns a PodClusterInformer
	Pods() PodClusterInformer
	// PodTemplates returns a PodTemplateClusterInformer
	PodTemplates() PodTemplateClusterInformer
	// ReplicationControllers returns a ReplicationControllerClusterInformer
	ReplicationControllers() ReplicationControllerClusterInformer
	// Services returns a ServiceClusterInformer
	Services() ServiceClusterInformer
	// ServiceAccounts returns a ServiceAccountClusterInformer
	ServiceAccounts() ServiceAccountClusterInformer
	// Endpoints returns a EndpointsClusterInformer
	Endpoints() EndpointsClusterInformer
	// Nodes returns a NodeClusterInformer
	Nodes() NodeClusterInformer
	// Namespaces returns a NamespaceClusterInformer
	Namespaces() NamespaceClusterInformer
	// Events returns a EventClusterInformer
	Events() EventClusterInformer
	// LimitRanges returns a LimitRangeClusterInformer
	LimitRanges() LimitRangeClusterInformer
	// ResourceQuotas returns a ResourceQuotaClusterInformer
	ResourceQuotas() ResourceQuotaClusterInformer
	// Secrets returns a SecretClusterInformer
	Secrets() SecretClusterInformer
	// ConfigMaps returns a ConfigMapClusterInformer
	ConfigMaps() ConfigMapClusterInformer
	// ComponentStatuses returns a ComponentStatusClusterInformer
	ComponentStatuses() ComponentStatusClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// PersistentVolumes returns a PersistentVolumeClusterInformer
func (v *version) PersistentVolumes() PersistentVolumeClusterInformer {
	return &persistentVolumeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PersistentVolumeClaims returns a PersistentVolumeClaimClusterInformer
func (v *version) PersistentVolumeClaims() PersistentVolumeClaimClusterInformer {
	return &persistentVolumeClaimClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Pods returns a PodClusterInformer
func (v *version) Pods() PodClusterInformer {
	return &podClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PodTemplates returns a PodTemplateClusterInformer
func (v *version) PodTemplates() PodTemplateClusterInformer {
	return &podTemplateClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ReplicationControllers returns a ReplicationControllerClusterInformer
func (v *version) ReplicationControllers() ReplicationControllerClusterInformer {
	return &replicationControllerClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceClusterInformer
func (v *version) Services() ServiceClusterInformer {
	return &serviceClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ServiceAccounts returns a ServiceAccountClusterInformer
func (v *version) ServiceAccounts() ServiceAccountClusterInformer {
	return &serviceAccountClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Endpoints returns a EndpointsClusterInformer
func (v *version) Endpoints() EndpointsClusterInformer {
	return &endpointsClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Nodes returns a NodeClusterInformer
func (v *version) Nodes() NodeClusterInformer {
	return &nodeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Namespaces returns a NamespaceClusterInformer
func (v *version) Namespaces() NamespaceClusterInformer {
	return &namespaceClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Events returns a EventClusterInformer
func (v *version) Events() EventClusterInformer {
	return &eventClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LimitRanges returns a LimitRangeClusterInformer
func (v *version) LimitRanges() LimitRangeClusterInformer {
	return &limitRangeClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ResourceQuotas returns a ResourceQuotaClusterInformer
func (v *version) ResourceQuotas() ResourceQuotaClusterInformer {
	return &resourceQuotaClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Secrets returns a SecretClusterInformer
func (v *version) Secrets() SecretClusterInformer {
	return &secretClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapClusterInformer
func (v *version) ConfigMaps() ConfigMapClusterInformer {
	return &configMapClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ComponentStatuses returns a ComponentStatusClusterInformer
func (v *version) ComponentStatuses() ComponentStatusClusterInformer {
	return &componentStatusClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
