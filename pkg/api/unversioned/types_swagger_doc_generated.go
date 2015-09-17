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

package unversioned

// This file contains a collection of methods that can be used from go-resful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser. (e.g. TODO(andronat):... || TODO:...) iff
// are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ListMeta = map[string]string{
	"":                "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
	"selfLink":        "SelfLink is a URL representing this object. Populated by the system. Read-only.",
	"resourceVersion": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#concurrency-control-and-consistency",
}

func (ListMeta) SwaggerDoc() map[string]string {
	return map_ListMeta
}

var map_Status = map[string]string{
	"":         "Status is a return value for calls that don't return other objects.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds",
	"status":   "Status of the operation. One of: \"Success\" or \"Failure\". More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"message":  "A human-readable description of the status of this operation.",
	"reason":   "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
	"details":  "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
	"code":     "Suggested HTTP return code for this status, 0 if not set.",
}

func (Status) SwaggerDoc() map[string]string {
	return map_Status
}

var map_StatusCause = map[string]string{
	"":        "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
	"reason":  "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
	"message": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
	"field":   "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
}

func (StatusCause) SwaggerDoc() map[string]string {
	return map_StatusCause
}

var map_StatusDetails = map[string]string{
	"":                  "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
	"name":              "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
	"kind":              "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds",
	"causes":            "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
	"retryAfterSeconds": "If specified, the time in seconds before the operation should be retried.",
}

func (StatusDetails) SwaggerDoc() map[string]string {
	return map_StatusDetails
}

var map_TypeMeta = map[string]string{
	"":           "TypeMeta describes an individual object in an API response or request with strings representing the type of the object and its API schema version. Structures that are versioned or persisted should inline TypeMeta.",
	"kind":       "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds",
	"apiVersion": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#resources",
}

func (TypeMeta) SwaggerDoc() map[string]string {
	return map_TypeMeta
}

// AUTO-GENERATED FUNCTIONS END HERE
