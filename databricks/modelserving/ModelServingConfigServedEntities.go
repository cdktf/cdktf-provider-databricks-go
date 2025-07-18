// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving


type ModelServingConfigServedEntities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#entity_name ModelServing#entity_name}.
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#entity_version ModelServing#entity_version}.
	EntityVersion *string `field:"optional" json:"entityVersion" yaml:"entityVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#environment_vars ModelServing#environment_vars}.
	EnvironmentVars *map[string]*string `field:"optional" json:"environmentVars" yaml:"environmentVars"`
	// external_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#external_model ModelServing#external_model}
	ExternalModel *ModelServingConfigServedEntitiesExternalModel `field:"optional" json:"externalModel" yaml:"externalModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#instance_profile_arn ModelServing#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#max_provisioned_concurrency ModelServing#max_provisioned_concurrency}.
	MaxProvisionedConcurrency *float64 `field:"optional" json:"maxProvisionedConcurrency" yaml:"maxProvisionedConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#max_provisioned_throughput ModelServing#max_provisioned_throughput}.
	MaxProvisionedThroughput *float64 `field:"optional" json:"maxProvisionedThroughput" yaml:"maxProvisionedThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#min_provisioned_concurrency ModelServing#min_provisioned_concurrency}.
	MinProvisionedConcurrency *float64 `field:"optional" json:"minProvisionedConcurrency" yaml:"minProvisionedConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#min_provisioned_throughput ModelServing#min_provisioned_throughput}.
	MinProvisionedThroughput *float64 `field:"optional" json:"minProvisionedThroughput" yaml:"minProvisionedThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#name ModelServing#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#provisioned_model_units ModelServing#provisioned_model_units}.
	ProvisionedModelUnits *float64 `field:"optional" json:"provisionedModelUnits" yaml:"provisionedModelUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#scale_to_zero_enabled ModelServing#scale_to_zero_enabled}.
	ScaleToZeroEnabled interface{} `field:"optional" json:"scaleToZeroEnabled" yaml:"scaleToZeroEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#workload_size ModelServing#workload_size}.
	WorkloadSize *string `field:"optional" json:"workloadSize" yaml:"workloadSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/model_serving#workload_type ModelServing#workload_type}.
	WorkloadType *string `field:"optional" json:"workloadType" yaml:"workloadType"`
}

