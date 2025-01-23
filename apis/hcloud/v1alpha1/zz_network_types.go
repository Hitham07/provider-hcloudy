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

type NetworkInitParameters struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	ExposeRoutesToVswitch *bool `json:"exposeRoutesToVswitch,omitempty" tf:"expose_routes_to_vswitch,omitempty"`

	// IP Range of the whole Network which must span all included subnets and route destinations. Must be one of the private ipv4 ranges of RFC1918.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type NetworkObservation struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	ExposeRoutesToVswitch *bool `json:"exposeRoutesToVswitch,omitempty" tf:"expose_routes_to_vswitch,omitempty"`

	// (int) Unique ID of the network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP Range of the whole Network which must span all included subnets and route destinations. Must be one of the private ipv4 ranges of RFC1918.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type NetworkParameters struct {

	// Enable or disable delete protection. See "Delete Protection" in the Provider Docs for details.
	// +kubebuilder:validation:Optional
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	// Enable or disable exposing the routes to the vSwitch connection. The exposing only takes effect if a vSwitch connection is active.
	// +kubebuilder:validation:Optional
	ExposeRoutesToVswitch *bool `json:"exposeRoutesToVswitch,omitempty" tf:"expose_routes_to_vswitch,omitempty"`

	// IP Range of the whole Network which must span all included subnets and route destinations. Must be one of the private ipv4 ranges of RFC1918.
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
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
	InitProvider NetworkInitParameters `json:"initProvider,omitempty"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Network is the Schema for the Networks API. Provides a Hetzner Cloud Network to represent a Network in the Hetzner Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloudy}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipRange) || (has(self.initProvider) && has(self.initProvider.ipRange))",message="spec.forProvider.ipRange is a required parameter"
	Spec   NetworkSpec   `json:"spec"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
