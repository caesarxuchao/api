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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	api "k8s.io/kubernetes/pkg/api"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1beta1_TokenReview,
		DeepCopy_v1beta1_TokenReviewSpec,
		DeepCopy_v1beta1_TokenReviewStatus,
		DeepCopy_v1beta1_UserInfo,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1beta1_TokenReview(in TokenReview, out *TokenReview, c *conversion.Cloner) error {
	out.TypeMeta = in.TypeMeta
	out.Spec = in.Spec
	if err := DeepCopy_v1beta1_TokenReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_TokenReviewSpec(in TokenReviewSpec, out *TokenReviewSpec, c *conversion.Cloner) error {
	out.Token = in.Token
	return nil
}

func DeepCopy_v1beta1_TokenReviewStatus(in TokenReviewStatus, out *TokenReviewStatus, c *conversion.Cloner) error {
	out.Authenticated = in.Authenticated
	if err := DeepCopy_v1beta1_UserInfo(in.User, &out.User, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_UserInfo(in UserInfo, out *UserInfo, c *conversion.Cloner) error {
	out.Username = in.Username
	out.UID = in.UID
	if in.Groups != nil {
		in, out := in.Groups, &out.Groups
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Groups = nil
	}
	if in.Extra != nil {
		in, out := in.Extra, &out.Extra
		*out = make(map[string][]string)
		for key, val := range in {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				(*out)[key] = newVal.([]string)
			}
		}
	} else {
		out.Extra = nil
	}
	return nil
}
