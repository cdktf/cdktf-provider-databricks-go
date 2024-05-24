// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwscustomermanagedkeys


type MwsCustomerManagedKeysGcpKeyInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.45.0/docs/resources/mws_customer_managed_keys#kms_key_id MwsCustomerManagedKeys#kms_key_id}.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

