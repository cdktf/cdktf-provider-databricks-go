// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package onlinetable

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OnlineTableStatusTriggeredUpdateStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOnlineTableStatusTriggeredUpdateStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

