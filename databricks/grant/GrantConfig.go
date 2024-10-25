// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#principal Grant#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#privileges Grant#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#catalog Grant#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#external_location Grant#external_location}.
	ExternalLocation *string `field:"optional" json:"externalLocation" yaml:"externalLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#foreign_connection Grant#foreign_connection}.
	ForeignConnection *string `field:"optional" json:"foreignConnection" yaml:"foreignConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#function Grant#function}.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#id Grant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#metastore Grant#metastore}.
	Metastore *string `field:"optional" json:"metastore" yaml:"metastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#model Grant#model}.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#pipeline Grant#pipeline}.
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#recipient Grant#recipient}.
	Recipient *string `field:"optional" json:"recipient" yaml:"recipient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#schema Grant#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#share Grant#share}.
	Share *string `field:"optional" json:"share" yaml:"share"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#storage_credential Grant#storage_credential}.
	StorageCredential *string `field:"optional" json:"storageCredential" yaml:"storageCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#table Grant#table}.
	Table *string `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/grant#volume Grant#volume}.
	Volume *string `field:"optional" json:"volume" yaml:"volume"`
}

