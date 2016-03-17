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

package unversioned

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	time "time"
)

func DeepCopy_unversioned_GroupKind(in GroupKind, out *GroupKind, c *conversion.Cloner) error {
	out.Group = in.Group
	out.Kind = in.Kind
	return nil
}

func DeepCopy_unversioned_GroupResource(in GroupResource, out *GroupResource, c *conversion.Cloner) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

func DeepCopy_unversioned_GroupVersion(in GroupVersion, out *GroupVersion, c *conversion.Cloner) error {
	out.Group = in.Group
	out.Version = in.Version
	return nil
}

func DeepCopy_unversioned_GroupVersionKind(in GroupVersionKind, out *GroupVersionKind, c *conversion.Cloner) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Kind = in.Kind
	return nil
}

func DeepCopy_unversioned_GroupVersionResource(in GroupVersionResource, out *GroupVersionResource, c *conversion.Cloner) error {
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	return nil
}

func DeepCopy_unversioned_ListMeta(in ListMeta, out *ListMeta, c *conversion.Cloner) error {
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func DeepCopy_unversioned_Time(in Time, out *Time, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Time); err != nil {
		return err
	} else {
		out.Time = newVal.(time.Time)
	}
	return nil
}

func DeepCopy_unversioned_TypeMeta(in TypeMeta, out *TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}
