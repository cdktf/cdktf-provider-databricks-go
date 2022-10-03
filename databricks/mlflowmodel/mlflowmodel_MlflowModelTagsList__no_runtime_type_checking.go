//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package mlflowmodel

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MlflowModelTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MlflowModelTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MlflowModelTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMlflowModelTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

