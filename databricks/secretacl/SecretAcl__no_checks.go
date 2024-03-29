// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretacl

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretAcl) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretAcl) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSecretAcl_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSecretAcl_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretAcl_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSecretAcl_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetPermissionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetPrincipalParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretAcl) validateSetScopeParameters(val *string) error {
	return nil
}

func validateNewSecretAclParameters(scope constructs.Construct, id *string, config *SecretAclConfig) error {
	return nil
}

