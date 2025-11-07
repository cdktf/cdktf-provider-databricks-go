// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connection

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectionProvisioningInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConnectionProvisioningInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConnectionProvisioningInfoList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConnectionProvisioningInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConnectionProvisioningInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConnectionProvisioningInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConnectionProvisioningInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

