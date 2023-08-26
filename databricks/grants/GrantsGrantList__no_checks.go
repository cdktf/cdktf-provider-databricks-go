// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package grants

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GrantsGrantList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GrantsGrantList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GrantsGrantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GrantsGrantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GrantsGrantList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GrantsGrantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGrantsGrantListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

