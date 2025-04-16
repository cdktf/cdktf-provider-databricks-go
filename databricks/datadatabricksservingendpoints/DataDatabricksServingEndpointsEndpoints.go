// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints


type DataDatabricksServingEndpointsEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#ai_gateway DataDatabricksServingEndpoints#ai_gateway}.
	AiGateway interface{} `field:"optional" json:"aiGateway" yaml:"aiGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#budget_policy_id DataDatabricksServingEndpoints#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#config DataDatabricksServingEndpoints#config}.
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#creation_timestamp DataDatabricksServingEndpoints#creation_timestamp}.
	CreationTimestamp *float64 `field:"optional" json:"creationTimestamp" yaml:"creationTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#creator DataDatabricksServingEndpoints#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#id DataDatabricksServingEndpoints#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#last_updated_timestamp DataDatabricksServingEndpoints#last_updated_timestamp}.
	LastUpdatedTimestamp *float64 `field:"optional" json:"lastUpdatedTimestamp" yaml:"lastUpdatedTimestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#name DataDatabricksServingEndpoints#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#state DataDatabricksServingEndpoints#state}.
	State interface{} `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#tags DataDatabricksServingEndpoints#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/data-sources/serving_endpoints#task DataDatabricksServingEndpoints#task}.
	Task *string `field:"optional" json:"task" yaml:"task"`
}

