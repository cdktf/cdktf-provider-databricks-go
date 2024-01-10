// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package instancepool

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstancePoolPreloadedDockerImageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_InstancePoolPreloadedDockerImageList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InstancePoolPreloadedDockerImageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InstancePoolPreloadedDockerImageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstancePoolPreloadedDockerImageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstancePoolPreloadedDockerImageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InstancePoolPreloadedDockerImageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInstancePoolPreloadedDockerImageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

