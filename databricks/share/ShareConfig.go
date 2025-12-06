// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package share

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#name Share#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#comment Share#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#object Share#object}
	Object interface{} `field:"optional" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#owner Share#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#provider_config Share#provider_config}
	ProviderConfig interface{} `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/share#storage_root Share#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
}

