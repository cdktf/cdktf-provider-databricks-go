// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alert

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Alert) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_Alert) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_Alert) validatePutConditionParameters(value *AlertCondition) error {
	return nil
}

func validateAlert_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAlert_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlert_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAlert_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetCustomBodyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetCustomSubjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetNotifyOnOkParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetOwnerUserNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetParentPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetQueryIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Alert) validateSetSecondsToRetriggerParameters(val *float64) error {
	return nil
}

func validateNewAlertParameters(scope constructs.Construct, id *string, config *AlertConfig) error {
	return nil
}

