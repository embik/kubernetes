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

package v1beta2

import (
	"github.com/kcp-dev/client-go/informers/internalinterfaces"
)

type ClusterInterface interface {
	// StatefulSets returns a StatefulSetClusterInformer
	StatefulSets() StatefulSetClusterInformer
	// Deployments returns a DeploymentClusterInformer
	Deployments() DeploymentClusterInformer
	// DaemonSets returns a DaemonSetClusterInformer
	DaemonSets() DaemonSetClusterInformer
	// ReplicaSets returns a ReplicaSetClusterInformer
	ReplicaSets() ReplicaSetClusterInformer
	// ControllerRevisions returns a ControllerRevisionClusterInformer
	ControllerRevisions() ControllerRevisionClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new ClusterInterface.
func New(f internalinterfaces.SharedInformerFactory, tweakListOptions internalinterfaces.TweakListOptionsFunc) ClusterInterface {
	return &version{factory: f, tweakListOptions: tweakListOptions}
}

// StatefulSets returns a StatefulSetClusterInformer
func (v *version) StatefulSets() StatefulSetClusterInformer {
	return &statefulSetClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentClusterInformer
func (v *version) Deployments() DeploymentClusterInformer {
	return &deploymentClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DaemonSets returns a DaemonSetClusterInformer
func (v *version) DaemonSets() DaemonSetClusterInformer {
	return &daemonSetClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ReplicaSets returns a ReplicaSetClusterInformer
func (v *version) ReplicaSets() ReplicaSetClusterInformer {
	return &replicaSetClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ControllerRevisions returns a ControllerRevisionClusterInformer
func (v *version) ControllerRevisions() ControllerRevisionClusterInformer {
	return &controllerRevisionClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
