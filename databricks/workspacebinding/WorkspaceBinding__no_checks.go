// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspacebinding

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspaceBinding) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateImportFromParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateMoveToIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBinding) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateWorkspaceBinding_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWorkspaceBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkspaceBinding_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWorkspaceBinding_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetBindingTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetCatalogNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetSecurableNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetSecurableTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceBinding) validateSetWorkspaceIdParameters(val *float64) error {
	return nil
}

func validateNewWorkspaceBindingParameters(scope constructs.Construct, id *string, config *WorkspaceBindingConfig) error {
	return nil
}

