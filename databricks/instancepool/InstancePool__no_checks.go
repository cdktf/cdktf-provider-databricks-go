// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package instancepool

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstancePool) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutAwsAttributesParameters(value *InstancePoolAwsAttributes) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutAzureAttributesParameters(value *InstancePoolAzureAttributes) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutDiskSpecParameters(value *InstancePoolDiskSpec) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutGcpAttributesParameters(value *InstancePoolGcpAttributes) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutInstancePoolFleetAttributesParameters(value *InstancePoolInstancePoolFleetAttributes) error {
	return nil
}

func (i *jsiiProxy_InstancePool) validatePutPreloadedDockerImageParameters(value interface{}) error {
	return nil
}

func validateInstancePool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInstancePool_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateInstancePool_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetCustomTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetEnableElasticDiskParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetIdleInstanceAutoterminationMinutesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetInstancePoolIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetInstancePoolNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetMaxCapacityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetMinIdleInstancesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetNodeTypeIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetPreloadedSparkVersionsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_InstancePool) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewInstancePoolParameters(scope constructs.Construct, id *string, config *InstancePoolConfig) error {
	return nil
}

