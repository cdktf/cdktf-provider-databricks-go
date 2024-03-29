// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package onlinetable

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusContinuousUpdateStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOnlineTableStatusContinuousUpdateStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

