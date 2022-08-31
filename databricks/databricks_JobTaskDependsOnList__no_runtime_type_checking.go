//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobTaskDependsOnList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobTaskDependsOnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobTaskDependsOnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

