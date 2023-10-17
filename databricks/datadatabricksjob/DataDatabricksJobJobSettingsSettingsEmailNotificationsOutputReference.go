// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference interface {
	cdktf.ComplexObject
	AlertOnLastAttempt() interface{}
	SetAlertOnLastAttempt(val interface{})
	AlertOnLastAttemptInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksJobJobSettingsSettingsEmailNotifications
	SetInternalValue(val *DataDatabricksJobJobSettingsSettingsEmailNotifications)
	NoAlertForSkippedRuns() interface{}
	SetNoAlertForSkippedRuns(val interface{})
	NoAlertForSkippedRunsInput() interface{}
	OnDurationWarningThresholdExceeded() *[]*string
	SetOnDurationWarningThresholdExceeded(val *[]*string)
	OnDurationWarningThresholdExceededInput() *[]*string
	OnFailure() *[]*string
	SetOnFailure(val *[]*string)
	OnFailureInput() *[]*string
	OnStart() *[]*string
	SetOnStart(val *[]*string)
	OnStartInput() *[]*string
	OnSuccess() *[]*string
	SetOnSuccess(val *[]*string)
	OnSuccessInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAlertOnLastAttempt()
	ResetNoAlertForSkippedRuns()
	ResetOnDurationWarningThresholdExceeded()
	ResetOnFailure()
	ResetOnStart()
	ResetOnSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) AlertOnLastAttempt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertOnLastAttempt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) AlertOnLastAttemptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertOnLastAttemptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) InternalValue() *DataDatabricksJobJobSettingsSettingsEmailNotifications {
	var returns *DataDatabricksJobJobSettingsSettingsEmailNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) NoAlertForSkippedRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) NoAlertForSkippedRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnDurationWarningThresholdExceeded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnDurationWarningThresholdExceededInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onDurationWarningThresholdExceededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnFailure() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnFailureInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnStart() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnStartInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnSuccess() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) OnSuccessInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference_Override(d DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetAlertOnLastAttempt(val interface{}) {
	if err := j.validateSetAlertOnLastAttemptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertOnLastAttempt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetInternalValue(val *DataDatabricksJobJobSettingsSettingsEmailNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetNoAlertForSkippedRuns(val interface{}) {
	if err := j.validateSetNoAlertForSkippedRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAlertForSkippedRuns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetOnDurationWarningThresholdExceeded(val *[]*string) {
	if err := j.validateSetOnDurationWarningThresholdExceededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDurationWarningThresholdExceeded",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetOnFailure(val *[]*string) {
	if err := j.validateSetOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFailure",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetOnStart(val *[]*string) {
	if err := j.validateSetOnStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStart",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetOnSuccess(val *[]*string) {
	if err := j.validateSetOnSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSuccess",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetAlertOnLastAttempt() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertOnLastAttempt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetNoAlertForSkippedRuns() {
	_jsii_.InvokeVoid(
		d,
		"resetNoAlertForSkippedRuns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetOnDurationWarningThresholdExceeded() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDurationWarningThresholdExceeded",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetOnFailure() {
	_jsii_.InvokeVoid(
		d,
		"resetOnFailure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetOnStart() {
	_jsii_.InvokeVoid(
		d,
		"resetOnStart",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ResetOnSuccess() {
	_jsii_.InvokeVoid(
		d,
		"resetOnSuccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsEmailNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

