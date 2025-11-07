// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PipelineIngestionDefinitionObjectsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPipelineIngestionDefinitionObjectsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

