// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksvolume


type DataDatabricksVolumeVolumeInfoEncryptionDetailsSseEncryptionDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/volume#algorithm DataDatabricksVolume#algorithm}.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.1/docs/data-sources/volume#aws_kms_key_arn DataDatabricksVolume#aws_kms_key_arn}.
	AwsKmsKeyArn *string `field:"optional" json:"awsKmsKeyArn" yaml:"awsKmsKeyArn"`
}

