// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineGatewayDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline#connection_id Pipeline#connection_id}.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline#connection_name Pipeline#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline#gateway_storage_catalog Pipeline#gateway_storage_catalog}.
	GatewayStorageCatalog *string `field:"optional" json:"gatewayStorageCatalog" yaml:"gatewayStorageCatalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline#gateway_storage_name Pipeline#gateway_storage_name}.
	GatewayStorageName *string `field:"optional" json:"gatewayStorageName" yaml:"gatewayStorageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/resources/pipeline#gateway_storage_schema Pipeline#gateway_storage_schema}.
	GatewayStorageSchema *string `field:"optional" json:"gatewayStorageSchema" yaml:"gatewayStorageSchema"`
}

