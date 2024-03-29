// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package onlinetable

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusProvisioningStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOnlineTableStatusProvisioningStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

