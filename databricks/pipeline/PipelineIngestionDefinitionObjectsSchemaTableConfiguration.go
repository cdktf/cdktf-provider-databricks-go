// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsSchemaTableConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/pipeline#primary_keys Pipeline#primary_keys}.
	PrimaryKeys *[]*string `field:"optional" json:"primaryKeys" yaml:"primaryKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/pipeline#salesforce_include_formula_fields Pipeline#salesforce_include_formula_fields}.
	SalesforceIncludeFormulaFields interface{} `field:"optional" json:"salesforceIncludeFormulaFields" yaml:"salesforceIncludeFormulaFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/pipeline#scd_type Pipeline#scd_type}.
	ScdType *string `field:"optional" json:"scdType" yaml:"scdType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/pipeline#sequence_by Pipeline#sequence_by}.
	SequenceBy *[]*string `field:"optional" json:"sequenceBy" yaml:"sequenceBy"`
}

