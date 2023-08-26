// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package library

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Library) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Library) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_Library) validatePutCranParameters(value *LibraryCran) error {
	return nil
}

func (l *jsiiProxy_Library) validatePutMavenParameters(value *LibraryMaven) error {
	return nil
}

func (l *jsiiProxy_Library) validatePutPypiParameters(value *LibraryPypi) error {
	return nil
}

func validateLibrary_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLibrary_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLibrary_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetClusterIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetEggParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetJarParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Library) validateSetWhlParameters(val *string) error {
	return nil
}

func validateNewLibraryParameters(scope constructs.Construct, id *string, config *LibraryConfig) error {
	return nil
}

