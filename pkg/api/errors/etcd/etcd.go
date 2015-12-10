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

package etcd

import (
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/storage"
)

// InterpretListError converts a generic error on a retrieval
// operation into the appropriate API error.
func InterpretListError(err error, kind string) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(kind, "")
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(kind, "list", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretGetError converts a generic error on a retrieval
// operation into the appropriate API error.
func InterpretGetError(err error, kind, name string) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(kind, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(kind, "get", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretCreateError converts a generic error on a create
// operation into the appropriate API error.
func InterpretCreateError(err error, kind, name string) error {
	switch {
	case storage.IsNodeExist(err):
		return errors.NewAlreadyExists(kind, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(kind, "create", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretUpdateError converts a generic error on a update
// operation into the appropriate API error.
func InterpretUpdateError(err error, kind, name string) error {
	switch {
	case storage.IsTestFailed(err), storage.IsNodeExist(err):
		return errors.NewConflict(kind, name, err)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(kind, "update", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretDeleteError converts a generic error on a delete
// operation into the appropriate API error.
func InterpretDeleteError(err error, kind, name string) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(kind, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(kind, "delete", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}
