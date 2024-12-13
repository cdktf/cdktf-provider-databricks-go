// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewClusterInitScriptsS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#canned_acl Job#canned_acl}.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#enable_encryption Job#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#encryption_type Job#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#endpoint Job#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#kms_key Job#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/job#region Job#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

