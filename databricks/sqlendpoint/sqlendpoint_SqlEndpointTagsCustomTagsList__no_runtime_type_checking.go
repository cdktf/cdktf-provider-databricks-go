//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package sqlendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlEndpointTagsCustomTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlEndpointTagsCustomTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointTagsCustomTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointTagsCustomTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointTagsCustomTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlEndpointTagsCustomTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlEndpointTagsCustomTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

