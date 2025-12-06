// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineGatewayDefinitionConnectionParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/pipeline#source_catalog Pipeline#source_catalog}.
	SourceCatalog *string `field:"optional" json:"sourceCatalog" yaml:"sourceCatalog"`
}

