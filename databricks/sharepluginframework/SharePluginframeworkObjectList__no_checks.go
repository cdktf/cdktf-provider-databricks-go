// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sharepluginframework

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharePluginframeworkObjectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SharePluginframeworkObjectList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SharePluginframeworkObjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SharePluginframeworkObjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SharePluginframeworkObjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SharePluginframeworkObjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SharePluginframeworkObjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSharePluginframeworkObjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

