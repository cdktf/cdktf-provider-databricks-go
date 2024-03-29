// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineClusterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PipelineClusterList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineClusterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineClusterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineClusterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

