// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package token

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Token) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_Token) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateToken_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateToken_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateToken_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetCreationTimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetExpiryTimeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetLifetimeSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetTokenIdParameters(val *string) error {
	return nil
}

func validateNewTokenParameters(scope constructs.Construct, id *string, config *TokenConfig) error {
	return nil
}

