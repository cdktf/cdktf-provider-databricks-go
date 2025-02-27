// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetAlertConfigurations struct {
	// action_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#action_configurations Budget#action_configurations}
	ActionConfigurations interface{} `field:"optional" json:"actionConfigurations" yaml:"actionConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#alert_configuration_id Budget#alert_configuration_id}.
	AlertConfigurationId *string `field:"optional" json:"alertConfigurationId" yaml:"alertConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#quantity_threshold Budget#quantity_threshold}.
	QuantityThreshold *string `field:"optional" json:"quantityThreshold" yaml:"quantityThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#quantity_type Budget#quantity_type}.
	QuantityType *string `field:"optional" json:"quantityType" yaml:"quantityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#time_period Budget#time_period}.
	TimePeriod *string `field:"optional" json:"timePeriod" yaml:"timePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.67.0/docs/resources/budget#trigger_type Budget#trigger_type}.
	TriggerType *string `field:"optional" json:"triggerType" yaml:"triggerType"`
}

