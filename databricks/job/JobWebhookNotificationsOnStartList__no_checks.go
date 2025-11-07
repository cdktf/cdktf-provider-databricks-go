// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnStartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobWebhookNotificationsOnStartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

