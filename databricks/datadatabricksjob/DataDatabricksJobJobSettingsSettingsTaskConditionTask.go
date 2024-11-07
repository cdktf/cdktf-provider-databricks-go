// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskConditionTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/job#left DataDatabricksJob#left}.
	Left *string `field:"required" json:"left" yaml:"left"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/job#op DataDatabricksJob#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/data-sources/job#right DataDatabricksJob#right}.
	Right *string `field:"required" json:"right" yaml:"right"`
}

