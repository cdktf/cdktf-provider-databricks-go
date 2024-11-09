// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickssharepluginframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksSharePluginframeworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#comment DataDatabricksSharePluginframework#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#created_at DataDatabricksSharePluginframework#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#created_by DataDatabricksSharePluginframework#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#effective_owner DataDatabricksSharePluginframework#effective_owner}.
	EffectiveOwner *string `field:"optional" json:"effectiveOwner" yaml:"effectiveOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#name DataDatabricksSharePluginframework#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#object DataDatabricksSharePluginframework#object}
	Object interface{} `field:"optional" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#owner DataDatabricksSharePluginframework#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#storage_location DataDatabricksSharePluginframework#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#storage_root DataDatabricksSharePluginframework#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#updated_at DataDatabricksSharePluginframework#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/share_pluginframework#updated_by DataDatabricksSharePluginframework#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

