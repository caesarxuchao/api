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

package v1beta3

import (
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

// Common string formats
// ---------------------
// Many fields in this API have formatting requirements.  The commonly used
// formats are defined here.
//
// C_IDENTIFIER:  This is a string that conforms to the definition of an "identifier"
//     in the C language.  This is captured by the following regex:
//         [A-Za-z_][A-Za-z0-9_]*
//     This defines the format, but not the length restriction, which should be
//     specified at the definition of any field of this type.
//
// DNS_LABEL:  This is a string, no more than 63 characters long, that conforms
//     to the definition of a "label" in RFCs 1035 and 1123.  This is captured
//     by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?
//
// DNS_SUBDOMAIN:  This is a string, no more than 253 characters long, that conforms
//      to the definition of a "subdomain" in RFCs 1035 and 1123.  This is captured
//      by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
//     or more simply:
//         DNS_LABEL(\.DNS_LABEL)*

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	Kind string `json:"kind,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	APIVersion string `json:"apiVersion,omitempty"`
}

// ListMeta describes metadata that synthetic resources must have, including lists and
// various status objects.
type ListMeta struct {
	// SelfLink is a URL representing this object.
	SelfLink string `json:"selfLink,omitempty"`

	// An opaque value that represents the version of this response for use with optimistic
	// concurrency and change monitoring endpoints.  Clients must treat these values as opaque
	// and values may only be valid for a particular resource or set of resources. Only servers
	// will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create.
type ObjectMeta struct {
	// Name is unique within a namespace.  Name is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	// SelfLink is a URL representing this object.
	SelfLink string `json:"selfLink,omitempty"`

	// UID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	UID string `json:"uid,omitempty"`

	// An opaque value that represents the version of this resource. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and values may only be valid for a particular
	// resource or set of resources. Only servers will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp util.Time `json:"creationTimestamp,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// TODO: replace map[string]string with labels.LabelSet type
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.
	Annotations map[string]string `json:"annotations,omitempty"`
}

const (
	// NamespaceDefault means the object is in the default namespace which is applied when not specified by clients
	NamespaceDefault string = "default"
	// NamespaceAll is the default argument to specify on a context when you want to list or filter resources across all namespaces
	NamespaceAll string = ""
)

//
//// ContainerManifest corresponds to the Container Manifest format, documented at:
//// https://developers.google.com/compute/docs/containers/container_vms#container_manifest
//// This is used as the representation of Kubernetes workloads.
//// DEPRECATED: Exists to allow backwards compatible storage for clients accessing etcd
//// directly.
//type ContainerManifest struct {
//	// Required: This must be a supported version string, such as "v1beta1".
//	Version string `json:"version"`
//	// Required: This must be a DNS_SUBDOMAIN.
//	// TODO: ID on Manifest is deprecated and will be removed in the future.
//	ID string `json:"id"`
//	// TODO: UUID on Manifest is deprecated in the future once we are done
//	// with the API refactoring. It is required for now to determine the instance
//	// of a Pod.
//	UUID          string        `json:"uuid,omitempty"`
//	Volumes       []Volume      `json:"volumes"`
//	Containers    []Container   `json:"containers"`
//	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
//}
//
//// ContainerManifestList is used to communicate container manifests to kubelet.
//// DEPRECATED: Exists to allow backwards compatible storage for clients accessing etcd
//// directly.
//type ContainerManifestList struct {
//	TypeMeta `json:",inline"`
//	// ID is the legacy field representing Name
//	ID string `json:"id,omitempty"`
//
//	Items []ContainerManifest `json:"items,omitempty"`
//}

// Volume represents a named volume in a pod that may be accessed by any containers in the pod.
type Volume struct {
	// Required: This must be a DNS_LABEL.  Each volume in a pod must have
	// a unique name.
	Name string `json:"name"`
	// Source represents the location and type of a volume to mount.
	// This is optional for now. If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	Source *VolumeSource `json:"source"`
}

// VolumeSource represents the source location of a valume to mount.
// Only one of its members may be specified.
type VolumeSource struct {
	// HostDir represents a pre-existing directory on the host machine that is directly
	// exposed to the container. This is generally used for system agents or other privileged
	// things that are allowed to see the host machine. Most containers will NOT need this.
	// TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not
	// mount host directories as read/write.
	HostDir *HostDir `json:"hostDir"`
	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	EmptyDir *EmptyDir `json:"emptyDir"`
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	GCEPersistentDisk *GCEPersistentDisk `json:"persistentDisk"`
	// GitRepo represents a git repository at a particular revision.
	GitRepo *GitRepo `json:"gitRepo"`
}

// HostDir represents bare host directory volume.
type HostDir struct {
	Path string `json:"path"`
}

type EmptyDir struct{}

// Protocol defines network protocols supported for things like conatiner ports.
type Protocol string

const (
	// ProtocolTCP is the TCP protocol.
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP is the UDP protocol.
	ProtocolUDP Protocol = "UDP"
)

// GCEPersistentDisk represents a Persistent Disk resource in Google Compute Engine.
//
// A GCE PD must exist and be formatted before mounting to a container.
// The disk must also be in the same GCE project and zone as the kubelet.
// A GCE PD can only be mounted as read/write once.
type GCEPersistentDisk struct {
	// Unique name of the PD resource. Used to identify the disk in GCE
	PDName string `json:"pdName"`
	// Required: Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs"
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	FSType string `json:"fsType,omitempty"`
	// Optional: Partition on the disk to mount.
	// If omitted, kubelet will attempt to mount the device name.
	// Ex. For /dev/sda1, this field is "1", for /dev/sda, this field is 0 or empty.
	Partition int `json:"partition,omitempty"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`
}

// GitRepo represents a volume that is pulled from git when the pod is created.
type GitRepo struct {
	// Repository URL
	Repository string `json:"repository"`
	// Commit hash, this is optional
	Revision string `json:"revision"`
}

// Port represents a network port in a single container.
type Port struct {
	// Optional: If specified, this must be a DNS_LABEL.  Each named port
	// in a pod must have a unique name.
	Name string `json:"name,omitempty"`
	// Optional: If specified, this must be a valid port number, 0 < x < 65536.
	HostPort int `json:"hostPort,omitempty"`
	// Required: This must be a valid port number, 0 < x < 65536.
	ContainerPort int `json:"containerPort"`
	// Optional: Supports "TCP" and "UDP".  Defaults to "TCP".
	Protocol Protocol `json:"protocol,omitempty"`
	// Optional: What host IP to bind the external port to.
	HostIP string `json:"hostIP,omitempty"`
}

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// Required: This must match the Name of a Volume [above].
	Name string `json:"name"`
	// Optional: Defaults to false (read-write).
	ReadOnly bool `json:"readOnly,omitempty"`
	// Required.
	MountPath string `json:"mountPath,omitempty"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Required: This must be a C_IDENTIFIER.
	Name string `json:"name"`
	// Optional: defaults to "".
	Value string `json:"value,omitempty"`
}

// HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetAction struct {
	// Optional: Path to access on the HTTP server.
	Path string `json:"path,omitempty"`
	// Required: Name or number of the port to access on the container.
	Port util.IntOrString `json:"port,omitempty"`
	// Optional: Host name to connect to, defaults to the pod IP.
	Host string `json:"host,omitempty"`
}

// TCPSocketAction describes an action based on opening a socket
type TCPSocketAction struct {
	// Required: Port to connect to.
	Port util.IntOrString `json:"port,omitempty"`
}

// ExecAction describes a "run in container" action.
type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem.  The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work.  To use
	// a shell, you need to explicitly call out to that shell.
	Command []string `json:"command,omitempty"`
}

// LivenessProbe describes how to probe a container for liveness.
// TODO: pass structured data to the actions, and document that data here.
type LivenessProbe struct {
	// Type of liveness probe.  Current legal values "HTTP", "TCP", "Exec"
	Type string `json:"type,omitempty"`
	// HTTPGetProbe parameters, required if Type == 'HTTP'
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`
	// TCPSocketProbe parameter, required if Type == 'TCP'
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty"`
	// ExecProbe parameter, required if Type == 'Exec'
	Exec *ExecAction `json:"exec,omitempty"`
	// Length of time before health checking is activated.  In seconds.
	InitialDelaySeconds int64 `json:"initialDelaySeconds,omitempty"`
}

// PullPolicy describes a policy for if/when to pull a container image
type PullPolicy string

const (
	// PullAlways means that kubelet always attempts to pull the latest image.  Container will fail If the pull fails.
	PullAlways PullPolicy = "PullAlways"
	// PullNever means that kubelet never pulls an image, but only uses a local image.  Container will fail if the image isn't present
	PullNever PullPolicy = "PullNever"
	// PullIfNotPresent means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "PullIfNotPresent"
)

// Container represents a single container that is expected to be run on the host.
type Container struct {
	// Required: This must be a DNS_LABEL.  Each container in a pod must
	// have a unique name.
	Name string `json:"name"`
	// Required.
	Image string `json:"image"`
	// Optional: Defaults to whatever is defined in the image.
	Command []string `json:"command,omitempty"`
	// Optional: Defaults to Docker's default.
	WorkingDir string   `json:"workingDir,omitempty"`
	Ports      []Port   `json:"ports,omitempty"`
	Env        []EnvVar `json:"env,omitempty"`
	// Optional: Defaults to unlimited.
	Memory int `json:"memory,omitempty"`
	// Optional: Defaults to unlimited.
	CPU           int            `json:"cpu,omitempty"`
	VolumeMounts  []VolumeMount  `json:"volumeMounts,omitempty"`
	LivenessProbe *LivenessProbe `json:"livenessProbe,omitempty"`
	Lifecycle     *Lifecycle     `json:"lifecycle,omitempty"`
	// Optional: Defaults to /dev/termination-log
	TerminationMessagePath string `json:"terminationMessagePath,omitempty"`
	// Optional: Default to false.
	Privileged bool `json:"privileged,omitempty"`
	// Optional: Policy for pulling images for this container
	ImagePullPolicy PullPolicy `json:"imagePullPolicy"`
}

// Handler defines a specific action that should be taken
// TODO: pass structured data to these actions, and document that data here.
type Handler struct {
	// One and only one of the following should be specified.
	// Exec specifies the action to take.
	Exec *ExecAction `json:"exec,omitempty"`
	// HTTPGet specifies the http request to perform.
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty"`
}

// Lifecycle describes actions that the management system should take in response to container lifecycle
// events.  For the PostStart and PreStop lifecycle handlers, management of the container blocks
// until the action is complete, unless the container process fails, in which case the handler is aborted.
type Lifecycle struct {
	// PostStart is called immediately after a container is created.  If the handler fails, the container
	// is terminated and restarted.
	PostStart *Handler `json:"postStart,omitempty"`
	// PreStop is called immediately before a container is terminated.  The reason for termination is
	// passed to the handler.  Regardless of the outcome of the handler, the container is eventually terminated.
	PreStop *Handler `json:"preStop,omitempty"`
}

// PodPhase is a label for the condition of a pod at the current time.
type PodPhase string

// These are the valid states of pods.
const (
	// PodPending means the pod has been accepted by the system, but one or more of the containers
	// has not been started. This includes time before being bound to a node, as well as time spent
	// pulling images onto the host.
	PodPending PodPhase = "Pending"
	// PodRunning means the pod has been bound to a node and all of the containers have been started.
	// At least one container is still running or is in the process of being restarted.
	PodRunning PodPhase = "Running"
	// PodSucceeded means that all containers in the pod have voluntarily terminated
	// with a container exit code of 0, and the system is not going to restart any of these containers.
	PodSucceeded PodPhase = "Succeeded"
	// PodFailed means that all containers in the pod have terminated, and at least one container has
	// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
	PodFailed PodPhase = "Failed"
)

type ContainerStateWaiting struct {
	// Reason could be pulling image,
	Reason string `json:"reason,omitempty"`
}

type ContainerStateRunning struct {
	StartedAt time.Time `json:"startedAt,omitempty"`
}

type ContainerStateTerminated struct {
	ExitCode   int       `json:"exitCode"`
	Signal     int       `json:"signal,omitempty"`
	Reason     string    `json:"reason,omitempty"`
	Message    string    `json:"message,omitempty"`
	StartedAt  time.Time `json:"startedAt,omitempty"`
	FinishedAt time.Time `json:"finishedAt,omitempty"`
}

// ContainerState holds a possible state of container.
// Only one of its members may be specified.
// If none of them is specified, the default one is ContainerStateWaiting.
type ContainerState struct {
	Waiting     *ContainerStateWaiting    `json:"waiting,omitempty"`
	Running     *ContainerStateRunning    `json:"running,omitempty"`
	Termination *ContainerStateTerminated `json:"termination,omitempty"`
}

type ContainerStatus struct {
	// TODO(dchen1107): Should we rename PodStatus to a more generic name or have a separate states
	// defined for container?
	State ContainerState `json:"state,omitempty"`
	// Note that this is calculated from dead containers.  But those containers are subject to
	// garbage collection.  This value will get capped at 5 by GC.
	RestartCount int `json:"restartCount"`
	// TODO(dchen1107): Introduce our own NetworkSettings struct here?
	// TODO(dchen1107): Which image the container is running with?
}

// PodInfo contains one entry for every container with available info.
type PodInfo map[string]ContainerStatus

type RestartPolicyAlways struct{}

// TODO(dchen1107): Define what kinds of failures should restart.
// TODO(dchen1107): Decide whether to support policy knobs, and, if so, which ones.
type RestartPolicyOnFailure struct{}

type RestartPolicyNever struct{}

type RestartPolicy struct {
	// Only one of the following restart policies may be specified.
	// If none of the following policies is specified, the default one
	// is RestartPolicyAlways.
	Always    *RestartPolicyAlways    `json:"always,omitempty"`
	OnFailure *RestartPolicyOnFailure `json:"onFailure,omitempty"`
	Never     *RestartPolicyNever     `json:"never,omitempty"`
}

// PodSpec is a description of a pod
type PodSpec struct {
	Volumes       []Volume      `json:"volumes"`
	Containers    []Container   `json:"containers"`
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

// PodStatus represents information about the status of a pod. Status may trail the actual
// state of a system.
type PodStatus struct {
	Phase PodPhase `json:"phase,omitempty"`
	// A human readable message indicating details about why the pod is in this state.
	Message string `json:"message,omitempty"`

	// Host is the name of the node that this Pod is currently bound to, or empty if no
	// assignment has been done.
	Host   string `json:"host,omitempty"`
	HostIP string `json:"hostIP,omitempty"`
	PodIP  string `json:"podIP,omitempty"`

	// The key of this map is the *name* of the container within the manifest; it has one
	// entry per container in the manifest. The value of this map is currently the output
	// of `docker inspect`. This output format is *not* final and should not be relied
	// upon.
	// TODO: Make real decisions about what our info should look like. Re-enable fuzz test
	// when we have done this.
	Info PodInfo `json:"info,omitempty"`
}

// Pod is a collection of containers that can run on a host. This resource is created
// by clients and scheduled onto hosts.  BoundPod represents the state of this resource
// to hosts.
type Pod struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod.
	Spec PodSpec `json:"spec,omitempty"`

	// Status represents the current information about a pod. This data may not be up
	// to date.
	Status PodStatus `json:"status,omitempty"`
}

// PodList is a list of Pods.
type PodList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Pod `json:"items"`
}

// PodTemplateSpec describes the data a pod should have when created from a template
type PodTemplateSpec struct {
	// Metadata of the pods created from this template.
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod.
	Spec PodSpec `json:"spec,omitempty"`
}

// PodTemplate describes a template for creating copies of a predefined pod.
type PodTemplate struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod.
	Spec PodTemplateSpec `json:"spec,omitempty"`
}

// PodTemplateList is a list of PodTemplates.
type PodTemplateList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []PodTemplate `json:"items"`
}

// BoundPod is a collection of containers that should be run on a host. A BoundPod
// defines how a Pod may change after a Binding is created. A Pod is a request to
// execute a pod, whereas a BoundPod is the specification that would be run on a server.
type BoundPod struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod.
	Spec PodSpec `json:"spec,omitempty"`
}

// BoundPods is a list of Pods bound to a common server. The resource version of
// the pod list is guaranteed to only change when the list of bound pods changes.
type BoundPods struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Host is the name of a node that these pods were bound to.
	Host string `json:"host"`

	// Items is the list of all pods bound to a given host.
	Items []BoundPod `json:"items"`
}

// ReplicationControllerSpec is the specification of a replication controller.
type ReplicationControllerSpec struct {
	// Replicas is the number of desired replicas.
	Replicas int `json:"replicas"`

	// Selector is a label query over pods that should match the Replicas count.
	Selector map[string]string `json:"selector,omitempty"`

	// Template is a reference to an object that describes the pod that will be created if
	// insufficient replicas are detected.
	Template ObjectReference `json:"template,omitempty"`
}

// ReplicationControllerStatus represents the current status of a replication
// controller.
type ReplicationControllerStatus struct {
	// Replicas is the number of actual replicas.
	Replicas int `json:"replicas"`
}

// ReplicationController represents the configuration of a replication controller.
type ReplicationController struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired behavior of this replication controller.
	Spec ReplicationControllerSpec `json:"spec,omitempty"`

	// Status is the current status of this replication controller. This data may be
	// out of date by some window of time.
	Status ReplicationControllerStatus `json:"status,omitempty"`
}

// ReplicationControllerList is a collection of replication controllers.
type ReplicationControllerList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []ReplicationController `json:"items"`
}

// ServiceStatus represents the current status of a service
type ServiceStatus struct{}

// ServiceSpec describes the attributes that a user creates on a service
type ServiceSpec struct {
	// Port is the TCP or UDP port that will be made available to each pod for connecting to the pods
	// proxied by this service.
	Port int `json:"port"`

	// Optional: Supports "TCP" and "UDP".  Defaults to "TCP".
	Protocol Protocol `json:"protocol,omitempty"`

	// This service will route traffic to pods having labels matching this selector. If null, no endpoints will be automatically created. If empty, all pods will be selected.
	Selector map[string]string `json:"selector"`

	// PortalIP is usually assigned by the master.  If specified by the user
	// we will try to respect it or else fail the request.  This field can
	// not be changed by updates.
	PortalIP string `json:"portalIP,omitempty"`

	// ProxyPort is assigned by the master.  If 0, the proxy will choose an ephemeral port.
	ProxyPort int `json:"proxyPort,omitempty"`

	// CreateExternalLoadBalancer indicates whether a load balancer should be created for this service.
	CreateExternalLoadBalancer bool `json:"createExternalLoadBalancer,omitempty"`
	// PublicIPs are used by external load balancers.
	PublicIPs []string `json:"publicIPs,omitempty"`

	// ContainerPort is the name of the port on the container to direct traffic to.
	// Optional, if unspecified use the first port on the container.
	ContainerPort util.IntOrString `json:"containerPort,omitempty"`
}

// Service is a named abstraction of software service (for example, mysql) consisting of local port
// (for example 3306) that the proxy listens on, and the selector that determines which pods
// will answer requests sent through the proxy.
type Service struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a service.
	Spec ServiceSpec `json:"spec,omitempty"`

	// Status represents the current status of a service.
	Status ServiceStatus `json:"status,omitempty"`
}

// ServiceList holds a list of services.
type ServiceList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Service `json:"items"`
}

// Endpoints is a collection of endpoints that implement the actual service, for example:
// Name: "mysql", Endpoints: ["10.10.1.1:1909", "10.10.2.2:8834"]
type Endpoints struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`

	// Endpoints is the list of host ports that satisfy the service selector
	Endpoints []string `json:"endpoints"`
}

// EndpointsList is a list of endpoints.
type EndpointsList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Endpoints `json:"items"`
}

// NodeSpec describes the attributes that a node is created with.
type NodeSpec struct {
	// Capacity represents the available resources of a node
	// see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/resources.md for more details.
	Capacity ResourceList `json:"capacity,omitempty"`
}

// NodeStatus is information about the current status of a node.
type NodeStatus struct {
	// Queried from cloud provider, if available.
	HostIP string `json:"hostIP,omitempty"`
}

type ResourceName string

type ResourceList map[ResourceName]util.IntOrString

// Node is a worker node in Kubernetes.
// The name of the node according to etcd is in ID.
type Node struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a node.
	Spec NodeSpec `json:"spec,omitempty"`

	// Status describes the current status of a Node
	Status NodeStatus `json:"status,omitempty"`
}

// NodeList is a list of minions.
type NodeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Node `json:"items"`
}

// Binding is written by a scheduler to cause a pod to be bound to a node. Name is not
// required for Bindings.
type Binding struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	// PodID is a Pod name to be bound to a node.
	PodID string `json:"podID"`
	// Host is the name of a node to bind to.
	Host string `json:"host"`
}

// Status is a return value for calls that don't return other objects.
type Status struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	// One of: "Success", "Failure", "Working" (for operations not yet completed)
	Status string `json:"status,omitempty"`
	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty"`
	// A machine-readable description of why this operation is in the
	// "Failure" or "Working" status. If this value is empty there
	// is no information available. A Reason clarifies an HTTP status
	// code but does not override it.
	Reason StatusReason `json:"reason,omitempty"`
	// Extended data associated with the reason.  Each reason may define its
	// own extended details. This field is optional and the data returned
	// is not guaranteed to conform to any schema except that defined by
	// the reason type.
	Details *StatusDetails `json:"details,omitempty"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code int `json:"code,omitempty"`
}

// StatusDetails is a set of additional properties that MAY be set by the
// server to provide additional information about a response. The Reason
// field of a Status object defines what attributes will be set. Clients
// must ignore fields that do not match the defined type of each attribute,
// and should assume that any attribute may be empty, invalid, or under
// defined.
type StatusDetails struct {
	// The ID attribute of the resource associated with the status StatusReason
	// (when there is a single ID which can be described).
	ID string `json:"id,omitempty"`
	// The kind attribute of the resource associated with the status StatusReason.
	// On some operations may differ from the requested resource Kind.
	Kind string `json:"kind,omitempty"`
	// The Causes array includes more details associated with the StatusReason
	// failure. Not all StatusReasons may provide detailed causes.
	Causes []StatusCause `json:"causes,omitempty"`
}

// Values of Status.Status
const (
	StatusSuccess = "Success"
	StatusFailure = "Failure"
	StatusWorking = "Working"
)

// StatusReason is an enumeration of possible failure causes.  Each StatusReason
// must map to a single HTTP status code, but multiple reasons may map
// to the same HTTP status code.
// TODO: move to apiserver
type StatusReason string

const (
	// StatusReasonUnknown means the server has declined to indicate a specific reason.
	// The details field may contain other information about this error.
	// Status code 500.
	StatusReasonUnknown StatusReason = ""

	// StatusReasonWorking means the server is processing this request and will complete
	// at a future time.
	// Details (optional):
	//   "kind" string - the name of the resource being referenced ("operation" today)
	//   "id"   string - the identifier of the Operation resource where updates
	//                   will be returned
	// Headers (optional):
	//   "Location" - HTTP header populated with a URL that can retrieved the final
	//                status of this operation.
	// Status code 202
	StatusReasonWorking StatusReason = "Working"

	// StatusReasonNotFound means one or more resources required for this operation
	// could not be found.
	// Details (optional):
	//   "kind" string - the kind attribute of the missing resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the missing resource
	// Status code 404
	StatusReasonNotFound StatusReason = "NotFound"

	// StatusReasonAlreadyExists means the resource you are creating already exists.
	// Details (optional):
	//   "kind" string - the kind attribute of the conflicting resource
	//   "id"   string - the identifier of the conflicting resource
	// Status code 409
	StatusReasonAlreadyExists StatusReason = "AlreadyExists"

	// StatusReasonConflict means the requested update operation cannot be completed
	// due to a conflict in the operation. The client may need to alter the request.
	// Each resource may define custom details that indicate the nature of the
	// conflict.
	// Status code 409
	StatusReasonConflict StatusReason = "Conflict"

	// StatusReasonInvalid means the requested create or update operation cannot be
	// completed due to invalid data provided as part of the request. The client may
	// need to alter the request. When set, the client may use the StatusDetails
	// message field as a summary of the issues encountered.
	// Details (optional):
	//   "kind" string - the kind attribute of the invalid resource
	//   "id"   string - the identifier of the invalid resource
	//   "causes"      - one or more StatusCause entries indicating the data in the
	//                   provided resource that was invalid.  The code, message, and
	//                   field attributes will be set.
	// Status code 422
	StatusReasonInvalid StatusReason = "Invalid"
)

// StatusCause provides more information about an api.Status failure, including
// cases when multiple errors are encountered.
type StatusCause struct {
	// A machine-readable description of the cause of the error. If this value is
	// empty there is no information available.
	Type CauseType `json:"reason,omitempty"`
	// A human-readable description of the cause of the error.  This field may be
	// presented as-is to a reader.
	Message string `json:"message,omitempty"`
	// The field of the resource that has caused this error, as named by its JSON
	// serialization. May include dot and postfix notation for nested attributes.
	// Arrays are zero-indexed.  Fields may appear more than once in an array of
	// causes due to fields having multiple errors.
	// Optional.
	//
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	Field string `json:"field,omitempty"`
}

// CauseType is a machine readable value providing more detail about what
// occured in a status response. An operation may have multiple causes for a
// status (whether Failure, Success, or Working).
type CauseType string

const (
	// CauseTypeFieldValueNotFound is used to report failure to find a requested value
	// (e.g. looking up an ID).
	CauseTypeFieldValueNotFound CauseType = "FieldValueNotFound"
	// CauseTypeFieldValueRequired is used to report required values that are not
	// provided (e.g. empty strings, null values, or empty arrays).
	CauseTypeFieldValueRequired CauseType = "FieldValueRequired"
	// CauseTypeFieldValueDuplicate is used to report collisions of values that must be
	// unique (e.g. unique IDs).
	CauseTypeFieldValueDuplicate CauseType = "FieldValueDuplicate"
	// CauseTypeFieldValueInvalid is used to report malformed values (e.g. failed regex
	// match).
	CauseTypeFieldValueInvalid CauseType = "FieldValueInvalid"
	// CauseTypeFieldValueNotSupported is used to report valid (as per formatting rules)
	// values that can not be handled (e.g. an enumerated string).
	CauseTypeFieldValueNotSupported CauseType = "FieldValueNotSupported"
)

// Operation is a request from a client that has not yet been satisfied. The name of an
// Operation is assigned by the server when an operation is started, and can be used by
// clients to retrieve the final result of the operation at a later time.
type Operation struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`
}

// OperationList is a list of operations, as delivered to API clients.
type OperationList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Operation `json:"items"`
}

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	Kind            string `json:"kind,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Name            string `json:"name,omitempty"`
	UID             string `json:"uid,omitempty"`
	APIVersion      string `json:"apiVersion,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// Optional. If referring to a piece of an object instead of an entire object, this string
	// should contain a valid field access statement. For example,
	// if the object reference is to a container within a pod, this would take on a value like:
	// "spec.containers[2]". Such statements are valid language constructs in
	// both go and JavaScript. This is syntax is chosen only to have some well-defined way of
	// referencing a part of an object.
	// TODO: this design is not final and this field is subject to change in the future.
	FieldPath string `json:"fieldPath,omitempty"`
}

// Event is a report of an event somewhere in the cluster.
// TODO: Decide whether to store these separately or with the object they apply to.
type Event struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`

	// Required. The object that this event is about.
	InvolvedObject ObjectReference `json:"involvedObject,omitempty"`

	// Should be a short, machine understandable string that describes the current status
	// of the referred object. This should not give the reason for being in this state.
	// Examples: "Running", "CantStart", "CantSchedule", "Deleted".
	// It's OK for components to make up statuses to report here, but the same string should
	// always be used for the same status.
	// TODO: define a way of making sure these are consistent and don't collide.
	// TODO: provide exact specification for format.
	Condition string `json:"condition,omitempty"`

	// Optional; this should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status. For example, if ObjectStatus is
	// "cantStart", StatusReason might be "imageNotFound".
	// TODO: provide exact specification for format.
	Reason string `json:"reason,omitempty"`

	// Optional. A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	Message string `json:"message,omitempty"`

	// Optional. The component reporting this event. Should be a short machine understandable string.
	// TODO: provide exact specification for format.
	Source string `json:"source,omitempty"`

	// The time at which the client recorded the event. (Time of server receipt is in TypeMeta.)
	Timestamp util.Time `json:"timestamp,omitempty"`
}

// EventList is a list of events.
type EventList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`

	Items []Event `json:"items"`
}
