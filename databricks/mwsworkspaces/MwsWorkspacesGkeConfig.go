// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesGkeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/mws_workspaces#connectivity_type MwsWorkspaces#connectivity_type}.
	ConnectivityType *string `field:"optional" json:"connectivityType" yaml:"connectivityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/mws_workspaces#master_ip_range MwsWorkspaces#master_ip_range}.
	MasterIpRange *string `field:"optional" json:"masterIpRange" yaml:"masterIpRange"`
}

