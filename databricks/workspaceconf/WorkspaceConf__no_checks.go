// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workspaceconf

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspaceConf) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateImportFromParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateMoveToIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceConf) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateWorkspaceConf_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWorkspaceConf_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkspaceConf_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWorkspaceConf_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetCustomConfigParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_WorkspaceConf) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewWorkspaceConfParameters(scope constructs.Construct, id *string, config *WorkspaceConfConfig) error {
	return nil
}

