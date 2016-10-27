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
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole,
		Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole,
		Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding,
		Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding,
		Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList,
		Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList,
		Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList,
		Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList,
		Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule,
		Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule,
		Convert_v1alpha1_Role_To_rbac_Role,
		Convert_rbac_Role_To_v1alpha1_Role,
		Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding,
		Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding,
		Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList,
		Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList,
		Convert_v1alpha1_RoleList_To_rbac_RoleList,
		Convert_rbac_RoleList_To_v1alpha1_RoleList,
		Convert_v1alpha1_RoleRef_To_rbac_RoleRef,
		Convert_rbac_RoleRef_To_v1alpha1_RoleRef,
		Convert_v1alpha1_Subject_To_rbac_Subject,
		Convert_rbac_Subject_To_v1alpha1_Subject,
	)
}

func autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbac.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in *ClusterRole, out *rbac.ClusterRole, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRole_To_rbac_ClusterRole(in, out, s)
}

func autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *ClusterRole, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in *rbac.ClusterRole, out *ClusterRole, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRole_To_v1alpha1_ClusterRole(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Subjects = *(*[]rbac.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in *ClusterRoleBinding, out *rbac.ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleBinding_To_rbac_ClusterRoleBinding(in, out, s)
}

func autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *ClusterRoleBinding, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Subjects = *(*[]Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

func Convert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in *rbac.ClusterRoleBinding, out *ClusterRoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBinding_To_v1alpha1_ClusterRoleBinding(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.ClusterRoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in *ClusterRoleBindingList, out *rbac.ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleBindingList_To_rbac_ClusterRoleBindingList(in, out, s)
}

func autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *ClusterRoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterRoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in *rbac.ClusterRoleBindingList, out *ClusterRoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleBindingList_To_v1alpha1_ClusterRoleBindingList(in, out, s)
}

func autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]rbac.ClusterRole, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ClusterRole_To_rbac_ClusterRole(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in *ClusterRoleList, out *rbac.ClusterRoleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterRoleList_To_rbac_ClusterRoleList(in, out, s)
}

func autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *ClusterRoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRole, len(*in))
		for i := range *in {
			if err := Convert_rbac_ClusterRole_To_v1alpha1_ClusterRole(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in *rbac.ClusterRoleList, out *ClusterRoleList, s conversion.Scope) error {
	return autoConvert_rbac_ClusterRoleList_To_v1alpha1_ClusterRoleList(in, out, s)
}

func autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

func Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in *PolicyRule, out *rbac.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1alpha1_PolicyRule_To_rbac_PolicyRule(in, out, s)
}

func autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.AttributeRestrictions, &out.AttributeRestrictions, s); err != nil {
		return err
	}
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

func Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in *rbac.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	return autoConvert_rbac_PolicyRule_To_v1alpha1_PolicyRule(in, out, s)
}

func autoConvert_v1alpha1_Role_To_rbac_Role(in *Role, out *rbac.Role, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbac.PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_PolicyRule_To_rbac_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func Convert_v1alpha1_Role_To_rbac_Role(in *Role, out *rbac.Role, s conversion.Scope) error {
	return autoConvert_v1alpha1_Role_To_rbac_Role(in, out, s)
}

func autoConvert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *Role, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			if err := Convert_rbac_PolicyRule_To_v1alpha1_PolicyRule(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func Convert_rbac_Role_To_v1alpha1_Role(in *rbac.Role, out *Role, s conversion.Scope) error {
	return autoConvert_rbac_Role_To_v1alpha1_Role(in, out, s)
}

func autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Subjects = *(*[]rbac.Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_v1alpha1_RoleRef_To_rbac_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in *RoleBinding, out *rbac.RoleBinding, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleBinding_To_rbac_RoleBinding(in, out, s)
}

func autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *RoleBinding, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Subjects = *(*[]Subject)(unsafe.Pointer(&in.Subjects))
	if err := Convert_rbac_RoleRef_To_v1alpha1_RoleRef(&in.RoleRef, &out.RoleRef, s); err != nil {
		return err
	}
	return nil
}

func Convert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in *rbac.RoleBinding, out *RoleBinding, s conversion.Scope) error {
	return autoConvert_rbac_RoleBinding_To_v1alpha1_RoleBinding(in, out, s)
}

func autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]rbac.RoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in *RoleBindingList, out *rbac.RoleBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleBindingList_To_rbac_RoleBindingList(in, out, s)
}

func autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]RoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in *rbac.RoleBindingList, out *RoleBindingList, s conversion.Scope) error {
	return autoConvert_rbac_RoleBindingList_To_v1alpha1_RoleBindingList(in, out, s)
}

func autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in *RoleList, out *rbac.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]rbac.Role, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Role_To_rbac_Role(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1alpha1_RoleList_To_rbac_RoleList(in *RoleList, out *rbac.RoleList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleList_To_rbac_RoleList(in, out, s)
}

func autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Role, len(*in))
		for i := range *in {
			if err := Convert_rbac_Role_To_v1alpha1_Role(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_rbac_RoleList_To_v1alpha1_RoleList(in *rbac.RoleList, out *RoleList, s conversion.Scope) error {
	return autoConvert_rbac_RoleList_To_v1alpha1_RoleList(in, out, s)
}

func autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in *RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

func Convert_v1alpha1_RoleRef_To_rbac_RoleRef(in *RoleRef, out *rbac.RoleRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_RoleRef_To_rbac_RoleRef(in, out, s)
}

func autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *RoleRef, s conversion.Scope) error {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return nil
}

func Convert_rbac_RoleRef_To_v1alpha1_RoleRef(in *rbac.RoleRef, out *RoleRef, s conversion.Scope) error {
	return autoConvert_rbac_RoleRef_To_v1alpha1_RoleRef(in, out, s)
}

func autoConvert_v1alpha1_Subject_To_rbac_Subject(in *Subject, out *rbac.Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

func Convert_v1alpha1_Subject_To_rbac_Subject(in *Subject, out *rbac.Subject, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subject_To_rbac_Subject(in, out, s)
}

func autoConvert_rbac_Subject_To_v1alpha1_Subject(in *rbac.Subject, out *Subject, s conversion.Scope) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

func Convert_rbac_Subject_To_v1alpha1_Subject(in *rbac.Subject, out *Subject, s conversion.Scope) error {
	return autoConvert_rbac_Subject_To_v1alpha1_Subject(in, out, s)
}
