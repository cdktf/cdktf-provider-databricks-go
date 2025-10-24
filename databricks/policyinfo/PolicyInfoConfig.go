// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policyinfo

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyInfoConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#for_securable_type PolicyInfo#for_securable_type}.
	ForSecurableType *string `field:"required" json:"forSecurableType" yaml:"forSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#policy_type PolicyInfo#policy_type}.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#to_principals PolicyInfo#to_principals}.
	ToPrincipals *[]*string `field:"required" json:"toPrincipals" yaml:"toPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#column_mask PolicyInfo#column_mask}.
	ColumnMask *PolicyInfoColumnMask `field:"optional" json:"columnMask" yaml:"columnMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#comment PolicyInfo#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#except_principals PolicyInfo#except_principals}.
	ExceptPrincipals *[]*string `field:"optional" json:"exceptPrincipals" yaml:"exceptPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#match_columns PolicyInfo#match_columns}.
	MatchColumns interface{} `field:"optional" json:"matchColumns" yaml:"matchColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#name PolicyInfo#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#on_securable_fullname PolicyInfo#on_securable_fullname}.
	OnSecurableFullname *string `field:"optional" json:"onSecurableFullname" yaml:"onSecurableFullname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#on_securable_type PolicyInfo#on_securable_type}.
	OnSecurableType *string `field:"optional" json:"onSecurableType" yaml:"onSecurableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#row_filter PolicyInfo#row_filter}.
	RowFilter *PolicyInfoRowFilter `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/policy_info#when_condition PolicyInfo#when_condition}.
	WhenCondition *string `field:"optional" json:"whenCondition" yaml:"whenCondition"`
}

