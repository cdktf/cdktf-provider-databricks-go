// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretscope

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretScope) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SecretScope) validatePutKeyvaultMetadataParameters(value *SecretScopeKeyvaultMetadata) error {
	return nil
}

func validateSecretScope_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSecretScope_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretScope_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSecretScope_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetBackendTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetInitialManagePrincipalParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretScope) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewSecretScopeParameters(scope constructs.Construct, id *string, config *SecretScopeConfig) error {
	return nil
}

