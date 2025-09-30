// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinitionObjectsReportTableConfigurationWorkdayReportParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/pipeline#incremental Pipeline#incremental}.
	Incremental interface{} `field:"optional" json:"incremental" yaml:"incremental"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/pipeline#parameters Pipeline#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// report_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/pipeline#report_parameters Pipeline#report_parameters}
	ReportParameters interface{} `field:"optional" json:"reportParameters" yaml:"reportParameters"`
}

