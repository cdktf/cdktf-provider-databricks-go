// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repo

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Repo) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_Repo) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Repo) validatePutSparseCheckoutParameters(value *RepoSparseCheckout) error {
	return nil
}

func validateRepo_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRepo_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRepo_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRepo_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetBranchParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetCommitHashParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetGitProviderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetTagParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repo) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewRepoParameters(scope constructs.Construct, id *string, config *RepoConfig) error {
	return nil
}

