// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwslogdelivery

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDelivery",
		reflect.TypeOf((*MwsLogDelivery)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configId", GoGetter: "ConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "configIdInput", GoGetter: "ConfigIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "configName", GoGetter: "ConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "configNameInput", GoGetter: "ConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsId", GoGetter: "CredentialsId"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsIdInput", GoGetter: "CredentialsIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryPathPrefix", GoGetter: "DeliveryPathPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryPathPrefixInput", GoGetter: "DeliveryPathPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStartTime", GoGetter: "DeliveryStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStartTimeInput", GoGetter: "DeliveryStartTimeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberProperty{JsiiProperty: "logTypeInput", GoGetter: "LogTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormatInput", GoGetter: "OutputFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigId", GoMethod: "ResetConfigId"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigName", GoMethod: "ResetConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryPathPrefix", GoMethod: "ResetDeliveryPathPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryStartTime", GoMethod: "ResetDeliveryStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceIdsFilter", GoMethod: "ResetWorkspaceIdsFilter"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageConfigurationId", GoGetter: "StorageConfigurationId"},
			_jsii_.MemberProperty{JsiiProperty: "storageConfigurationIdInput", GoGetter: "StorageConfigurationIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsFilter", GoGetter: "WorkspaceIdsFilter"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdsFilterInput", GoGetter: "WorkspaceIdsFilterInput"},
		},
		func() interface{} {
			j := jsiiProxy_MwsLogDelivery{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.mwsLogDelivery.MwsLogDeliveryConfig",
		reflect.TypeOf((*MwsLogDeliveryConfig)(nil)).Elem(),
	)
}
