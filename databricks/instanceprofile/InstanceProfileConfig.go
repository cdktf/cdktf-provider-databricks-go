// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instanceprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstanceProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/instance_profile#instance_profile_arn InstanceProfile#instance_profile_arn}.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/instance_profile#iam_role_arn InstanceProfile#iam_role_arn}.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/instance_profile#id InstanceProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/instance_profile#is_meta_instance_profile InstanceProfile#is_meta_instance_profile}.
	IsMetaInstanceProfile interface{} `field:"optional" json:"isMetaInstanceProfile" yaml:"isMetaInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.2/docs/resources/instance_profile#skip_validation InstanceProfile#skip_validation}.
	SkipValidation interface{} `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

