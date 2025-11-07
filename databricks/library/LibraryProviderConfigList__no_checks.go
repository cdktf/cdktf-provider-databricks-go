// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package library

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LibraryProviderConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LibraryProviderConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LibraryProviderConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LibraryProviderConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LibraryProviderConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LibraryProviderConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LibraryProviderConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLibraryProviderConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

