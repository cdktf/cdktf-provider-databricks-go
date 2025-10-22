// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickspolicyinfo

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksPolicyInfoConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#for_securable_type DataDatabricksPolicyInfo#for_securable_type}.
	ForSecurableType *string `field:"required" json:"forSecurableType" yaml:"forSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#policy_type DataDatabricksPolicyInfo#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#to_principals DataDatabricksPolicyInfo#to_principals}.
	ToPrincipals *[]*string `field:"required" json:"toPrincipals" yaml:"toPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#column_mask DataDatabricksPolicyInfo#column_mask}.
	ColumnMask *DataDatabricksPolicyInfoColumnMask `field:"optional" json:"columnMask" yaml:"columnMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#comment DataDatabricksPolicyInfo#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#except_principals DataDatabricksPolicyInfo#except_principals}.
	ExceptPrincipals *[]*string `field:"optional" json:"exceptPrincipals" yaml:"exceptPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#match_columns DataDatabricksPolicyInfo#match_columns}.
	MatchColumns interface{} `field:"optional" json:"matchColumns" yaml:"matchColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#name DataDatabricksPolicyInfo#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#on_securable_fullname DataDatabricksPolicyInfo#on_securable_fullname}.
	OnSecurableFullname *string `field:"optional" json:"onSecurableFullname" yaml:"onSecurableFullname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#on_securable_type DataDatabricksPolicyInfo#on_securable_type}.
	OnSecurableType *string `field:"optional" json:"onSecurableType" yaml:"onSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#row_filter DataDatabricksPolicyInfo#row_filter}.
	RowFilter *DataDatabricksPolicyInfoRowFilter `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/data-sources/policy_info#when_condition DataDatabricksPolicyInfo#when_condition}.
	WhenCondition *string `field:"optional" json:"whenCondition" yaml:"whenCondition"`
}

