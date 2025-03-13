// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineIngestionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/pipeline#connection_name Pipeline#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/pipeline#ingestion_gateway_id Pipeline#ingestion_gateway_id}.
	IngestionGatewayId *string `field:"optional" json:"ingestionGatewayId" yaml:"ingestionGatewayId"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/pipeline#objects Pipeline#objects}
	Objects interface{} `field:"optional" json:"objects" yaml:"objects"`
	// table_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/pipeline#table_configuration Pipeline#table_configuration}
	TableConfiguration *PipelineIngestionDefinitionTableConfiguration `field:"optional" json:"tableConfiguration" yaml:"tableConfiguration"`
}

