// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package app

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_App) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_App) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_App) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_App) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_App) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_App) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_App) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_App) validatePutResourcesParameters(value interface{}) error {
	return nil
}

func validateApp_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApp_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateApp_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_App) validateSetBudgetPolicyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_App) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_App) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_App) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_App) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_App) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_App) validateSetNoComputeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_App) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_App) validateSetUserApiScopesParameters(val *[]*string) error {
	return nil
}

func validateNewAppParameters(scope constructs.Construct, id *string, config *AppConfig) error {
	return nil
}

