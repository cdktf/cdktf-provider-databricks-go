// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmetastore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastore",
		reflect.TypeOf((*DataDatabricksMetastore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreId", GoGetter: "MetastoreId"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreIdInput", GoGetter: "MetastoreIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreInfo", GoGetter: "MetastoreInfo"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreInfoInput", GoGetter: "MetastoreInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putMetastoreInfo", GoMethod: "PutMetastoreInfo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetastoreInfo", GoMethod: "ResetMetastoreInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksMetastore{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastoreConfig",
		reflect.TypeOf((*DataDatabricksMetastoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastoreMetastoreInfo",
		reflect.TypeOf((*DataDatabricksMetastoreMetastoreInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastoreMetastoreInfoOutputReference",
		reflect.TypeOf((*DataDatabricksMetastoreMetastoreInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloud", GoGetter: "Cloud"},
			_jsii_.MemberProperty{JsiiProperty: "cloudInput", GoGetter: "CloudInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "createdAtInput", GoGetter: "CreatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdBy", GoGetter: "CreatedBy"},
			_jsii_.MemberProperty{JsiiProperty: "createdByInput", GoGetter: "CreatedByInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDataAccessConfigId", GoGetter: "DefaultDataAccessConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDataAccessConfigIdInput", GoGetter: "DefaultDataAccessConfigIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingOrganizationName", GoGetter: "DeltaSharingOrganizationName"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingOrganizationNameInput", GoGetter: "DeltaSharingOrganizationNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingRecipientTokenLifetimeInSeconds", GoGetter: "DeltaSharingRecipientTokenLifetimeInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingRecipientTokenLifetimeInSecondsInput", GoGetter: "DeltaSharingRecipientTokenLifetimeInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingScope", GoGetter: "DeltaSharingScope"},
			_jsii_.MemberProperty{JsiiProperty: "deltaSharingScopeInput", GoGetter: "DeltaSharingScopeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "globalMetastoreId", GoGetter: "GlobalMetastoreId"},
			_jsii_.MemberProperty{JsiiProperty: "globalMetastoreIdInput", GoGetter: "GlobalMetastoreIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreId", GoGetter: "MetastoreId"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreIdInput", GoGetter: "MetastoreIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerInput", GoGetter: "OwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "privilegeModelVersion", GoGetter: "PrivilegeModelVersion"},
			_jsii_.MemberProperty{JsiiProperty: "privilegeModelVersionInput", GoGetter: "PrivilegeModelVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloud", GoMethod: "ResetCloud"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedAt", GoMethod: "ResetCreatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedBy", GoMethod: "ResetCreatedBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultDataAccessConfigId", GoMethod: "ResetDefaultDataAccessConfigId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeltaSharingOrganizationName", GoMethod: "ResetDeltaSharingOrganizationName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeltaSharingRecipientTokenLifetimeInSeconds", GoMethod: "ResetDeltaSharingRecipientTokenLifetimeInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeltaSharingScope", GoMethod: "ResetDeltaSharingScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetGlobalMetastoreId", GoMethod: "ResetGlobalMetastoreId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetastoreId", GoMethod: "ResetMetastoreId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwner", GoMethod: "ResetOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivilegeModelVersion", GoMethod: "ResetPrivilegeModelVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageRoot", GoMethod: "ResetStorageRoot"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageRootCredentialId", GoMethod: "ResetStorageRootCredentialId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageRootCredentialName", GoMethod: "ResetStorageRootCredentialName"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdatedAt", GoMethod: "ResetUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdatedBy", GoMethod: "ResetUpdatedBy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageRoot", GoGetter: "StorageRoot"},
			_jsii_.MemberProperty{JsiiProperty: "storageRootCredentialId", GoGetter: "StorageRootCredentialId"},
			_jsii_.MemberProperty{JsiiProperty: "storageRootCredentialIdInput", GoGetter: "StorageRootCredentialIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageRootCredentialName", GoGetter: "StorageRootCredentialName"},
			_jsii_.MemberProperty{JsiiProperty: "storageRootCredentialNameInput", GoGetter: "StorageRootCredentialNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageRootInput", GoGetter: "StorageRootInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAtInput", GoGetter: "UpdatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "updatedBy", GoGetter: "UpdatedBy"},
			_jsii_.MemberProperty{JsiiProperty: "updatedByInput", GoGetter: "UpdatedByInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}