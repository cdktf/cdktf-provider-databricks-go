// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package query

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_Query) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateImportFromParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (q *jsiiProxy_Query) validateMoveToIdParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_Query) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (q *jsiiProxy_Query) validatePutParameterParameters(value interface{}) error {
	return nil
}

func validateQuery_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateQuery_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQuery_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateQuery_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetApplyAutoLimitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetCatalogParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetOwnerUserNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetParentPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetQueryTextParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetRunAsModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Query) validateSetWarehouseIdParameters(val *string) error {
	return nil
}

func validateNewQueryParameters(scope constructs.Construct, id *string, config *QueryConfig) error {
	return nil
}

