// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mount


type MountGs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/mount#bucket_name Mount#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/mount#service_account Mount#service_account}.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

