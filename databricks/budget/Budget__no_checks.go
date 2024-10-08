// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package budget

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Budget) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateImportFromParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (b *jsiiProxy_Budget) validateMoveToIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (b *jsiiProxy_Budget) validatePutAlertConfigurationsParameters(value interface{}) error {
	return nil
}

func (b *jsiiProxy_Budget) validatePutFilterParameters(value *BudgetFilter) error {
	return nil
}

func validateBudget_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateBudget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBudget_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBudget_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetBudgetConfigurationIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetCreateTimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Budget) validateSetUpdateTimeParameters(val *float64) error {
	return nil
}

func validateNewBudgetParameters(scope constructs.Construct, id *string, config *BudgetConfig) error {
	return nil
}

