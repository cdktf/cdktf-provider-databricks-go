// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskSqlTaskDashboardSubscriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#destination_id Job#destination_id}.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/job#user_name Job#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

