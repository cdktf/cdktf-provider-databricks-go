// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pipeline) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateMoveToIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutClusterParameters(value interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutDeploymentParameters(value *PipelineDeployment) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutEnvironmentParameters(value *PipelineEnvironment) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutEventLogParameters(value *PipelineEventLog) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutFiltersParameters(value *PipelineFilters) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutGatewayDefinitionParameters(value *PipelineGatewayDefinition) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutIngestionDefinitionParameters(value *PipelineIngestionDefinition) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutLatestUpdatesParameters(value interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutLibraryParameters(value interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutNotificationParameters(value interface{}) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutRestartWindowParameters(value *PipelineRestartWindow) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutRunAsParameters(value *PipelineRunAs) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutTimeoutsParameters(value *PipelineTimeouts) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validatePutTriggerParameters(value *PipelineTrigger) error {
	return nil
}

func validatePipeline_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validatePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipeline_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validatePipeline_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetAllowDuplicateNamesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetBudgetPolicyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetCatalogParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetCauseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetChannelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetClusterIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetConfigurationParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetContinuousParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetCreatorUserNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetDevelopmentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetEditionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetExpectedLastModifiedParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetHealthParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetLastModifiedParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetPhotonParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetRootPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetRunAsUserNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetServerlessParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetStateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetStorageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetTagsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetTargetParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Pipeline) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewPipelineParameters(scope constructs.Construct, id *string, config *PipelineConfig) error {
	return nil
}

