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

type ClusterProfilePackManifestObservation struct {
}

type ClusterProfilePackManifestParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type EksBackupPolicyObservation struct {
}

type EksBackupPolicyParameters struct {

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

type EksCloudConfigObservation struct {
}

type EksCloudConfigParameters struct {

	// +kubebuilder:validation:Optional
	AzSubnets map[string]*string `json:"azSubnets,omitempty" tf:"az_subnets,omitempty"`

	// +kubebuilder:validation:Optional
	Azs []*string `json:"azs,omitempty" tf:"azs,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointAccess *string `json:"endpointAccess,omitempty" tf:"endpoint_access,omitempty"`

	// +kubebuilder:validation:Optional
	PublicAccessCidrs []*string `json:"publicAccessCidrs,omitempty" tf:"public_access_cidrs,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SSHKeyName *string `json:"sshKeyName,omitempty" tf:"ssh_key_name,omitempty"`

	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type EksClusterProfileObservation struct {
}

type EksClusterProfilePackObservation struct {
}

type EksClusterProfilePackParameters struct {

	// +kubebuilder:validation:Optional
	Manifest []ClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type EksClusterProfileParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []EksClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type EksClusterRbacBindingObservation struct {
}

type EksClusterRbacBindingParameters struct {

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []EksClusterRbacBindingSubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EksClusterRbacBindingSubjectsObservation struct {
}

type EksClusterRbacBindingSubjectsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type EksMachinePoolObservation struct {
}

type EksMachinePoolParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLabels map[string]*string `json:"additionalLabels,omitempty" tf:"additional_labels,omitempty"`

	// +kubebuilder:validation:Optional
	AzSubnets map[string]*string `json:"azSubnets,omitempty" tf:"az_subnets,omitempty"`

	// +kubebuilder:validation:Optional
	Azs []*string `json:"azs,omitempty" tf:"azs,omitempty"`

	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type,omitempty"`

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// +kubebuilder:validation:Optional
	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price,omitempty"`

	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []EksMachinePoolTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateStrategy *string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
}

type EksMachinePoolTaintsObservation struct {
}

type EksMachinePoolTaintsParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type EksNamespacesObservation struct {
}

type EksNamespacesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type EksObservation struct {
	CloudConfigID *string `json:"cloudConfigId,omitempty" tf:"cloud_config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kubeconfig *string `json:"kubeconfig,omitempty" tf:"kubeconfig,omitempty"`
}

type EksPackObservation struct {
}

type EksPackParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type EksParameters struct {

	// +kubebuilder:validation:Optional
	BackupPolicy []EksBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountID *string `json:"cloudAccountId" tf:"cloud_account_id,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig []EksCloudConfigParameters `json:"cloudConfig" tf:"cloud_config,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []EksClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfileID *string `json:"clusterProfileId,omitempty" tf:"cluster_profile_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []EksClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Optional
	FargateProfile []FargateProfileParameters `json:"fargateProfile,omitempty" tf:"fargate_profile,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool []EksMachinePoolParameters `json:"machinePool" tf:"machine_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []EksNamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []EksPackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// +kubebuilder:validation:Optional
	ScanPolicy []EksScanPolicyParameters `json:"scanPolicy,omitempty" tf:"scan_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EksScanPolicyObservation struct {
}

type EksScanPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationScanSchedule *string `json:"configurationScanSchedule" tf:"configuration_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	ConformanceScanSchedule *string `json:"conformanceScanSchedule" tf:"conformance_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	PenetrationScanSchedule *string `json:"penetrationScanSchedule" tf:"penetration_scan_schedule,omitempty"`
}

type FargateProfileObservation struct {
}

type FargateProfileParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalTags map[string]*string `json:"additionalTags,omitempty" tf:"additional_tags,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Selector []SelectorParameters `json:"selector" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

type SelectorObservation struct {
}

type SelectorParameters struct {

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

// EksSpec defines the desired state of Eks
type EksSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EksParameters `json:"forProvider"`
}

// EksStatus defines the observed state of Eks.
type EksStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Eks is the Schema for the Ekss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palettejet}
type Eks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksSpec   `json:"spec"`
	Status            EksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksList contains a list of Ekss
type EksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eks `json:"items"`
}

// Repository type metadata.
var (
	Eks_Kind             = "Eks"
	Eks_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Eks_Kind}.String()
	Eks_KindAPIVersion   = Eks_Kind + "." + CRDGroupVersion.String()
	Eks_GroupVersionKind = CRDGroupVersion.WithKind(Eks_Kind)
)

func init() {
	SchemeBuilder.Register(&Eks{}, &EksList{})
}