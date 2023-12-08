// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskSqlTaskQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.1/docs/resources/job#query_id Job#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
}

