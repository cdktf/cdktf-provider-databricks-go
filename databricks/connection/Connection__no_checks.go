// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package connection

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Connection) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Connection) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Connection) validatePutEnvironmentSettingsParameters(value *ConnectionEnvironmentSettings) error {
	return nil
}

func validateConnection_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConnection_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateConnection_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetConnectionTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetOptionsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetOwnerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetPropertiesParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Connection) validateSetReadOnlyParameters(val interface{}) error {
	return nil
}

func validateNewConnectionParameters(scope constructs.Construct, id *string, config *ConnectionConfig) error {
	return nil
}

