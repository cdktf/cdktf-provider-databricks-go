//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobWebhookNotificationsOnSuccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobWebhookNotificationsOnSuccessListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

