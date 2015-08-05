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

package api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/runtime"
)

var (
	// Errors that could be returned by GetReference.
	ErrNilObject  = errors.New("can't reference a nil object")
	ErrNoSelfLink = errors.New("selfLink was empty, can't make reference")
)

// ForTesting_ReferencesAllowBlankSelfLinks can be set to true in tests to avoid
// "ErrNoSelfLink" errors.
var ForTesting_ReferencesAllowBlankSelfLinks = false

// GetReference returns an ObjectReference which refers to the given
// object, or an error if the object doesn't follow the conventions
// that would allow this.
// TODO: should take a meta.Interface see https://github.com/GoogleCloudPlatform/kubernetes/issues/7127
func GetReference(obj runtime.Object) (*ObjectReference, error) {
	if obj == nil {
		return nil, ErrNilObject
	}
	if ref, ok := obj.(*ObjectReference); ok {
		// Don't make a reference to a reference.
		return ref, nil
	}
	meta, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}

	// if the object referenced is actually persisted, we can just get kind from meta
	// if we are building an object reference to something not yet persisted, we should fallback to scheme
	kind := meta.Kind()
	if kind == "" {
		_, kind, err = Scheme.ObjectVersionAndKind(obj)
		if err != nil {
			return nil, err
		}
	}

	// if the object referenced is actually persisted, we can also get version from meta
	version := meta.APIVersion()
	if version == "" {
		selfLink := meta.SelfLink()
		if selfLink == "" {
			if ForTesting_ReferencesAllowBlankSelfLinks {
				version = "testing"
			} else {
				return nil, ErrNoSelfLink
			}
		} else {
			selfLinkUrl, err := url.Parse(selfLink)
			if err != nil {
				return nil, err
			}
			// example paths: /<prefix>/<version>/*
			parts := strings.Split(selfLinkUrl.Path, "/")
			if len(parts) < 3 {
				return nil, fmt.Errorf("unexpected self link format: '%v'; got version '%v'", selfLink, version)
			}
			version = parts[2]
		}
	}

	return &ObjectReference{
		Kind:            kind,
		APIVersion:      version,
		Name:            meta.Name(),
		Namespace:       meta.Namespace(),
		UID:             meta.UID(),
		ResourceVersion: meta.ResourceVersion(),
	}, nil
}

// GetPartialReference is exactly like GetReference, but allows you to set the FieldPath.
func GetPartialReference(obj runtime.Object, fieldPath string) (*ObjectReference, error) {
	ref, err := GetReference(obj)
	if err != nil {
		return nil, err
	}
	ref.FieldPath = fieldPath
	return ref, nil
}

// IsAnAPIObject allows clients to preemptively get a reference to an API object and pass it to places that
// intend only to get a reference to that object. This simplifies the event recording interface.
func (*ObjectReference) IsAnAPIObject() {}
