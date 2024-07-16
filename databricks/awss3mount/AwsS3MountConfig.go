// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awss3mount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AwsS3MountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/aws_s3_mount#mount_name AwsS3Mount#mount_name}.
	MountName *string `field:"required" json:"mountName" yaml:"mountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/aws_s3_mount#s3_bucket_name AwsS3Mount#s3_bucket_name}.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/aws_s3_mount#cluster_id AwsS3Mount#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/aws_s3_mount#id AwsS3Mount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/aws_s3_mount#instance_profile AwsS3Mount#instance_profile}.
	InstanceProfile *string `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
}

