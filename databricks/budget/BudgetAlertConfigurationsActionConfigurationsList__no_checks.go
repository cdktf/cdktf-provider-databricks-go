// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package budget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BudgetAlertConfigurationsActionConfigurationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBudgetAlertConfigurationsActionConfigurationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

