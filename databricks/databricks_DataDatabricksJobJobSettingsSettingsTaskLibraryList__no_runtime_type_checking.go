//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskLibraryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDatabricksJobJobSettingsSettingsTaskLibraryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
