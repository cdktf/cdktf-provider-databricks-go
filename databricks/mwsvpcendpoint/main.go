// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsvpcendpoint

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpoint",
		reflect.TypeOf((*MwsVpcEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountIdInput", GoGetter: "AwsAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsEndpointServiceId", GoGetter: "AwsEndpointServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "awsEndpointServiceIdInput", GoGetter: "AwsEndpointServiceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsVpcEndpointId", GoGetter: "AwsVpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "awsVpcEndpointIdInput", GoGetter: "AwsVpcEndpointIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "gcpVpcEndpointInfo", GoGetter: "GcpVpcEndpointInfo"},
			_jsii_.MemberProperty{JsiiProperty: "gcpVpcEndpointInfoInput", GoGetter: "GcpVpcEndpointInfoInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putGcpVpcEndpointInfo", GoMethod: "PutGcpVpcEndpointInfo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccountId", GoMethod: "ResetAwsAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsEndpointServiceId", GoMethod: "ResetAwsEndpointServiceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsVpcEndpointId", GoMethod: "ResetAwsVpcEndpointId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGcpVpcEndpointInfo", GoMethod: "ResetGcpVpcEndpointInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseCase", GoMethod: "ResetUseCase"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcEndpointId", GoMethod: "ResetVpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useCase", GoGetter: "UseCase"},
			_jsii_.MemberProperty{JsiiProperty: "useCaseInput", GoGetter: "UseCaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointId", GoGetter: "VpcEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointIdInput", GoGetter: "VpcEndpointIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointName", GoGetter: "VpcEndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointNameInput", GoGetter: "VpcEndpointNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_MwsVpcEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpointConfig",
		reflect.TypeOf((*MwsVpcEndpointConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpointGcpVpcEndpointInfo",
		reflect.TypeOf((*MwsVpcEndpointGcpVpcEndpointInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpointGcpVpcEndpointInfoOutputReference",
		reflect.TypeOf((*MwsVpcEndpointGcpVpcEndpointInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointRegion", GoGetter: "EndpointRegion"},
			_jsii_.MemberProperty{JsiiProperty: "endpointRegionInput", GoGetter: "EndpointRegionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pscConnectionId", GoGetter: "PscConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "pscConnectionIdInput", GoGetter: "PscConnectionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pscEndpointName", GoGetter: "PscEndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "pscEndpointNameInput", GoGetter: "PscEndpointNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPscConnectionId", GoMethod: "ResetPscConnectionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAttachmentId", GoMethod: "ResetServiceAttachmentId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAttachmentId", GoGetter: "ServiceAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAttachmentIdInput", GoGetter: "ServiceAttachmentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
