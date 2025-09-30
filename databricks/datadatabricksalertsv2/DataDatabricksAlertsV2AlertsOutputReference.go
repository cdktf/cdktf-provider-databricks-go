// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksalertsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAlertsV2AlertsOutputReference interface {
	cdktf.ComplexObject
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomDescription() *string
	SetCustomDescription(val *string)
	CustomDescriptionInput() *string
	CustomSummary() *string
	SetCustomSummary(val *string)
	CustomSummaryInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveRunAs() DataDatabricksAlertsV2AlertsEffectiveRunAsOutputReference
	Evaluation() DataDatabricksAlertsV2AlertsEvaluationOutputReference
	EvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataDatabricksAlertsV2Alerts
	SetInternalValue(val *DataDatabricksAlertsV2Alerts)
	LifecycleState() *string
	OwnerUserName() *string
	ParentPath() *string
	SetParentPath(val *string)
	ParentPathInput() *string
	QueryText() *string
	SetQueryText(val *string)
	QueryTextInput() *string
	RunAs() DataDatabricksAlertsV2AlertsRunAsOutputReference
	RunAsInput() interface{}
	RunAsUserName() *string
	SetRunAsUserName(val *string)
	RunAsUserNameInput() *string
	Schedule() DataDatabricksAlertsV2AlertsScheduleOutputReference
	ScheduleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
	WarehouseId() *string
	SetWarehouseId(val *string)
	WarehouseIdInput() *string
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
	PutEvaluation(value *DataDatabricksAlertsV2AlertsEvaluation)
	PutRunAs(value *DataDatabricksAlertsV2AlertsRunAs)
	PutSchedule(value *DataDatabricksAlertsV2AlertsSchedule)
	ResetCustomDescription()
	ResetCustomSummary()
	ResetDisplayName()
	ResetEvaluation()
	ResetParentPath()
	ResetQueryText()
	ResetRunAs()
	ResetRunAsUserName()
	ResetSchedule()
	ResetWarehouseId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAlertsV2AlertsOutputReference
type jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CustomDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CustomDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CustomSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) CustomSummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) EffectiveRunAs() DataDatabricksAlertsV2AlertsEffectiveRunAsOutputReference {
	var returns DataDatabricksAlertsV2AlertsEffectiveRunAsOutputReference
	_jsii_.Get(
		j,
		"effectiveRunAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) Evaluation() DataDatabricksAlertsV2AlertsEvaluationOutputReference {
	var returns DataDatabricksAlertsV2AlertsEvaluationOutputReference
	_jsii_.Get(
		j,
		"evaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) EvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) InternalValue() *DataDatabricksAlertsV2Alerts {
	var returns *DataDatabricksAlertsV2Alerts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) OwnerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ParentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ParentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) QueryText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) QueryTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) RunAs() DataDatabricksAlertsV2AlertsRunAsOutputReference {
	var returns DataDatabricksAlertsV2AlertsRunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) RunAsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) RunAsUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) RunAsUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) Schedule() DataDatabricksAlertsV2AlertsScheduleOutputReference {
	var returns DataDatabricksAlertsV2AlertsScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAlertsV2AlertsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAlertsV2AlertsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAlertsV2AlertsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2AlertsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAlertsV2AlertsOutputReference_Override(d DataDatabricksAlertsV2AlertsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2AlertsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetCustomDescription(val *string) {
	if err := j.validateSetCustomDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDescription",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetCustomSummary(val *string) {
	if err := j.validateSetCustomSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSummary",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetInternalValue(val *DataDatabricksAlertsV2Alerts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetParentPath(val *string) {
	if err := j.validateSetParentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetQueryText(val *string) {
	if err := j.validateSetQueryTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryText",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetRunAsUserName(val *string) {
	if err := j.validateSetRunAsUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) PutEvaluation(value *DataDatabricksAlertsV2AlertsEvaluation) {
	if err := d.validatePutEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvaluation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) PutRunAs(value *DataDatabricksAlertsV2AlertsRunAs) {
	if err := d.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunAs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) PutSchedule(value *DataDatabricksAlertsV2AlertsSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetCustomDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetCustomSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetEvaluation() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetParentPath() {
	_jsii_.InvokeVoid(
		d,
		"resetParentPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetQueryText() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetRunAs() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetRunAsUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetRunAsUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2AlertsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

