// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfos


type DataDatabricksPolicyInfosPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/policy_infos#name DataDatabricksPolicyInfos#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/policy_infos#on_securable_fullname DataDatabricksPolicyInfos#on_securable_fullname}.
	OnSecurableFullname *string `field:"required" json:"onSecurableFullname" yaml:"onSecurableFullname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/data-sources/policy_infos#on_securable_type DataDatabricksPolicyInfos#on_securable_type}.
	OnSecurableType *string `field:"required" json:"onSecurableType" yaml:"onSecurableType"`
}

