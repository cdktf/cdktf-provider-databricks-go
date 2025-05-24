// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputConfigServedEntities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/model_serving_provisioned_throughput#entity_name ModelServingProvisionedThroughput#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/model_serving_provisioned_throughput#entity_version ModelServingProvisionedThroughput#entity_version}.
	EntityVersion *string `field:"required" json:"entityVersion" yaml:"entityVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/model_serving_provisioned_throughput#provisioned_model_units ModelServingProvisionedThroughput#provisioned_model_units}.
	ProvisionedModelUnits *float64 `field:"required" json:"provisionedModelUnits" yaml:"provisionedModelUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/model_serving_provisioned_throughput#name ModelServingProvisionedThroughput#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

