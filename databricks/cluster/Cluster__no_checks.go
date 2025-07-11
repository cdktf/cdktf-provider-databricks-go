// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutAutoscaleParameters(value *ClusterAutoscale) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutAwsAttributesParameters(value *ClusterAwsAttributes) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutAzureAttributesParameters(value *ClusterAzureAttributes) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutClusterLogConfParameters(value *ClusterClusterLogConf) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutClusterMountInfoParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutDockerImageParameters(value *ClusterDockerImage) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutGcpAttributesParameters(value *ClusterGcpAttributes) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutInitScriptsParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutLibraryParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutTimeoutsParameters(value *ClusterTimeouts) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutWorkloadTypeParameters(value *ClusterWorkloadType) error {
	return nil
}

func validateCluster_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetApplyPolicyDefaultValuesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAutoterminationMinutesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetClusterNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetCustomTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetDataSecurityModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetDriverInstancePoolIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetDriverNodeTypeIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetEnableElasticDiskParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetEnableLocalDiskEncryptionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIdempotencyTokenParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetInstancePoolIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIsPinnedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIsSingleNodeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetKindParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetNodeTypeIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetNoWaitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetNumWorkersParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetPolicyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetRemoteDiskThroughputParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetRuntimeEngineParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetSingleUserNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetSparkConfParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetSparkEnvVarsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetSparkVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetSshPublicKeysParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetTotalInitialRemoteDiskSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetUseMlRuntimeParameters(val interface{}) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, config *ClusterConfig) error {
	return nil
}

