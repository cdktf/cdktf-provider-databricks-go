// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#name DatabaseInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#capacity DatabaseInstance#capacity}.
	Capacity *string `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#custom_tags DatabaseInstance#custom_tags}.
	CustomTags interface{} `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#enable_pg_native_login DatabaseInstance#enable_pg_native_login}.
	EnablePgNativeLogin interface{} `field:"optional" json:"enablePgNativeLogin" yaml:"enablePgNativeLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#enable_readable_secondaries DatabaseInstance#enable_readable_secondaries}.
	EnableReadableSecondaries interface{} `field:"optional" json:"enableReadableSecondaries" yaml:"enableReadableSecondaries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#node_count DatabaseInstance#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#parent_instance_ref DatabaseInstance#parent_instance_ref}.
	ParentInstanceRef *DatabaseInstanceParentInstanceRef `field:"optional" json:"parentInstanceRef" yaml:"parentInstanceRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#purge_on_delete DatabaseInstance#purge_on_delete}.
	PurgeOnDelete interface{} `field:"optional" json:"purgeOnDelete" yaml:"purgeOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#retention_window_in_days DatabaseInstance#retention_window_in_days}.
	RetentionWindowInDays *float64 `field:"optional" json:"retentionWindowInDays" yaml:"retentionWindowInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#stopped DatabaseInstance#stopped}.
	Stopped interface{} `field:"optional" json:"stopped" yaml:"stopped"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/database_instance#usage_policy_id DatabaseInstance#usage_policy_id}.
	UsagePolicyId *string `field:"optional" json:"usagePolicyId" yaml:"usagePolicyId"`
}

