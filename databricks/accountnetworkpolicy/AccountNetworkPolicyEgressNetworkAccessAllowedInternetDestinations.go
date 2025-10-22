// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgressNetworkAccessAllowedInternetDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/account_network_policy#destination AccountNetworkPolicy#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/account_network_policy#internet_destination_type AccountNetworkPolicy#internet_destination_type}.
	InternetDestinationType *string `field:"optional" json:"internetDestinationType" yaml:"internetDestinationType"`
}

