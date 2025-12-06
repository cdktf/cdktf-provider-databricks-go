// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedInternetDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/account_network_policies#destination DataDatabricksAccountNetworkPolicies#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/account_network_policies#internet_destination_type DataDatabricksAccountNetworkPolicies#internet_destination_type}.
	InternetDestinationType *string `field:"optional" json:"internetDestinationType" yaml:"internetDestinationType"`
}

