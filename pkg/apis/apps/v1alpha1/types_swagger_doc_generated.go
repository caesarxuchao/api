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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_StatefulSet = map[string]string{
	"":       "StatefulSet represents a set of pods with consistent identities. Identities are defined as:\n - Network: A single stable DNS and hostname.\n - Storage: As many VolumeClaims as requested.\nThe StatefulSet guarantees that a given network identity will always map to the same storage identity. StatefulSet is currently in alpha and subject to change without notice.",
	"spec":   "Spec defines the desired identities of pods in this set.",
	"status": "Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.",
}

func (StatefulSet) SwaggerDoc() map[string]string {
	return map_StatefulSet
}

var map_StatefulSetList = map[string]string{
	"": "StatefulSetList is a collection of StatefulSets.",
}

func (StatefulSetList) SwaggerDoc() map[string]string {
	return map_StatefulSetList
}

var map_StatefulSetSpec = map[string]string{
	"":                     "A StatefulSetSpec is the specification of a StatefulSet.",
	"replicas":             "Replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.",
	"selector":             "Selector is a label query over pods that should match the replica count. If empty, defaulted to labels on the pod template. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
	"template":             "Template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.",
	"volumeClaimTemplates": "VolumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.",
	"serviceName":          "ServiceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where \"pod-specific-string\" is managed by the StatefulSet controller.",
}

func (StatefulSetSpec) SwaggerDoc() map[string]string {
	return map_StatefulSetSpec
}

var map_StatefulSetStatus = map[string]string{
	"":                   "StatefulSetStatus represents the current state of a StatefulSet.",
	"observedGeneration": "most recent generation observed by this autoscaler.",
	"replicas":           "Replicas is the number of actual replicas.",
}

func (StatefulSetStatus) SwaggerDoc() map[string]string {
	return map_StatefulSetStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
