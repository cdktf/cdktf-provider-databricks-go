// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwscustomermanagedkeys


type MwsCustomerManagedKeysAwsKeyInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/mws_customer_managed_keys#key_arn MwsCustomerManagedKeys#key_arn}.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/mws_customer_managed_keys#key_alias MwsCustomerManagedKeys#key_alias}.
	KeyAlias *string `field:"optional" json:"keyAlias" yaml:"keyAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.82.0/docs/resources/mws_customer_managed_keys#key_region MwsCustomerManagedKeys#key_region}.
	KeyRegion *string `field:"optional" json:"keyRegion" yaml:"keyRegion"`
}

