// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#comment DataDatabricksShare#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#name DataDatabricksShare#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#object DataDatabricksShare#object}.
	Object interface{} `field:"optional" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#owner DataDatabricksShare#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#provider_config DataDatabricksShare#provider_config}.
	ProviderConfig *DataDatabricksShareProviderConfig `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/data-sources/share#storage_root DataDatabricksShare#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
}

