// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoClusterLogConfS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#destination DataDatabricksCluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#canned_acl DataDatabricksCluster#canned_acl}.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#enable_encryption DataDatabricksCluster#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#encryption_type DataDatabricksCluster#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#endpoint DataDatabricksCluster#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#kms_key DataDatabricksCluster#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/data-sources/cluster#region DataDatabricksCluster#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

