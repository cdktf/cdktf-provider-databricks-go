// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions


type DataDatabricksFunctionsFunctionsInputParamsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#name DataDatabricksFunctions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#position DataDatabricksFunctions#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_name DataDatabricksFunctions#type_name}.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_text DataDatabricksFunctions#type_text}.
	TypeText *string `field:"required" json:"typeText" yaml:"typeText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#comment DataDatabricksFunctions#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#parameter_default DataDatabricksFunctions#parameter_default}.
	ParameterDefault *string `field:"optional" json:"parameterDefault" yaml:"parameterDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#parameter_mode DataDatabricksFunctions#parameter_mode}.
	ParameterMode *string `field:"optional" json:"parameterMode" yaml:"parameterMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#parameter_type DataDatabricksFunctions#parameter_type}.
	ParameterType *string `field:"optional" json:"parameterType" yaml:"parameterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_interval_type DataDatabricksFunctions#type_interval_type}.
	TypeIntervalType *string `field:"optional" json:"typeIntervalType" yaml:"typeIntervalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_json DataDatabricksFunctions#type_json}.
	TypeJson *string `field:"optional" json:"typeJson" yaml:"typeJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_precision DataDatabricksFunctions#type_precision}.
	TypePrecision *float64 `field:"optional" json:"typePrecision" yaml:"typePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.0/docs/data-sources/functions#type_scale DataDatabricksFunctions#type_scale}.
	TypeScale *float64 `field:"optional" json:"typeScale" yaml:"typeScale"`
}

