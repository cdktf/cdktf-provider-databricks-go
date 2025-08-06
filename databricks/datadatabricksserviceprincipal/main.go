// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksserviceprincipal

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipal.DataDatabricksServicePrincipal",
		reflect.TypeOf((*DataDatabricksServicePrincipal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aclPrincipalId", GoGetter: "AclPrincipalId"},
			_jsii_.MemberProperty{JsiiProperty: "aclPrincipalIdInput", GoGetter: "AclPrincipalIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "active", GoGetter: "Active"},
			_jsii_.MemberProperty{JsiiProperty: "activeInput", GoGetter: "ActiveInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "applicationIdInput", GoGetter: "ApplicationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalId", GoGetter: "ExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdInput", GoGetter: "ExternalIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "home", GoGetter: "Home"},
			_jsii_.MemberProperty{JsiiProperty: "homeInput", GoGetter: "HomeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "repos", GoGetter: "Repos"},
			_jsii_.MemberProperty{JsiiProperty: "reposInput", GoGetter: "ReposInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAclPrincipalId", GoMethod: "ResetAclPrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetActive", GoMethod: "ResetActive"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationId", GoMethod: "ResetApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalId", GoMethod: "ResetExternalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHome", GoMethod: "ResetHome"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepos", GoMethod: "ResetRepos"},
			_jsii_.MemberMethod{JsiiMethod: "resetScimId", GoMethod: "ResetScimId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpId", GoMethod: "ResetSpId"},
			_jsii_.MemberProperty{JsiiProperty: "scimId", GoGetter: "ScimId"},
			_jsii_.MemberProperty{JsiiProperty: "scimIdInput", GoGetter: "ScimIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "spId", GoGetter: "SpId"},
			_jsii_.MemberProperty{JsiiProperty: "spIdInput", GoGetter: "SpIdInput"},
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
			j := jsiiProxy_DataDatabricksServicePrincipal{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksServicePrincipal.DataDatabricksServicePrincipalConfig",
		reflect.TypeOf((*DataDatabricksServicePrincipalConfig)(nil)).Elem(),
	)
}
