// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package vectorsearchindex

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorSearchIndexStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VectorSearchIndexStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VectorSearchIndexStatusList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VectorSearchIndexStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VectorSearchIndexStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VectorSearchIndexStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVectorSearchIndexStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

