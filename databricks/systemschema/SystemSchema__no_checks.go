// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package systemschema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SystemSchema) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SystemSchema) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSystemSchema_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSystemSchema_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSystemSchema_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSystemSchema_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SystemSchema) validateSetSchemaParameters(val *string) error {
	return nil
}

func validateNewSystemSchemaParameters(scope constructs.Construct, id *string, config *SystemSchemaConfig) error {
	return nil
}

