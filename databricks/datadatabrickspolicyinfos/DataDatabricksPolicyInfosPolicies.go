// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfos


type DataDatabricksPolicyInfosPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#for_securable_type DataDatabricksPolicyInfos#for_securable_type}.
	ForSecurableType *string `field:"required" json:"forSecurableType" yaml:"forSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#policy_type DataDatabricksPolicyInfos#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#to_principals DataDatabricksPolicyInfos#to_principals}.
	ToPrincipals *[]*string `field:"required" json:"toPrincipals" yaml:"toPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#column_mask DataDatabricksPolicyInfos#column_mask}.
	ColumnMask *DataDatabricksPolicyInfosPoliciesColumnMask `field:"optional" json:"columnMask" yaml:"columnMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#comment DataDatabricksPolicyInfos#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#except_principals DataDatabricksPolicyInfos#except_principals}.
	ExceptPrincipals *[]*string `field:"optional" json:"exceptPrincipals" yaml:"exceptPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#match_columns DataDatabricksPolicyInfos#match_columns}.
	MatchColumns interface{} `field:"optional" json:"matchColumns" yaml:"matchColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#name DataDatabricksPolicyInfos#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#on_securable_fullname DataDatabricksPolicyInfos#on_securable_fullname}.
	OnSecurableFullname *string `field:"optional" json:"onSecurableFullname" yaml:"onSecurableFullname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#on_securable_type DataDatabricksPolicyInfos#on_securable_type}.
	OnSecurableType *string `field:"optional" json:"onSecurableType" yaml:"onSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#row_filter DataDatabricksPolicyInfos#row_filter}.
	RowFilter *DataDatabricksPolicyInfosPoliciesRowFilter `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/policy_infos#when_condition DataDatabricksPolicyInfos#when_condition}.
	WhenCondition *string `field:"optional" json:"whenCondition" yaml:"whenCondition"`
}

