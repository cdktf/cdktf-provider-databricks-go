//go:build no_runtime_type_checking

package sqlwidget

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlWidgetParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlWidgetParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlWidgetParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlWidgetParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

