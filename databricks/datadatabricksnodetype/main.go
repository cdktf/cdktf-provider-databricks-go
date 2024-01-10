// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksnodetype

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeType",
		reflect.TypeOf((*DataDatabricksNodeType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
			_jsii_.MemberProperty{JsiiProperty: "categoryInput", GoGetter: "CategoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fleet", GoGetter: "Fleet"},
			_jsii_.MemberProperty{JsiiProperty: "fleetInput", GoGetter: "FleetInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gbPerCore", GoGetter: "GbPerCore"},
			_jsii_.MemberProperty{JsiiProperty: "gbPerCoreInput", GoGetter: "GbPerCoreInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "graviton", GoGetter: "Graviton"},
			_jsii_.MemberProperty{JsiiProperty: "gravitonInput", GoGetter: "GravitonInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isIoCacheEnabled", GoGetter: "IsIoCacheEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isIoCacheEnabledInput", GoGetter: "IsIoCacheEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localDisk", GoGetter: "LocalDisk"},
			_jsii_.MemberProperty{JsiiProperty: "localDiskInput", GoGetter: "LocalDiskInput"},
			_jsii_.MemberProperty{JsiiProperty: "localDiskMinSize", GoGetter: "LocalDiskMinSize"},
			_jsii_.MemberProperty{JsiiProperty: "localDiskMinSizeInput", GoGetter: "LocalDiskMinSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "minCores", GoGetter: "MinCores"},
			_jsii_.MemberProperty{JsiiProperty: "minCoresInput", GoGetter: "MinCoresInput"},
			_jsii_.MemberProperty{JsiiProperty: "minGpus", GoGetter: "MinGpus"},
			_jsii_.MemberProperty{JsiiProperty: "minGpusInput", GoGetter: "MinGpusInput"},
			_jsii_.MemberProperty{JsiiProperty: "minMemoryGb", GoGetter: "MinMemoryGb"},
			_jsii_.MemberProperty{JsiiProperty: "minMemoryGbInput", GoGetter: "MinMemoryGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "photonDriverCapable", GoGetter: "PhotonDriverCapable"},
			_jsii_.MemberProperty{JsiiProperty: "photonDriverCapableInput", GoGetter: "PhotonDriverCapableInput"},
			_jsii_.MemberProperty{JsiiProperty: "photonWorkerCapable", GoGetter: "PhotonWorkerCapable"},
			_jsii_.MemberProperty{JsiiProperty: "photonWorkerCapableInput", GoGetter: "PhotonWorkerCapableInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCategory", GoMethod: "ResetCategory"},
			_jsii_.MemberMethod{JsiiMethod: "resetFleet", GoMethod: "ResetFleet"},
			_jsii_.MemberMethod{JsiiMethod: "resetGbPerCore", GoMethod: "ResetGbPerCore"},
			_jsii_.MemberMethod{JsiiMethod: "resetGraviton", GoMethod: "ResetGraviton"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsIoCacheEnabled", GoMethod: "ResetIsIoCacheEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalDisk", GoMethod: "ResetLocalDisk"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalDiskMinSize", GoMethod: "ResetLocalDiskMinSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinCores", GoMethod: "ResetMinCores"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinGpus", GoMethod: "ResetMinGpus"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinMemoryGb", GoMethod: "ResetMinMemoryGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhotonDriverCapable", GoMethod: "ResetPhotonDriverCapable"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhotonWorkerCapable", GoMethod: "ResetPhotonWorkerCapable"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportPortForwarding", GoMethod: "ResetSupportPortForwarding"},
			_jsii_.MemberProperty{JsiiProperty: "supportPortForwarding", GoGetter: "SupportPortForwarding"},
			_jsii_.MemberProperty{JsiiProperty: "supportPortForwardingInput", GoGetter: "SupportPortForwardingInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksNodeType{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksNodeType.DataDatabricksNodeTypeConfig",
		reflect.TypeOf((*DataDatabricksNodeTypeConfig)(nil)).Elem(),
	)
}
