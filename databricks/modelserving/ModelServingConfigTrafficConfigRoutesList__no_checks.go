// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package modelserving

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ModelServingConfigTrafficConfigRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewModelServingConfigTrafficConfigRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

