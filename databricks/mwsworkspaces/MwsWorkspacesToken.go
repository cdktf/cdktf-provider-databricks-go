// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesToken struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/mws_workspaces#comment MwsWorkspaces#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/mws_workspaces#lifetime_seconds MwsWorkspaces#lifetime_seconds}.
	LifetimeSeconds *float64 `field:"optional" json:"lifetimeSeconds" yaml:"lifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/mws_workspaces#token_id MwsWorkspaces#token_id}.
	TokenId *string `field:"optional" json:"tokenId" yaml:"tokenId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/resources/mws_workspaces#token_value MwsWorkspaces#token_value}.
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}

