// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabricksProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DatabricksProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDatabricksProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDatabricksProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabricksProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDatabricksProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabricksProvider) validateSetAzureUseMsiParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabricksProvider) validateSetDebugHeadersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabricksProvider) validateSetSkipVerifyParameters(val interface{}) error {
	return nil
}

func validateNewDatabricksProviderParameters(scope constructs.Construct, id *string, config *DatabricksProviderConfig) error {
	return nil
}

