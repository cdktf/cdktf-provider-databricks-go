// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqlquery

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlQueryParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlQueryParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlQueryParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlQueryParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlQueryParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlQueryParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlQueryParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

