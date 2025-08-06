// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskDbtCloudTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#connection_resource_name Job#connection_resource_name}.
	ConnectionResourceName *string `field:"optional" json:"connectionResourceName" yaml:"connectionResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#dbt_cloud_job_id Job#dbt_cloud_job_id}.
	DbtCloudJobId *float64 `field:"optional" json:"dbtCloudJobId" yaml:"dbtCloudJobId"`
}

