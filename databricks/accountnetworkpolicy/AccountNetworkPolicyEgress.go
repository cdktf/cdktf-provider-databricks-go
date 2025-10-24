// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/account_network_policy#network_access AccountNetworkPolicy#network_access}.
	NetworkAccess *AccountNetworkPolicyEgressNetworkAccess `field:"optional" json:"networkAccess" yaml:"networkAccess"`
}

