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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#name DataDatabricksDatabaseInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#capacity DataDatabricksDatabaseInstance#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#enable_readable_secondaries DataDatabricksDatabaseInstance#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#node_count DataDatabricksDatabaseInstance#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#parent_instance_ref DataDatabricksDatabaseInstance#parent_instance_ref}.
	ParentInstanceRef *DataDatabricksDatabaseInstanceParentInstanceRef `field:"optional" json:"parentInstanceRef" yaml:"parentInstanceRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#retention_window_in_days DataDatabricksDatabaseInstance#retention_window_in_days}.
	RetentionWindowInDays *float64 `field:"optional" json:"retentionWindowInDays" yaml:"retentionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/data-sources/database_instance#stopped DataDatabricksDatabaseInstance#stopped}.
	Stopped interface{} `field:"optional" json:"stopped" yaml:"stopped"`
}

