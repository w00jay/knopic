/*
Linstor Operator
Copyright 2019 LINBIT USA, LLC.

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

package spec

import corev1 "k8s.io/api/core/v1"

// Shared consts common to container volumes.
const (
	DevDir             = "/dev/"
	DevDirName         = "device-dir"
	LinstorConfDir     = "/etc/linstor"
	LinstorConfDirName = "linstor-conf"
	ModulesDir         = "/lib/modules/" // "/usr/lib/modules/"
	ModulesDirName     = "modules-dir"
	SrcDir             = "/usr/src"
	SrcDirName         = "src-dir"
	UdevDir            = "/run/udev"
	UdevDirName        = "udev"
)

// Shared consts common to container volumes. These need to be vars, so they
// are addressible.
var (
	HostPathDirectoryType         = corev1.HostPathDirectory
	HostPathDirectoryOrCreateType = corev1.HostPathDirectoryOrCreate
	MountPropagationBidirectional = corev1.MountPropagationBidirectional
)

// Shared consts common to container security. These need to be vars, so they
// are addressible.
var (
	Privileged = true
)

const selectorPrefix = "linstor.linbit.com/"

// Kubernetes node labels that are an opt-in selector to run linstor pods when
// set to "true".
const (
	// LinstorSatelliteNode label to mark node eligible to run linstor-node pods.
	LinstorNode = selectorPrefix + "linstor-node"
)

// LinstorPriorityClassName is the name of the PriorityClass set up in the
// example yaml used by important linstor components.
const LinstorCSPriorityClassName = "linstor-cs-priority-class"
const LinstorNSPriorityClassName = "linstor-ns-priority-class"

const (
	// LinstorServerImage is the repo/tag for linstor-server. The controller
	// and node use the same image with different commands.
	//LinstorServerImage = "quay.io/piraeusdatastore/piraeus-server"
	LinstorServerImage = "woojay/my-piraeus-server"
	// LinstorVersion must must match exactly in the ControllerSet and the NodeSet
	// since the linstor controller and satellite versions must also match exactly.
	LinstorVersion = "v1.2.1"
	// LinstorKernelModImage is the worker (aka satellite) image for each node
	// LinstorKernelModImage = "quay.io/piraeusdatastore/drbd9-centos7"
	LinstorKernelModImage = "woojay/my-drbd"
	// LinstorKernelModVersion is the release tag for the above image
	LinstorKernelModVersion = "v9.0.21"
)
