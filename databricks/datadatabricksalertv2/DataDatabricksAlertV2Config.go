// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAlertV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#custom_description DataDatabricksAlertV2#custom_description}.
	CustomDescription *string `field:"optional" json:"customDescription" yaml:"customDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#custom_summary DataDatabricksAlertV2#custom_summary}.
	CustomSummary *string `field:"optional" json:"customSummary" yaml:"customSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#display_name DataDatabricksAlertV2#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#evaluation DataDatabricksAlertV2#evaluation}.
	Evaluation *DataDatabricksAlertV2Evaluation `field:"optional" json:"evaluation" yaml:"evaluation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#parent_path DataDatabricksAlertV2#parent_path}.
	ParentPath *string `field:"optional" json:"parentPath" yaml:"parentPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#query_text DataDatabricksAlertV2#query_text}.
	QueryText *string `field:"optional" json:"queryText" yaml:"queryText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#schedule DataDatabricksAlertV2#schedule}.
	Schedule *DataDatabricksAlertV2Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/alert_v2#warehouse_id DataDatabricksAlertV2#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

