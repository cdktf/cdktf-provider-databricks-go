// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externallocation

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalLocation) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateImportFromParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateMoveToIdParameters(id *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validatePutEncryptionDetailsParameters(value *ExternalLocationEncryptionDetails) error {
	return nil
}

func (e *jsiiProxy_ExternalLocation) validatePutFileEventQueueParameters(value *ExternalLocationFileEventQueue) error {
	return nil
}

func validateExternalLocation_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateExternalLocation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalLocation_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateExternalLocation_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetCredentialNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetEnableFileEventsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetFallbackParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetForceDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetForceUpdateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetIsolationModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetMetastoreIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetOwnerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetReadOnlyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetSkipValidationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalLocation) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewExternalLocationParameters(scope constructs.Construct, id *string, config *ExternalLocationConfig) error {
	return nil
}

