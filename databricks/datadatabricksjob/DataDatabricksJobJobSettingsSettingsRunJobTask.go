// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsRunJobTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/job#job_id DataDatabricksJob#job_id}.
	JobId *float64 `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/job#job_parameters DataDatabricksJob#job_parameters}.
	JobParameters *map[string]*string `field:"optional" json:"jobParameters" yaml:"jobParameters"`
}

