// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package modelserving

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ModelServingRateLimitsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ModelServingRateLimitsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ModelServingRateLimitsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ModelServingRateLimitsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewModelServingRateLimitsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

