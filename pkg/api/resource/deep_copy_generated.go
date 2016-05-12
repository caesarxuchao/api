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

package resource

import (
	inf_v0 "gopkg.in/inf.v0"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func DeepCopy_resource_Quantity(in Quantity, out *Quantity, c *conversion.Cloner) error {
	if in.Amount != nil {
		in, out := in.Amount, &out.Amount
		*out = new(inf_v0.Dec)
		if newVal, err := c.DeepCopy(*in); err != nil {
			return err
		} else {
			**out = newVal.(inf_v0.Dec)
		}
	} else {
		out.Amount = nil
	}
	out.Format = in.Format
	return nil
}

func DeepCopy_resource_QuantityProto(in QuantityProto, out *QuantityProto, c *conversion.Cloner) error {
	out.Format = in.Format
	out.Scale = in.Scale
	if in.Bigint != nil {
		in, out := in.Bigint, &out.Bigint
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Bigint = nil
	}
	return nil
}
