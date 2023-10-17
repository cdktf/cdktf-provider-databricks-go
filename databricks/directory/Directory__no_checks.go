// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package directory

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Directory) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Directory) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Directory) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDirectory_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDirectory_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDirectory_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDirectory_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetDeleteRecursiveParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetObjectIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Directory) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewDirectoryParameters(scope constructs.Construct, id *string, config *DirectoryConfig) error {
	return nil
}

