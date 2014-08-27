/*
Copyright 2014 Google Inc. All rights reserved.

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

package api

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta1"
)

func init() {
	// TODO: Consider inverting dependency chain-- imagine v1beta1 package
	// registering all of these functions. Then, if you want to be able to understand
	// v1beta1 objects, you just import that package for its side effects.
	AddConversionFuncs(
		// EnvVar's Key is deprecated in favor of Name.
		func(in *EnvVar, out *v1beta1.EnvVar) error {
			out.Value = in.Value
			out.Key = in.Name
			out.Name = in.Name
			return nil
		},
		func(in *v1beta1.EnvVar, out *EnvVar) error {
			out.Value = in.Value
			if in.Name != "" {
				out.Name = in.Name
			} else {
				out.Name = in.Key
			}
			return nil
		},

		// Path & MountType are deprecated.
		func(in *VolumeMount, out *v1beta1.VolumeMount) error {
			out.Name = in.Name
			out.ReadOnly = in.ReadOnly
			out.MountPath = in.MountPath
			out.Path = in.MountPath
			out.MountType = "" // MountType is ignored.
			return nil
		},
		func(in *v1beta1.VolumeMount, out *VolumeMount) error {
			out.Name = in.Name
			out.ReadOnly = in.ReadOnly
			if in.MountPath == "" {
				out.MountPath = in.Path
			} else {
				out.MountPath = in.MountPath
			}
			return nil
		},
	)

}
