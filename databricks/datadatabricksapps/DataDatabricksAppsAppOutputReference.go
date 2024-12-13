// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksapps/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAppsAppOutputReference interface {
	cdktf.ComplexObject
	ActiveDeployment() DataDatabricksAppsAppActiveDeploymentOutputReference
	ActiveDeploymentInput() interface{}
	AppStatus() DataDatabricksAppsAppAppStatusOutputReference
	AppStatusInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComputeStatus() DataDatabricksAppsAppComputeStatusOutputReference
	ComputeStatusInput() interface{}
	CreateTime() *string
	SetCreateTime(val *string)
	CreateTimeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Creator() *string
	SetCreator(val *string)
	CreatorInput() *string
	DefaultSourceCodePath() *string
	SetDefaultSourceCodePath(val *string)
	DefaultSourceCodePathInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksAppsApp
	SetInternalValue(val *DataDatabricksAppsApp)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PendingDeployment() DataDatabricksAppsAppPendingDeploymentOutputReference
	PendingDeploymentInput() interface{}
	Resources() DataDatabricksAppsAppResourcesList
	ResourcesInput() interface{}
	ServicePrincipalClientId() *string
	SetServicePrincipalClientId(val *string)
	ServicePrincipalClientIdInput() *string
	ServicePrincipalId() *float64
	SetServicePrincipalId(val *float64)
	ServicePrincipalIdInput() *float64
	ServicePrincipalName() *string
	SetServicePrincipalName(val *string)
	ServicePrincipalNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Updater() *string
	SetUpdater(val *string)
	UpdaterInput() *string
	UpdateTime() *string
	SetUpdateTime(val *string)
	UpdateTimeInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutActiveDeployment(value *DataDatabricksAppsAppActiveDeployment)
	PutAppStatus(value *DataDatabricksAppsAppAppStatus)
	PutComputeStatus(value *DataDatabricksAppsAppComputeStatus)
	PutPendingDeployment(value *DataDatabricksAppsAppPendingDeployment)
	PutResources(value interface{})
	ResetActiveDeployment()
	ResetAppStatus()
	ResetComputeStatus()
	ResetCreateTime()
	ResetCreator()
	ResetDefaultSourceCodePath()
	ResetDescription()
	ResetPendingDeployment()
	ResetResources()
	ResetServicePrincipalClientId()
	ResetServicePrincipalId()
	ResetServicePrincipalName()
	ResetUpdater()
	ResetUpdateTime()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAppsAppOutputReference
type jsiiProxy_DataDatabricksAppsAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ActiveDeployment() DataDatabricksAppsAppActiveDeploymentOutputReference {
	var returns DataDatabricksAppsAppActiveDeploymentOutputReference
	_jsii_.Get(
		j,
		"activeDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ActiveDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) AppStatus() DataDatabricksAppsAppAppStatusOutputReference {
	var returns DataDatabricksAppsAppAppStatusOutputReference
	_jsii_.Get(
		j,
		"appStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) AppStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ComputeStatus() DataDatabricksAppsAppComputeStatusOutputReference {
	var returns DataDatabricksAppsAppComputeStatusOutputReference
	_jsii_.Get(
		j,
		"computeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ComputeStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) CreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) CreatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) DefaultSourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) DefaultSourceCodePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) InternalValue() *DataDatabricksAppsApp {
	var returns *DataDatabricksAppsApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) PendingDeployment() DataDatabricksAppsAppPendingDeploymentOutputReference {
	var returns DataDatabricksAppsAppPendingDeploymentOutputReference
	_jsii_.Get(
		j,
		"pendingDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) PendingDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pendingDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Resources() DataDatabricksAppsAppResourcesList {
	var returns DataDatabricksAppsAppResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) ServicePrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Updater() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updater",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) UpdaterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updaterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) UpdateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsAppOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsAppOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsAppOutputReference_Override(d DataDatabricksAppsAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApps.DataDatabricksAppsAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetCreateTime(val *string) {
	if err := j.validateSetCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetCreator(val *string) {
	if err := j.validateSetCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creator",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetDefaultSourceCodePath(val *string) {
	if err := j.validateSetDefaultSourceCodePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSourceCodePath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetInternalValue(val *DataDatabricksAppsApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetServicePrincipalClientId(val *string) {
	if err := j.validateSetServicePrincipalClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalClientId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetServicePrincipalId(val *float64) {
	if err := j.validateSetServicePrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetServicePrincipalName(val *string) {
	if err := j.validateSetServicePrincipalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetUpdater(val *string) {
	if err := j.validateSetUpdaterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updater",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetUpdateTime(val *string) {
	if err := j.validateSetUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsAppOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) PutActiveDeployment(value *DataDatabricksAppsAppActiveDeployment) {
	if err := d.validatePutActiveDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putActiveDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) PutAppStatus(value *DataDatabricksAppsAppAppStatus) {
	if err := d.validatePutAppStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAppStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) PutComputeStatus(value *DataDatabricksAppsAppComputeStatus) {
	if err := d.validatePutComputeStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComputeStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) PutPendingDeployment(value *DataDatabricksAppsAppPendingDeployment) {
	if err := d.validatePutPendingDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPendingDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) PutResources(value interface{}) {
	if err := d.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetActiveDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetAppStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetAppStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetComputeStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetCreateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetCreator() {
	_jsii_.InvokeVoid(
		d,
		"resetCreator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetDefaultSourceCodePath() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultSourceCodePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetPendingDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetPendingDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetServicePrincipalClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetServicePrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetServicePrincipalName() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetUpdater() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdater",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetUpdateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

