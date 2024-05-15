// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesExternalCustomerInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_workspaces#authoritative_user_email MwsWorkspaces#authoritative_user_email}.
	AuthoritativeUserEmail *string `field:"required" json:"authoritativeUserEmail" yaml:"authoritativeUserEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_workspaces#authoritative_user_full_name MwsWorkspaces#authoritative_user_full_name}.
	AuthoritativeUserFullName *string `field:"required" json:"authoritativeUserFullName" yaml:"authoritativeUserFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.43.0/docs/resources/mws_workspaces#customer_name MwsWorkspaces#customer_name}.
	CustomerName *string `field:"required" json:"customerName" yaml:"customerName"`
}

