// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabaseinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/database_instance#name DataDatabricksDatabaseInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/database_instance#admin_password DataDatabricksDatabaseInstance#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/database_instance#admin_rolename DataDatabricksDatabaseInstance#admin_rolename}.
	AdminRolename *string `field:"optional" json:"adminRolename" yaml:"adminRolename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/database_instance#capacity DataDatabricksDatabaseInstance#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/data-sources/database_instance#stopped DataDatabricksDatabaseInstance#stopped}.
	Stopped interface{} `field:"optional" json:"stopped" yaml:"stopped"`
}

