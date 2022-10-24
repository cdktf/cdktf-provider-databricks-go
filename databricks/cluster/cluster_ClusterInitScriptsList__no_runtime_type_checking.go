//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterInitScriptsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterInitScriptsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterInitScriptsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterInitScriptsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterInitScriptsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterInitScriptsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterInitScriptsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

