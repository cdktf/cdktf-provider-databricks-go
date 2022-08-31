//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDatabricksClusterClusterInfoExecutorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

