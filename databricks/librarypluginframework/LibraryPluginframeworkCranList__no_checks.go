// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package librarypluginframework

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LibraryPluginframeworkCranList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LibraryPluginframeworkCranList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LibraryPluginframeworkCranList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LibraryPluginframeworkCranList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LibraryPluginframeworkCranList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LibraryPluginframeworkCranList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LibraryPluginframeworkCranList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLibraryPluginframeworkCranListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

