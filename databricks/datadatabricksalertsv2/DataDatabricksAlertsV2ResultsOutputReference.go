// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksalertsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksalertsv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAlertsV2ResultsOutputReference interface {
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
	Evaluation() DataDatabricksAlertsV2ResultsEvaluationOutputReference
	EvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataDatabricksAlertsV2Results
	SetInternalValue(val *DataDatabricksAlertsV2Results)
	LifecycleState() *string
	OwnerUserName() *string
	ParentPath() *string
	SetParentPath(val *string)
	ParentPathInput() *string
	QueryText() *string
	SetQueryText(val *string)
	QueryTextInput() *string
	RunAsUserName() *string
	Schedule() DataDatabricksAlertsV2ResultsScheduleOutputReference
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
	PutEvaluation(value *DataDatabricksAlertsV2ResultsEvaluation)
	PutSchedule(value *DataDatabricksAlertsV2ResultsSchedule)
	ResetCustomDescription()
	ResetCustomSummary()
	ResetDisplayName()
	ResetEvaluation()
	ResetParentPath()
	ResetQueryText()
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

// The jsii proxy struct for DataDatabricksAlertsV2ResultsOutputReference
type jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CustomDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CustomDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CustomSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) CustomSummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) Evaluation() DataDatabricksAlertsV2ResultsEvaluationOutputReference {
	var returns DataDatabricksAlertsV2ResultsEvaluationOutputReference
	_jsii_.Get(
		j,
		"evaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) EvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) InternalValue() *DataDatabricksAlertsV2Results {
	var returns *DataDatabricksAlertsV2Results
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) OwnerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ParentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ParentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) QueryText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) QueryTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) RunAsUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) Schedule() DataDatabricksAlertsV2ResultsScheduleOutputReference {
	var returns DataDatabricksAlertsV2ResultsScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) WarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) WarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseIdInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAlertsV2ResultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAlertsV2ResultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAlertsV2ResultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2ResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAlertsV2ResultsOutputReference_Override(d DataDatabricksAlertsV2ResultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAlertsV2.DataDatabricksAlertsV2ResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetCustomDescription(val *string) {
	if err := j.validateSetCustomDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDescription",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetCustomSummary(val *string) {
	if err := j.validateSetCustomSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSummary",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetInternalValue(val *DataDatabricksAlertsV2Results) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetParentPath(val *string) {
	if err := j.validateSetParentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetQueryText(val *string) {
	if err := j.validateSetQueryTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryText",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference)SetWarehouseId(val *string) {
	if err := j.validateSetWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseId",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) PutEvaluation(value *DataDatabricksAlertsV2ResultsEvaluation) {
	if err := d.validatePutEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvaluation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) PutSchedule(value *DataDatabricksAlertsV2ResultsSchedule) {
	if err := d.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetCustomDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetCustomSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetEvaluation() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetParentPath() {
	_jsii_.InvokeVoid(
		d,
		"resetParentPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetQueryText() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ResetWarehouseId() {
	_jsii_.InvokeVoid(
		d,
		"resetWarehouseId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAlertsV2ResultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

