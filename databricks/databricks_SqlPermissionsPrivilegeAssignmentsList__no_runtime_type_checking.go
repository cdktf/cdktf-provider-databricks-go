//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlPermissionsPrivilegeAssignmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlPermissionsPrivilegeAssignmentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
