// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqlendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlEndpointHealthList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlEndpointHealthList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlEndpointHealthList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointHealthList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointHealthList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointHealthList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlEndpointHealthListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

