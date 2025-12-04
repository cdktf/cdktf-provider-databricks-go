// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineEventLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#name Pipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#catalog Pipeline#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.98.0/docs/resources/pipeline#schema Pipeline#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

