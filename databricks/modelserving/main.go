// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServing",
		reflect.TypeOf((*ModelServing)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConfig", GoMethod: "PutConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putRateLimits", GoMethod: "PutRateLimits"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimits", GoGetter: "RateLimits"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimitsInput", GoGetter: "RateLimitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRateLimits", GoMethod: "ResetRateLimits"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouteOptimized", GoMethod: "ResetRouteOptimized"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "routeOptimized", GoGetter: "RouteOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "routeOptimizedInput", GoGetter: "RouteOptimizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "servingEndpointId", GoGetter: "ServingEndpointId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServing{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfig",
		reflect.TypeOf((*ModelServingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigA",
		reflect.TypeOf((*ModelServingConfigA)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigAOutputReference",
		reflect.TypeOf((*ModelServingConfigAOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoCaptureConfig", GoGetter: "AutoCaptureConfig"},
			_jsii_.MemberProperty{JsiiProperty: "autoCaptureConfigInput", GoGetter: "AutoCaptureConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoCaptureConfig", GoMethod: "PutAutoCaptureConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putServedEntities", GoMethod: "PutServedEntities"},
			_jsii_.MemberMethod{JsiiMethod: "putServedModels", GoMethod: "PutServedModels"},
			_jsii_.MemberMethod{JsiiMethod: "putTrafficConfig", GoMethod: "PutTrafficConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoCaptureConfig", GoMethod: "ResetAutoCaptureConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetServedEntities", GoMethod: "ResetServedEntities"},
			_jsii_.MemberMethod{JsiiMethod: "resetServedModels", GoMethod: "ResetServedModels"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficConfig", GoMethod: "ResetTrafficConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "servedEntities", GoGetter: "ServedEntities"},
			_jsii_.MemberProperty{JsiiProperty: "servedEntitiesInput", GoGetter: "ServedEntitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "servedModels", GoGetter: "ServedModels"},
			_jsii_.MemberProperty{JsiiProperty: "servedModelsInput", GoGetter: "ServedModelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficConfig", GoGetter: "TrafficConfig"},
			_jsii_.MemberProperty{JsiiProperty: "trafficConfigInput", GoGetter: "TrafficConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigAOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigAutoCaptureConfig",
		reflect.TypeOf((*ModelServingConfigAutoCaptureConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigAutoCaptureConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigAutoCaptureConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "catalogName", GoGetter: "CatalogName"},
			_jsii_.MemberProperty{JsiiProperty: "catalogNameInput", GoGetter: "CatalogNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCatalogName", GoMethod: "ResetCatalogName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaName", GoMethod: "ResetSchemaName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTableNamePrefix", GoMethod: "ResetTableNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemaName", GoGetter: "SchemaName"},
			_jsii_.MemberProperty{JsiiProperty: "schemaNameInput", GoGetter: "SchemaNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "tableNamePrefix", GoGetter: "TableNamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "tableNamePrefixInput", GoGetter: "TableNamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigAutoCaptureConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntities",
		reflect.TypeOf((*ModelServingConfigServedEntities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModel",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAi21LabsConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAi21LabsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsApiKey", GoGetter: "Ai21LabsApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsApiKeyInput", GoGetter: "Ai21LabsApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsApiKeyPlaintext", GoGetter: "Ai21LabsApiKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsApiKeyPlaintextInput", GoGetter: "Ai21LabsApiKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAi21LabsApiKey", GoMethod: "ResetAi21LabsApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAi21LabsApiKeyPlaintext", GoMethod: "ResetAi21LabsApiKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyId", GoGetter: "AwsAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdInput", GoGetter: "AwsAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdPlaintext", GoGetter: "AwsAccessKeyIdPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdPlaintextInput", GoGetter: "AwsAccessKeyIdPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRegion", GoGetter: "AwsRegion"},
			_jsii_.MemberProperty{JsiiProperty: "awsRegionInput", GoGetter: "AwsRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKey", GoGetter: "AwsSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyInput", GoGetter: "AwsSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyPlaintext", GoGetter: "AwsSecretAccessKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyPlaintextInput", GoGetter: "AwsSecretAccessKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "bedrockProvider", GoGetter: "BedrockProvider"},
			_jsii_.MemberProperty{JsiiProperty: "bedrockProviderInput", GoGetter: "BedrockProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessKeyId", GoMethod: "ResetAwsAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessKeyIdPlaintext", GoMethod: "ResetAwsAccessKeyIdPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretAccessKey", GoMethod: "ResetAwsSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretAccessKeyPlaintext", GoMethod: "ResetAwsSecretAccessKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAnthropicConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAnthropicConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anthropicApiKey", GoGetter: "AnthropicApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "anthropicApiKeyInput", GoGetter: "AnthropicApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "anthropicApiKeyPlaintext", GoGetter: "AnthropicApiKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "anthropicApiKeyPlaintextInput", GoGetter: "AnthropicApiKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnthropicApiKey", GoMethod: "ResetAnthropicApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnthropicApiKeyPlaintext", GoMethod: "ResetAnthropicApiKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCohereConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelCohereConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cohereApiBase", GoGetter: "CohereApiBase"},
			_jsii_.MemberProperty{JsiiProperty: "cohereApiBaseInput", GoGetter: "CohereApiBaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "cohereApiKey", GoGetter: "CohereApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "cohereApiKeyInput", GoGetter: "CohereApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cohereApiKeyPlaintext", GoGetter: "CohereApiKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "cohereApiKeyPlaintextInput", GoGetter: "CohereApiKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCohereApiBase", GoMethod: "ResetCohereApiBase"},
			_jsii_.MemberMethod{JsiiMethod: "resetCohereApiKey", GoMethod: "ResetCohereApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetCohereApiKeyPlaintext", GoMethod: "ResetCohereApiKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databricksApiToken", GoGetter: "DatabricksApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "databricksApiTokenInput", GoGetter: "DatabricksApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksApiTokenPlaintext", GoGetter: "DatabricksApiTokenPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "databricksApiTokenPlaintextInput", GoGetter: "DatabricksApiTokenPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksWorkspaceUrl", GoGetter: "DatabricksWorkspaceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "databricksWorkspaceUrlInput", GoGetter: "DatabricksWorkspaceUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabricksApiToken", GoMethod: "ResetDatabricksApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabricksApiTokenPlaintext", GoMethod: "ResetDatabricksApiTokenPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPlaintext", GoGetter: "PrivateKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPlaintextInput", GoGetter: "PrivateKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyPlaintext", GoMethod: "ResetPrivateKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOpenaiConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelOpenaiConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientId", GoGetter: "MicrosoftEntraClientId"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientIdInput", GoGetter: "MicrosoftEntraClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientSecret", GoGetter: "MicrosoftEntraClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientSecretInput", GoGetter: "MicrosoftEntraClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientSecretPlaintext", GoGetter: "MicrosoftEntraClientSecretPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraClientSecretPlaintextInput", GoGetter: "MicrosoftEntraClientSecretPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraTenantId", GoGetter: "MicrosoftEntraTenantId"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftEntraTenantIdInput", GoGetter: "MicrosoftEntraTenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiBase", GoGetter: "OpenaiApiBase"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiBaseInput", GoGetter: "OpenaiApiBaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiKey", GoGetter: "OpenaiApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiKeyInput", GoGetter: "OpenaiApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiKeyPlaintext", GoGetter: "OpenaiApiKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiKeyPlaintextInput", GoGetter: "OpenaiApiKeyPlaintextInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiType", GoGetter: "OpenaiApiType"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiTypeInput", GoGetter: "OpenaiApiTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiVersion", GoGetter: "OpenaiApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "openaiApiVersionInput", GoGetter: "OpenaiApiVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiDeploymentName", GoGetter: "OpenaiDeploymentName"},
			_jsii_.MemberProperty{JsiiProperty: "openaiDeploymentNameInput", GoGetter: "OpenaiDeploymentNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiOrganization", GoGetter: "OpenaiOrganization"},
			_jsii_.MemberProperty{JsiiProperty: "openaiOrganizationInput", GoGetter: "OpenaiOrganizationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftEntraClientId", GoMethod: "ResetMicrosoftEntraClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftEntraClientSecret", GoMethod: "ResetMicrosoftEntraClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftEntraClientSecretPlaintext", GoMethod: "ResetMicrosoftEntraClientSecretPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftEntraTenantId", GoMethod: "ResetMicrosoftEntraTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiApiBase", GoMethod: "ResetOpenaiApiBase"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiApiKey", GoMethod: "ResetOpenaiApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiApiKeyPlaintext", GoMethod: "ResetOpenaiApiKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiApiType", GoMethod: "ResetOpenaiApiType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiApiVersion", GoMethod: "ResetOpenaiApiVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiDeploymentName", GoMethod: "ResetOpenaiDeploymentName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiOrganization", GoMethod: "ResetOpenaiOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsConfig", GoGetter: "Ai21LabsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ai21LabsConfigInput", GoGetter: "Ai21LabsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "amazonBedrockConfig", GoGetter: "AmazonBedrockConfig"},
			_jsii_.MemberProperty{JsiiProperty: "amazonBedrockConfigInput", GoGetter: "AmazonBedrockConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "anthropicConfig", GoGetter: "AnthropicConfig"},
			_jsii_.MemberProperty{JsiiProperty: "anthropicConfigInput", GoGetter: "AnthropicConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cohereConfig", GoGetter: "CohereConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cohereConfigInput", GoGetter: "CohereConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databricksModelServingConfig", GoGetter: "DatabricksModelServingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "databricksModelServingConfigInput", GoGetter: "DatabricksModelServingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudVertexAiConfig", GoGetter: "GoogleCloudVertexAiConfig"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudVertexAiConfigInput", GoGetter: "GoogleCloudVertexAiConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "openaiConfig", GoGetter: "OpenaiConfig"},
			_jsii_.MemberProperty{JsiiProperty: "openaiConfigInput", GoGetter: "OpenaiConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "palmConfig", GoGetter: "PalmConfig"},
			_jsii_.MemberProperty{JsiiProperty: "palmConfigInput", GoGetter: "PalmConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerInput", GoGetter: "ProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAi21LabsConfig", GoMethod: "PutAi21LabsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putAmazonBedrockConfig", GoMethod: "PutAmazonBedrockConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putAnthropicConfig", GoMethod: "PutAnthropicConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putCohereConfig", GoMethod: "PutCohereConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putDatabricksModelServingConfig", GoMethod: "PutDatabricksModelServingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudVertexAiConfig", GoMethod: "PutGoogleCloudVertexAiConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOpenaiConfig", GoMethod: "PutOpenaiConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPalmConfig", GoMethod: "PutPalmConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAi21LabsConfig", GoMethod: "ResetAi21LabsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonBedrockConfig", GoMethod: "ResetAmazonBedrockConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnthropicConfig", GoMethod: "ResetAnthropicConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetCohereConfig", GoMethod: "ResetCohereConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabricksModelServingConfig", GoMethod: "ResetDatabricksModelServingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudVertexAiConfig", GoMethod: "ResetGoogleCloudVertexAiConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenaiConfig", GoMethod: "ResetOpenaiConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetPalmConfig", GoMethod: "ResetPalmConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "task", GoGetter: "Task"},
			_jsii_.MemberProperty{JsiiProperty: "taskInput", GoGetter: "TaskInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelPalmConfig",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelPalmConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "palmApiKey", GoGetter: "PalmApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "palmApiKeyInput", GoGetter: "PalmApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "palmApiKeyPlaintext", GoGetter: "PalmApiKeyPlaintext"},
			_jsii_.MemberProperty{JsiiProperty: "palmApiKeyPlaintextInput", GoGetter: "PalmApiKeyPlaintextInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPalmApiKey", GoMethod: "ResetPalmApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPalmApiKeyPlaintext", GoMethod: "ResetPalmApiKeyPlaintext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesList",
		reflect.TypeOf((*ModelServingConfigServedEntitiesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesOutputReference",
		reflect.TypeOf((*ModelServingConfigServedEntitiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "entityName", GoGetter: "EntityName"},
			_jsii_.MemberProperty{JsiiProperty: "entityNameInput", GoGetter: "EntityNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "entityVersion", GoGetter: "EntityVersion"},
			_jsii_.MemberProperty{JsiiProperty: "entityVersionInput", GoGetter: "EntityVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVars", GoGetter: "EnvironmentVars"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVarsInput", GoGetter: "EnvironmentVarsInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalModel", GoGetter: "ExternalModel"},
			_jsii_.MemberProperty{JsiiProperty: "externalModelInput", GoGetter: "ExternalModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArn", GoGetter: "InstanceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArnInput", GoGetter: "InstanceProfileArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxProvisionedThroughput", GoGetter: "MaxProvisionedThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "maxProvisionedThroughputInput", GoGetter: "MaxProvisionedThroughputInput"},
			_jsii_.MemberProperty{JsiiProperty: "minProvisionedThroughput", GoGetter: "MinProvisionedThroughput"},
			_jsii_.MemberProperty{JsiiProperty: "minProvisionedThroughputInput", GoGetter: "MinProvisionedThroughputInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalModel", GoMethod: "PutExternalModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetEntityName", GoMethod: "ResetEntityName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEntityVersion", GoMethod: "ResetEntityVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentVars", GoMethod: "ResetEnvironmentVars"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalModel", GoMethod: "ResetExternalModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfileArn", GoMethod: "ResetInstanceProfileArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxProvisionedThroughput", GoMethod: "ResetMaxProvisionedThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinProvisionedThroughput", GoMethod: "ResetMinProvisionedThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleToZeroEnabled", GoMethod: "ResetScaleToZeroEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkloadSize", GoMethod: "ResetWorkloadSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkloadType", GoMethod: "ResetWorkloadType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scaleToZeroEnabled", GoGetter: "ScaleToZeroEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "scaleToZeroEnabledInput", GoGetter: "ScaleToZeroEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workloadSize", GoGetter: "WorkloadSize"},
			_jsii_.MemberProperty{JsiiProperty: "workloadSizeInput", GoGetter: "WorkloadSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "workloadType", GoGetter: "WorkloadType"},
			_jsii_.MemberProperty{JsiiProperty: "workloadTypeInput", GoGetter: "WorkloadTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedEntitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedModels",
		reflect.TypeOf((*ModelServingConfigServedModels)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedModelsList",
		reflect.TypeOf((*ModelServingConfigServedModelsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedModelsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedModelsOutputReference",
		reflect.TypeOf((*ModelServingConfigServedModelsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVars", GoGetter: "EnvironmentVars"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVarsInput", GoGetter: "EnvironmentVarsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArn", GoGetter: "InstanceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileArnInput", GoGetter: "InstanceProfileArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelName", GoGetter: "ModelName"},
			_jsii_.MemberProperty{JsiiProperty: "modelNameInput", GoGetter: "ModelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelVersion", GoGetter: "ModelVersion"},
			_jsii_.MemberProperty{JsiiProperty: "modelVersionInput", GoGetter: "ModelVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentVars", GoMethod: "ResetEnvironmentVars"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfileArn", GoMethod: "ResetInstanceProfileArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleToZeroEnabled", GoMethod: "ResetScaleToZeroEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkloadType", GoMethod: "ResetWorkloadType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scaleToZeroEnabled", GoGetter: "ScaleToZeroEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "scaleToZeroEnabledInput", GoGetter: "ScaleToZeroEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workloadSize", GoGetter: "WorkloadSize"},
			_jsii_.MemberProperty{JsiiProperty: "workloadSizeInput", GoGetter: "WorkloadSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "workloadType", GoGetter: "WorkloadType"},
			_jsii_.MemberProperty{JsiiProperty: "workloadTypeInput", GoGetter: "WorkloadTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigServedModelsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigTrafficConfig",
		reflect.TypeOf((*ModelServingConfigTrafficConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigTrafficConfigOutputReference",
		reflect.TypeOf((*ModelServingConfigTrafficConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putRoutes", GoMethod: "PutRoutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoutes", GoMethod: "ResetRoutes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "routes", GoGetter: "Routes"},
			_jsii_.MemberProperty{JsiiProperty: "routesInput", GoGetter: "RoutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigTrafficConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigTrafficConfigRoutes",
		reflect.TypeOf((*ModelServingConfigTrafficConfigRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigTrafficConfigRoutesList",
		reflect.TypeOf((*ModelServingConfigTrafficConfigRoutesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigTrafficConfigRoutesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigTrafficConfigRoutesOutputReference",
		reflect.TypeOf((*ModelServingConfigTrafficConfigRoutesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "servedModelName", GoGetter: "ServedModelName"},
			_jsii_.MemberProperty{JsiiProperty: "servedModelNameInput", GoGetter: "ServedModelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficPercentage", GoGetter: "TrafficPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "trafficPercentageInput", GoGetter: "TrafficPercentageInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingConfigTrafficConfigRoutesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingRateLimits",
		reflect.TypeOf((*ModelServingRateLimits)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingRateLimitsList",
		reflect.TypeOf((*ModelServingRateLimitsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingRateLimitsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingRateLimitsOutputReference",
		reflect.TypeOf((*ModelServingRateLimitsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "calls", GoGetter: "Calls"},
			_jsii_.MemberProperty{JsiiProperty: "callsInput", GoGetter: "CallsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "renewalPeriod", GoGetter: "RenewalPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "renewalPeriodInput", GoGetter: "RenewalPeriodInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingRateLimitsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingTags",
		reflect.TypeOf((*ModelServingTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingTagsList",
		reflect.TypeOf((*ModelServingTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingTagsOutputReference",
		reflect.TypeOf((*ModelServingTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.modelServing.ModelServingTimeouts",
		reflect.TypeOf((*ModelServingTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.modelServing.ModelServingTimeoutsOutputReference",
		reflect.TypeOf((*ModelServingTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_ModelServingTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
