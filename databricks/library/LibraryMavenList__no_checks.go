// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package library

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LibraryMavenList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LibraryMavenList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LibraryMavenList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LibraryMavenList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LibraryMavenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LibraryMavenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LibraryMavenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLibraryMavenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

