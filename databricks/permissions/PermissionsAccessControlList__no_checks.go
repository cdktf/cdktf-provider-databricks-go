// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package permissions

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PermissionsAccessControlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PermissionsAccessControlList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PermissionsAccessControlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PermissionsAccessControlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PermissionsAccessControlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PermissionsAccessControlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PermissionsAccessControlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPermissionsAccessControlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

