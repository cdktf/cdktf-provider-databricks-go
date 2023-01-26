//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterClusterMountInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterClusterMountInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterClusterMountInfoList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterClusterMountInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterClusterMountInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterClusterMountInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterClusterMountInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

