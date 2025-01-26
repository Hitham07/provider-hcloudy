// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetInitParameters struct {

	// Range to allocate IPs from. Must be a subnet of the ip_range of the Network and must not overlap with any other subnets or with any destinations in routes.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// ID of the Network the subnet should be added to.
	// +crossplane:generate:reference:type=github.com/Hitham07/provider-hcloudy/apis/hcloud/v1alpha1.Network
	NetworkID *float64 `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in hcloud to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in hcloud to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Name of network zone.
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	// Type of subnet. server, cloud or vswitch
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// ID of the vswitch, Required if type is vswitch
	VswitchID *float64 `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

type SubnetObservation struct {
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// (string) ID of the Network subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Range to allocate IPs from. Must be a subnet of the ip_range of the Network and must not overlap with any other subnets or with any destinations in routes.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// ID of the Network the subnet should be added to.
	NetworkID *float64 `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Name of network zone.
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	// Type of subnet. server, cloud or vswitch
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// ID of the vswitch, Required if type is vswitch
	VswitchID *float64 `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

type SubnetParameters struct {

	// Range to allocate IPs from. Must be a subnet of the ip_range of the Network and must not overlap with any other subnets or with any destinations in routes.
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// ID of the Network the subnet should be added to.
	// +crossplane:generate:reference:type=github.com/Hitham07/provider-hcloudy/apis/hcloud/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *float64 `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in hcloud to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in hcloud to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Name of network zone.
	// +kubebuilder:validation:Optional
	NetworkZone *string `json:"networkZone,omitempty" tf:"network_zone,omitempty"`

	// Type of subnet. server, cloud or vswitch
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// ID of the vswitch, Required if type is vswitch
	// +kubebuilder:validation:Optional
	VswitchID *float64 `json:"vswitchId,omitempty" tf:"vswitch_id,omitempty"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SubnetInitParameters `json:"initProvider,omitempty"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Subnet is the Schema for the Subnets API. Provides a Hetzner Cloud Network Subnet to represent a Subnet in the Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloudy}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipRange) || (has(self.initProvider) && has(self.initProvider.ipRange))",message="spec.forProvider.ipRange is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkZone) || (has(self.initProvider) && has(self.initProvider.networkZone))",message="spec.forProvider.networkZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
