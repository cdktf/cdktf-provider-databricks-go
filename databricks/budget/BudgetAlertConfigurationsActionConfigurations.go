// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budget


type BudgetAlertConfigurationsActionConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/budget#action_configuration_id Budget#action_configuration_id}.
	ActionConfigurationId *string `field:"optional" json:"actionConfigurationId" yaml:"actionConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/budget#action_type Budget#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/budget#target Budget#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

