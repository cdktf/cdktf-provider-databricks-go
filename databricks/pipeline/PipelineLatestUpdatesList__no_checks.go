// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineLatestUpdatesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PipelineLatestUpdatesList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineLatestUpdatesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineLatestUpdatesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineLatestUpdatesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineLatestUpdatesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineLatestUpdatesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineLatestUpdatesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

