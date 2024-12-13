// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#comment DataDatabricksTable#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// mask block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#mask DataDatabricksTable#mask}
	Mask *DataDatabricksTableTableInfoColumnsMask `field:"optional" json:"mask" yaml:"mask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#name DataDatabricksTable#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#nullable DataDatabricksTable#nullable}.
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#partition_index DataDatabricksTable#partition_index}.
	PartitionIndex *float64 `field:"optional" json:"partitionIndex" yaml:"partitionIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#position DataDatabricksTable#position}.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_interval_type DataDatabricksTable#type_interval_type}.
	TypeIntervalType *string `field:"optional" json:"typeIntervalType" yaml:"typeIntervalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_json DataDatabricksTable#type_json}.
	TypeJson *string `field:"optional" json:"typeJson" yaml:"typeJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_name DataDatabricksTable#type_name}.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_precision DataDatabricksTable#type_precision}.
	TypePrecision *float64 `field:"optional" json:"typePrecision" yaml:"typePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_scale DataDatabricksTable#type_scale}.
	TypeScale *float64 `field:"optional" json:"typeScale" yaml:"typeScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/table#type_text DataDatabricksTable#type_text}.
	TypeText *string `field:"optional" json:"typeText" yaml:"typeText"`
}

