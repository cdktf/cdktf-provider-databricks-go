// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package policyinfo

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyInfoColumnMaskUsingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PolicyInfoColumnMaskUsingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PolicyInfoColumnMaskUsingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPolicyInfoColumnMaskUsingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

