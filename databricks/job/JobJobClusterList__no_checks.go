// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobJobClusterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobJobClusterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobJobClusterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

