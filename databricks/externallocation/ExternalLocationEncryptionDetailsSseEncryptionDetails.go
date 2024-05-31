// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externallocation


type ExternalLocationEncryptionDetailsSseEncryptionDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/external_location#algorithm ExternalLocation#algorithm}.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.46.0/docs/resources/external_location#aws_kms_key_arn ExternalLocation#aws_kms_key_arn}.
	AwsKmsKeyArn *string `field:"optional" json:"awsKmsKeyArn" yaml:"awsKmsKeyArn"`
}

