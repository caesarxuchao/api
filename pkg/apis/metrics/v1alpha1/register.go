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

package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/runtime"
)

// GroupName is the group name use in this package
const GroupName = "metrics"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1alpha1"}

var Codec = runtime.CodecFor(api.Scheme, SchemeGroupVersion)

func init() {
	// Register the API.
	addKnownTypes()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes() {
	api.Scheme.AddKnownTypes(SchemeGroupVersion,
		&RawNode{},
		&RawPod{},
		&v1.DeleteOptions{},
	)
}

func (obj *RawNode) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
func (obj *RawPod) GetObjectKind() unversioned.ObjectKind  { return &obj.TypeMeta }
