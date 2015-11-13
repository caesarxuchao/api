/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

// TODO: move everything in this file to pkg/api/rest
package meta

import (
	"fmt"
	"strings"

	"k8s.io/kubernetes/pkg/api/unversioned"
)

// Implements RESTScope interface
type restScope struct {
	name             RESTScopeName
	paramName        string
	argumentName     string
	paramDescription string
}

func (r *restScope) Name() RESTScopeName {
	return r.name
}
func (r *restScope) ParamName() string {
	return r.paramName
}
func (r *restScope) ArgumentName() string {
	return r.argumentName
}
func (r *restScope) ParamDescription() string {
	return r.paramDescription
}

var RESTScopeNamespace = &restScope{
	name:             RESTScopeNameNamespace,
	paramName:        "namespaces",
	argumentName:     "namespace",
	paramDescription: "object name and auth scope, such as for teams and projects",
}

var RESTScopeRoot = &restScope{
	name: RESTScopeNameRoot,
}

// typeMeta is used as a key for lookup in the mapping between REST path and
// API object.
type typeMeta struct {
	APIVersion string
	Kind       string
}

// DefaultRESTMapper exposes mappings between the types defined in a
// runtime.Scheme. It assumes that all types defined the provided scheme
// can be mapped with the provided MetadataAccessor and Codec interfaces.
//
// The resource name of a Kind is defined as the lowercase,
// English-plural version of the Kind string.
// When converting from resource to Kind, the singular version of the
// resource name is also accepted for convenience.
//
// TODO: Only accept plural for some operations for increased control?
// (`get pod bar` vs `get pods bar`)
// TODO these maps should be keyed based on GroupVersionKinds
type DefaultRESTMapper struct {
	mapping        map[string]typeMeta
	reverse        map[typeMeta]string
	scopes         map[typeMeta]RESTScope
	groupVersions  []unversioned.GroupVersion
	plurals        map[string]string
	singulars      map[string]string
	interfacesFunc VersionInterfacesFunc
}

// VersionInterfacesFunc returns the appropriate codec, typer, and metadata accessor for a
// given api version, or an error if no such api version exists.
type VersionInterfacesFunc func(apiVersion string) (*VersionInterfaces, error)

// NewDefaultRESTMapper initializes a mapping between Kind and APIVersion
// to a resource name and back based on the objects in a runtime.Scheme
// and the Kubernetes API conventions. Takes a group name, a priority list of the versions
// to search when an object has no default version (set empty to return an error),
// and a function that retrieves the correct codec and metadata for a given version.
// TODO remove group when this API is fixed.  It is no longer used.
// The external API for a RESTMapper is cross-version and this is currently called using
// group/version tuples.  In the end, the structure may be easier to understand with
// a GroupRESTMapper and CrossGroupRESTMapper, but for now, this one is constructed and
// used a CrossGroupRESTMapper.
func NewDefaultRESTMapper(group string, gvStrings []string, f VersionInterfacesFunc) *DefaultRESTMapper {
	mapping := make(map[string]typeMeta)
	reverse := make(map[typeMeta]string)
	scopes := make(map[typeMeta]RESTScope)
	plurals := make(map[string]string)
	singulars := make(map[string]string)
	// TODO: verify name mappings work correctly when versions differ

	gvs := []unversioned.GroupVersion{}
	for _, gvString := range gvStrings {
		gvs = append(gvs, unversioned.ParseGroupVersionOrDie(gvString))
	}

	return &DefaultRESTMapper{
		mapping:        mapping,
		reverse:        reverse,
		scopes:         scopes,
		groupVersions:  gvs,
		plurals:        plurals,
		singulars:      singulars,
		interfacesFunc: f,
	}
}

func (m *DefaultRESTMapper) Add(scope RESTScope, kind string, gvString string, mixedCase bool) {
	gv := unversioned.ParseGroupVersionOrDie(gvString)

	plural, singular := KindToResource(kind, mixedCase)
	m.plurals[singular] = plural
	m.singulars[plural] = singular
	meta := typeMeta{APIVersion: gv.String(), Kind: kind}
	_, ok1 := m.mapping[plural]
	_, ok2 := m.mapping[strings.ToLower(plural)]
	if !ok1 && !ok2 {
		m.mapping[plural] = meta
		m.mapping[singular] = meta
		if strings.ToLower(plural) != plural {
			m.mapping[strings.ToLower(plural)] = meta
			m.mapping[strings.ToLower(singular)] = meta
		}
	}
	m.reverse[meta] = plural
	m.scopes[meta] = scope
}

// KindToResource converts Kind to a resource name.
func KindToResource(kind string, mixedCase bool) (plural, singular string) {
	if len(kind) == 0 {
		return
	}
	if mixedCase {
		// Legacy support for mixed case names
		singular = strings.ToLower(kind[:1]) + kind[1:]
	} else {
		singular = strings.ToLower(kind)
	}
	if strings.HasSuffix(singular, "endpoints") {
		plural = singular
	} else {
		switch string(singular[len(singular)-1]) {
		case "s":
			plural = singular + "es"
		case "y":
			plural = strings.TrimSuffix(singular, "y") + "ies"
		default:
			plural = singular + "s"
		}
	}
	return
}

// ResourceSingularizer implements RESTMapper
// It converts a resource name from plural to singular (e.g., from pods to pod)
func (m *DefaultRESTMapper) ResourceSingularizer(resource string) (singular string, err error) {
	singular, ok := m.singulars[resource]
	if !ok {
		return resource, fmt.Errorf("no singular of resource %q has been defined", resource)
	}
	return singular, nil
}

// VersionAndKindForResource implements RESTMapper
func (m *DefaultRESTMapper) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	meta, ok := m.mapping[strings.ToLower(resource)]
	if !ok {
		return "", "", fmt.Errorf("in version and kind for resource, no resource %q has been defined", resource)
	}
	return meta.APIVersion, meta.Kind, nil
}

func (m *DefaultRESTMapper) GroupForResource(resource string) (string, error) {
	typemeta, exists := m.mapping[strings.ToLower(resource)]
	if !exists {
		return "", fmt.Errorf("in group for resource, no resource %q has been defined", resource)
	}

	gv, err := unversioned.ParseGroupVersion(typemeta.APIVersion)
	if err != nil {
		return "", err
	}
	return gv.Group, nil
}

// RESTMapping returns a struct representing the resource path and conversion interfaces a
// RESTClient should use to operate on the provided kind in order of versions. If a version search
// order is not provided, the search order provided to DefaultRESTMapper will be used to resolve which
// APIVersion should be used to access the named kind.
// TODO version here in this RESTMapper means just APIVersion, but the RESTMapper API is intended to handle multiple groups
// So this API is broken.  The RESTMapper test made it clear that versions here were API versions, but the code tries to use
// them with group/version tuples.
// TODO this should probably become RESTMapping(GroupKind, versions ...string)
func (m *DefaultRESTMapper) RESTMapping(kind string, versions ...string) (*RESTMapping, error) {
	// Pick an appropriate version
	var groupVersion *unversioned.GroupVersion
	hadVersion := false
	for _, v := range versions {
		if len(v) == 0 {
			continue
		}
		currGroupVersion, err := unversioned.ParseGroupVersion(v)
		if err != nil {
			return nil, err
		}

		hadVersion = true
		if _, ok := m.reverse[typeMeta{APIVersion: currGroupVersion.String(), Kind: kind}]; ok {
			groupVersion = &currGroupVersion
			break
		}
	}
	// Use the default preferred versions
	if !hadVersion && (groupVersion == nil) {
		for _, currGroupVersion := range m.groupVersions {
			if _, ok := m.reverse[typeMeta{APIVersion: currGroupVersion.String(), Kind: kind}]; ok {
				groupVersion = &currGroupVersion
				break
			}
		}
	}
	if groupVersion == nil {
		return nil, fmt.Errorf("no kind named %q is registered in versions %q", kind, versions)
	}

	gvk := unversioned.NewGroupVersionKind(*groupVersion, kind)

	// Ensure we have a REST mapping
	resource, ok := m.reverse[typeMeta{APIVersion: gvk.GroupVersion().String(), Kind: gvk.Kind}]
	if !ok {
		found := []unversioned.GroupVersion{}
		for _, gv := range m.groupVersions {
			if _, ok := m.reverse[typeMeta{APIVersion: gv.String(), Kind: kind}]; ok {
				found = append(found, gv)
			}
		}
		if len(found) > 0 {
			return nil, fmt.Errorf("object with kind %q exists in versions %v, not %v", kind, found, *groupVersion)
		}
		return nil, fmt.Errorf("the provided version %q and kind %q cannot be mapped to a supported object", groupVersion, kind)
	}

	// Ensure we have a REST scope
	scope, ok := m.scopes[typeMeta{APIVersion: gvk.GroupVersion().String(), Kind: gvk.Kind}]
	if !ok {
		return nil, fmt.Errorf("the provided version %q and kind %q cannot be mapped to a supported scope", gvk.GroupVersion().String(), gvk.Kind)
	}

	interfaces, err := m.interfacesFunc(gvk.GroupVersion().String())
	if err != nil {
		return nil, fmt.Errorf("the provided version %q has no relevant versions", gvk.GroupVersion().String())
	}

	retVal := &RESTMapping{
		Resource:         resource,
		GroupVersionKind: gvk,
		Scope:            scope,

		Codec:            interfaces.Codec,
		ObjectConvertor:  interfaces.ObjectConvertor,
		MetadataAccessor: interfaces.MetadataAccessor,
	}

	return retVal, nil
}

// aliasToResource is used for mapping aliases to resources
var aliasToResource = map[string][]string{}

// AddResourceAlias maps aliases to resources
func (m *DefaultRESTMapper) AddResourceAlias(alias string, resources ...string) {
	if len(resources) == 0 {
		return
	}
	aliasToResource[alias] = resources
}

// AliasesForResource returns whether a resource has an alias or not
func (m *DefaultRESTMapper) AliasesForResource(alias string) ([]string, bool) {
	if res, ok := aliasToResource[alias]; ok {
		return res, true
	}
	return nil, false
}

// ResourceIsValid takes a string (kind) and checks if it's a valid resource
func (m *DefaultRESTMapper) ResourceIsValid(resource string) bool {
	_, _, err := m.VersionAndKindForResource(resource)
	return err == nil
}

// MultiRESTMapper is a wrapper for multiple RESTMappers.
type MultiRESTMapper []RESTMapper

// ResourceSingularizer converts a REST resource name from plural to singular (e.g., from pods to pod)
// This implementation supports multiple REST schemas and return the first match.
func (m MultiRESTMapper) ResourceSingularizer(resource string) (singular string, err error) {
	for _, t := range m {
		singular, err = t.ResourceSingularizer(resource)
		if err == nil {
			return
		}
	}
	return
}

// VersionAndKindForResource provides the Version and Kind  mappings for the
// REST resources. This implementation supports multiple REST schemas and return
// the first match.
func (m MultiRESTMapper) VersionAndKindForResource(resource string) (defaultVersion, kind string, err error) {
	for _, t := range m {
		defaultVersion, kind, err = t.VersionAndKindForResource(resource)
		if err == nil {
			return
		}
	}
	return
}

// GroupForResource provides the Group mappings for the REST resources. This
// implementation supports multiple REST schemas and returns the first match.
func (m MultiRESTMapper) GroupForResource(resource string) (group string, err error) {
	for _, t := range m {
		group, err = t.GroupForResource(resource)
		if err == nil {
			return
		}
	}
	return
}

// RESTMapping provides the REST mapping for the resource based on the resource
// kind and version. This implementation supports multiple REST schemas and
// return the first match.
func (m MultiRESTMapper) RESTMapping(kind string, versions ...string) (mapping *RESTMapping, err error) {
	for _, t := range m {
		mapping, err = t.RESTMapping(kind, versions...)
		if err == nil {
			return
		}
	}
	return
}

// AliasesForResource finds the first alias response for the provided mappers.
func (m MultiRESTMapper) AliasesForResource(alias string) (aliases []string, ok bool) {
	for _, t := range m {
		if aliases, ok = t.AliasesForResource(alias); ok {
			return
		}
	}
	return nil, false
}

// ResourceIsValid takes a string (either group/kind or kind) and checks if it's a valid resource
func (m MultiRESTMapper) ResourceIsValid(resource string) bool {
	for _, t := range m {
		if t.ResourceIsValid(resource) {
			return true
		}
	}
	return false
}
