// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package library

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LibraryPypiList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LibraryPypiList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LibraryPypiList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LibraryPypiList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LibraryPypiList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LibraryPypiList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LibraryPypiList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLibraryPypiListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

