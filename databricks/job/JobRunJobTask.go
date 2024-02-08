// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobRunJobTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/job#job_id Job#job_id}.
	JobId *float64 `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/job#job_parameters Job#job_parameters}.
	JobParameters *map[string]*string `field:"optional" json:"jobParameters" yaml:"jobParameters"`
}

