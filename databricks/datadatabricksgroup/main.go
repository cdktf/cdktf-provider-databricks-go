package datadatabricksgroup

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		reflect.TypeOf((*DataDatabricksGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowClusterCreate", GoGetter: "AllowClusterCreate"},
			_jsii_.MemberProperty{JsiiProperty: "allowClusterCreateInput", GoGetter: "AllowClusterCreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowInstancePoolCreate", GoGetter: "AllowInstancePoolCreate"},
			_jsii_.MemberProperty{JsiiProperty: "allowInstancePoolCreateInput", GoGetter: "AllowInstancePoolCreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "childGroups", GoGetter: "ChildGroups"},
			_jsii_.MemberProperty{JsiiProperty: "childGroupsInput", GoGetter: "ChildGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "databricksSqlAccess", GoGetter: "DatabricksSqlAccess"},
			_jsii_.MemberProperty{JsiiProperty: "databricksSqlAccessInput", GoGetter: "DatabricksSqlAccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "groupsInput", GoGetter: "GroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfiles", GoGetter: "InstanceProfiles"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfilesInput", GoGetter: "InstanceProfilesInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membersInput", GoGetter: "MembersInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recursive", GoGetter: "Recursive"},
			_jsii_.MemberProperty{JsiiProperty: "recursiveInput", GoGetter: "RecursiveInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowClusterCreate", GoMethod: "ResetAllowClusterCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowInstancePoolCreate", GoMethod: "ResetAllowInstancePoolCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetChildGroups", GoMethod: "ResetChildGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabricksSqlAccess", GoMethod: "ResetDatabricksSqlAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalId", GoMethod: "ResetExternalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroups", GoMethod: "ResetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceProfiles", GoMethod: "ResetInstanceProfiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembers", GoMethod: "ResetMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecursive", GoMethod: "ResetRecursive"},
			_jsii_.MemberMethod{JsiiMethod: "resetServicePrincipals", GoMethod: "ResetServicePrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsers", GoMethod: "ResetUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceAccess", GoMethod: "ResetWorkspaceAccess"},
			_jsii_.MemberProperty{JsiiProperty: "servicePrincipals", GoGetter: "ServicePrincipals"},
			_jsii_.MemberProperty{JsiiProperty: "servicePrincipalsInput", GoGetter: "ServicePrincipalsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "usersInput", GoGetter: "UsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccess", GoGetter: "WorkspaceAccess"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccessInput", GoGetter: "WorkspaceAccessInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroupConfig",
		reflect.TypeOf((*DataDatabricksGroupConfig)(nil)).Elem(),
	)
}
