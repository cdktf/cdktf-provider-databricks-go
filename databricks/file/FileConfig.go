// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package file

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#path File#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#content_base64 File#content_base64}.
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#id File#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#md5 File#md5}.
	Md5 *string `field:"optional" json:"md5" yaml:"md5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#remote_file_modified File#remote_file_modified}.
	RemoteFileModified interface{} `field:"optional" json:"remoteFileModified" yaml:"remoteFileModified"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/file#source File#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}
