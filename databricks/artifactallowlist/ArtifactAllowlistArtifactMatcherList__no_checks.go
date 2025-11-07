// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package artifactallowlist

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ArtifactAllowlistArtifactMatcherList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewArtifactAllowlistArtifactMatcherListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

