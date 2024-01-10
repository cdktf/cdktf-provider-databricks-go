// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package share

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShareObjectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ShareObjectList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ShareObjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ShareObjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ShareObjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ShareObjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ShareObjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewShareObjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

