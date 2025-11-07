// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package clusterpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterPolicyLibrariesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClusterPolicyLibrariesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterPolicyLibrariesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterPolicyLibrariesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterPolicyLibrariesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterPolicyLibrariesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterPolicyLibrariesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterPolicyLibrariesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

