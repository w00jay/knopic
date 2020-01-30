/*
Knopic Operator
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

// Kubernetes node labels that are an opt-in selector to run knopic pods when
// set to "true".
const (
	// KnopicSatelliteNode label to mark node eligible to run knopic-node pods.
	KnopicNode = selectorPrefix + "knopic-node"
)

// KnopicPriorityClassName is the name of the PriorityClass set up in the
// example yaml used by important knopic components.
const KnopicCSPriorityClassName = "knopic-cs-priority-class"
const KnopicNSPriorityClassName = "knopic-ns-priority-class"

const (
	// KnopicServerImage is the repo/tag for knopic-server. The controller
	// and node use the same image with different commands.
	//KnopicServerImage = "quay.io/piraeusdatastore/piraeus-server"
	KnopicServerImage = "woojay/my-piraeus-server"
	// KnopicVersion must must match exactly in the ControllerSet and the NodeSet
	// since the linstor controller and satellite versions must also match exactly.
	KnopicVersion = "v1.2.1"
	// KnopicKernelModImage is the worker (aka satellite) image for each node
	// KnopicKernelModImage = "quay.io/piraeusdatastore/drbd9-centos7"
	KnopicKernelModImage = "woojay/my-drbd"
	// KnopicKernelModVersion is the release tag for the above image
	KnopicKernelModVersion = "v9.0.21"
)
