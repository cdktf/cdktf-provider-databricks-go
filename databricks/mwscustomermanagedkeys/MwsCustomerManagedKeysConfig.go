// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwscustomermanagedkeys

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsCustomerManagedKeysConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#account_id MwsCustomerManagedKeys#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#use_cases MwsCustomerManagedKeys#use_cases}.
	UseCases *[]*string `field:"required" json:"useCases" yaml:"useCases"`
	// aws_key_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#aws_key_info MwsCustomerManagedKeys#aws_key_info}
	AwsKeyInfo *MwsCustomerManagedKeysAwsKeyInfo `field:"optional" json:"awsKeyInfo" yaml:"awsKeyInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#creation_time MwsCustomerManagedKeys#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#customer_managed_key_id MwsCustomerManagedKeys#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
	// gcp_key_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#gcp_key_info MwsCustomerManagedKeys#gcp_key_info}
	GcpKeyInfo *MwsCustomerManagedKeysGcpKeyInfo `field:"optional" json:"gcpKeyInfo" yaml:"gcpKeyInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/mws_customer_managed_keys#id MwsCustomerManagedKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

