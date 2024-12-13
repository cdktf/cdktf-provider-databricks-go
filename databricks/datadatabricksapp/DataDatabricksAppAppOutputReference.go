// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAppAppOutputReference interface {
	cdktf.ComplexObject
	ActiveDeployment() DataDatabricksAppAppActiveDeploymentOutputReference
	ActiveDeploymentInput() interface{}
	AppStatus() DataDatabricksAppAppAppStatusOutputReference
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
	ComputeStatus() DataDatabricksAppAppComputeStatusOutputReference
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
	InternalValue() *DataDatabricksAppApp
	SetInternalValue(val *DataDatabricksAppApp)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PendingDeployment() DataDatabricksAppAppPendingDeploymentOutputReference
	PendingDeploymentInput() interface{}
	Resources() DataDatabricksAppAppResourcesList
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
	PutActiveDeployment(value *DataDatabricksAppAppActiveDeployment)
	PutAppStatus(value *DataDatabricksAppAppAppStatus)
	PutComputeStatus(value *DataDatabricksAppAppComputeStatus)
	PutPendingDeployment(value *DataDatabricksAppAppPendingDeployment)
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

// The jsii proxy struct for DataDatabricksAppAppOutputReference
type jsiiProxy_DataDatabricksAppAppOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ActiveDeployment() DataDatabricksAppAppActiveDeploymentOutputReference {
	var returns DataDatabricksAppAppActiveDeploymentOutputReference
	_jsii_.Get(
		j,
		"activeDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ActiveDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) AppStatus() DataDatabricksAppAppAppStatusOutputReference {
	var returns DataDatabricksAppAppAppStatusOutputReference
	_jsii_.Get(
		j,
		"appStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) AppStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeStatus() DataDatabricksAppAppComputeStatusOutputReference {
	var returns DataDatabricksAppAppComputeStatusOutputReference
	_jsii_.Get(
		j,
		"computeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Creator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) CreatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DefaultSourceCodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DefaultSourceCodePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSourceCodePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) InternalValue() *DataDatabricksAppApp {
	var returns *DataDatabricksAppApp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) PendingDeployment() DataDatabricksAppAppPendingDeploymentOutputReference {
	var returns DataDatabricksAppAppPendingDeploymentOutputReference
	_jsii_.Get(
		j,
		"pendingDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) PendingDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pendingDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Resources() DataDatabricksAppAppResourcesList {
	var returns DataDatabricksAppAppResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"servicePrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) ServicePrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Updater() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updater",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UpdaterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updaterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UpdateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppAppOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksAppAppOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppAppOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppAppOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApp.DataDatabricksAppAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksAppAppOutputReference_Override(d DataDatabricksAppAppOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksApp.DataDatabricksAppAppOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetCreateTime(val *string) {
	if err := j.validateSetCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetCreator(val *string) {
	if err := j.validateSetCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creator",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetDefaultSourceCodePath(val *string) {
	if err := j.validateSetDefaultSourceCodePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSourceCodePath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetInternalValue(val *DataDatabricksAppApp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetServicePrincipalClientId(val *string) {
	if err := j.validateSetServicePrincipalClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalClientId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetServicePrincipalId(val *float64) {
	if err := j.validateSetServicePrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetServicePrincipalName(val *string) {
	if err := j.validateSetServicePrincipalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetUpdater(val *string) {
	if err := j.validateSetUpdaterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updater",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetUpdateTime(val *string) {
	if err := j.validateSetUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppAppOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutActiveDeployment(value *DataDatabricksAppAppActiveDeployment) {
	if err := d.validatePutActiveDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putActiveDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutAppStatus(value *DataDatabricksAppAppAppStatus) {
	if err := d.validatePutAppStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAppStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutComputeStatus(value *DataDatabricksAppAppComputeStatus) {
	if err := d.validatePutComputeStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComputeStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutPendingDeployment(value *DataDatabricksAppAppPendingDeployment) {
	if err := d.validatePutPendingDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPendingDeployment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) PutResources(value interface{}) {
	if err := d.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetActiveDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetAppStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetAppStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetComputeStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetCreateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetCreator() {
	_jsii_.InvokeVoid(
		d,
		"resetCreator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetDefaultSourceCodePath() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultSourceCodePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetPendingDeployment() {
	_jsii_.InvokeVoid(
		d,
		"resetPendingDeployment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetServicePrincipalClientId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalClientId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetServicePrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetServicePrincipalName() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipalName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetUpdater() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdater",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetUpdateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppAppOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

