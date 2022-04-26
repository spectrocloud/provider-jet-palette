/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterRbacBindingSubjectsObservation struct {
}

type ClusterRbacBindingSubjectsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EdgeVsphereBackupPolicyObservation struct {
}

type EdgeVsphereBackupPolicyParameters struct {

	// +kubebuilder:validation:Required
	BackupLocationID *string `json:"backupLocationId" tf:"backup_location_id,omitempty"`

	// +kubebuilder:validation:Required
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type EdgeVsphereCloudConfigObservation struct {
}

type EdgeVsphereCloudConfigParameters struct {

	// +kubebuilder:validation:Required
	Datacenter *string `json:"datacenter" tf:"datacenter,omitempty"`

	// +kubebuilder:validation:Required
	Folder *string `json:"folder" tf:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkSearchDomain *string `json:"networkSearchDomain,omitempty" tf:"network_search_domain,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Optional
	SSHKey *string `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`

	// +kubebuilder:validation:Optional
	StaticIP *bool `json:"staticIp,omitempty" tf:"static_ip,omitempty"`

	// +kubebuilder:validation:Required
	Vip *string `json:"vip" tf:"vip,omitempty"`
}

type EdgeVsphereClusterProfileObservation struct {
}

type EdgeVsphereClusterProfilePackObservation struct {
}

type EdgeVsphereClusterProfilePackParameters struct {

	// +kubebuilder:validation:Optional
	Manifest []PackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type EdgeVsphereClusterProfileParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []EdgeVsphereClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type EdgeVsphereClusterRbacBindingObservation struct {
}

type EdgeVsphereClusterRbacBindingParameters struct {

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []ClusterRbacBindingSubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EdgeVsphereMachinePoolObservation struct {
}

type EdgeVsphereMachinePoolParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLabels map[string]*string `json:"additionalLabels,omitempty" tf:"additional_labels,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlane *bool `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlaneAsWorker *bool `json:"controlPlaneAsWorker,omitempty" tf:"control_plane_as_worker,omitempty"`

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType []InstanceTypeParameters `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Placement []PlacementParameters `json:"placement" tf:"placement,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []MachinePoolTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateStrategy *string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
}

type EdgeVsphereNamespacesObservation struct {
}

type EdgeVsphereNamespacesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type EdgeVsphereObservation struct {
	CloudConfigID *string `json:"cloudConfigId,omitempty" tf:"cloud_config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kubeconfig *string `json:"kubeconfig,omitempty" tf:"kubeconfig,omitempty"`
}

type EdgeVspherePackObservation struct {
}

type EdgeVspherePackParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type EdgeVsphereParameters struct {

	// +kubebuilder:validation:Optional
	BackupPolicy []EdgeVsphereBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig []EdgeVsphereCloudConfigParameters `json:"cloudConfig" tf:"cloud_config,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []EdgeVsphereClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []EdgeVsphereClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Required
	EdgeHostUID *string `json:"edgeHostUid" tf:"edge_host_uid,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool []EdgeVsphereMachinePoolParameters `json:"machinePool" tf:"machine_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []EdgeVsphereNamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchAfter *string `json:"osPatchAfter,omitempty" tf:"os_patch_after,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchOnBoot *bool `json:"osPatchOnBoot,omitempty" tf:"os_patch_on_boot,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchSchedule *string `json:"osPatchSchedule,omitempty" tf:"os_patch_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []EdgeVspherePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// +kubebuilder:validation:Optional
	ScanPolicy []EdgeVsphereScanPolicyParameters `json:"scanPolicy,omitempty" tf:"scan_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EdgeVsphereScanPolicyObservation struct {
}

type EdgeVsphereScanPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationScanSchedule *string `json:"configurationScanSchedule" tf:"configuration_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	ConformanceScanSchedule *string `json:"conformanceScanSchedule" tf:"conformance_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	PenetrationScanSchedule *string `json:"penetrationScanSchedule" tf:"penetration_scan_schedule,omitempty"`
}

type InstanceTypeObservation struct {
}

type InstanceTypeParameters struct {

	// +kubebuilder:validation:Required
	CPU *float64 `json:"cpu" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	MemoryMb *float64 `json:"memoryMb" tf:"memory_mb,omitempty"`
}

type MachinePoolTaintsObservation struct {
}

type MachinePoolTaintsParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PackManifestObservation struct {
}

type PackManifestParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type PlacementObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PlacementParameters struct {

	// +kubebuilder:validation:Required
	Cluster *string `json:"cluster" tf:"cluster,omitempty"`

	// +kubebuilder:validation:Required
	Datastore *string `json:"datastore" tf:"datastore,omitempty"`

	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePool *string `json:"resourcePool" tf:"resource_pool,omitempty"`

	// +kubebuilder:validation:Optional
	StaticIPPoolID *string `json:"staticIpPoolId,omitempty" tf:"static_ip_pool_id,omitempty"`
}

// EdgeVsphereSpec defines the desired state of EdgeVsphere
type EdgeVsphereSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeVsphereParameters `json:"forProvider"`
}

// EdgeVsphereStatus defines the observed state of EdgeVsphere.
type EdgeVsphereStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeVsphereObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeVsphere is the Schema for the EdgeVspheres API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palettejet}
type EdgeVsphere struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeVsphereSpec   `json:"spec"`
	Status            EdgeVsphereStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeVsphereList contains a list of EdgeVspheres
type EdgeVsphereList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeVsphere `json:"items"`
}

// Repository type metadata.
var (
	EdgeVsphere_Kind             = "EdgeVsphere"
	EdgeVsphere_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeVsphere_Kind}.String()
	EdgeVsphere_KindAPIVersion   = EdgeVsphere_Kind + "." + CRDGroupVersion.String()
	EdgeVsphere_GroupVersionKind = CRDGroupVersion.WithKind(EdgeVsphere_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeVsphere{}, &EdgeVsphereList{})
}