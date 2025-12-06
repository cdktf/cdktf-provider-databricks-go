// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registeredmodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegisteredModelConfig struct {
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
	// aliases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#aliases RegisteredModel#aliases}
	Aliases interface{} `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#browse_only RegisteredModel#browse_only}.
	BrowseOnly interface{} `field:"optional" json:"browseOnly" yaml:"browseOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#catalog_name RegisteredModel#catalog_name}.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#comment RegisteredModel#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#created_at RegisteredModel#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#created_by RegisteredModel#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#full_name RegisteredModel#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#id RegisteredModel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#metastore_id RegisteredModel#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#name RegisteredModel#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#owner RegisteredModel#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#schema_name RegisteredModel#schema_name}.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#storage_location RegisteredModel#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#updated_at RegisteredModel#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/registered_model#updated_by RegisteredModel#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

