// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package table


type TableColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#name Table#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#position Table#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_name Table#type_name}.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_text Table#type_text}.
	TypeText *string `field:"required" json:"typeText" yaml:"typeText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#comment Table#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#nullable Table#nullable}.
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#partition_index Table#partition_index}.
	PartitionIndex *float64 `field:"optional" json:"partitionIndex" yaml:"partitionIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_interval_type Table#type_interval_type}.
	TypeIntervalType *string `field:"optional" json:"typeIntervalType" yaml:"typeIntervalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_json Table#type_json}.
	TypeJson *string `field:"optional" json:"typeJson" yaml:"typeJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_precision Table#type_precision}.
	TypePrecision *float64 `field:"optional" json:"typePrecision" yaml:"typePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/table#type_scale Table#type_scale}.
	TypeScale *float64 `field:"optional" json:"typeScale" yaml:"typeScale"`
}

