//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineClusterInitScriptsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineClusterInitScriptsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterInitScriptsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterInitScriptsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterInitScriptsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterInitScriptsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineClusterInitScriptsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

