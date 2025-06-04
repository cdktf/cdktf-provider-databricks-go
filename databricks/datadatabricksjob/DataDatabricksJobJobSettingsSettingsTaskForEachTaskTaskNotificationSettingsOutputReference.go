// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference interface {
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
	InternalValue() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings
	SetInternalValue(val *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings)
	NoAlertForCanceledRuns() interface{}
	SetNoAlertForCanceledRuns(val interface{})
	NoAlertForCanceledRunsInput() interface{}
	NoAlertForSkippedRuns() interface{}
	SetNoAlertForSkippedRuns(val interface{})
	NoAlertForSkippedRunsInput() interface{}
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
	ResetNoAlertForCanceledRuns()
	ResetNoAlertForSkippedRuns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) AlertOnLastAttempt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertOnLastAttempt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) AlertOnLastAttemptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertOnLastAttemptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) InternalValue() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) NoAlertForCanceledRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForCanceledRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) NoAlertForCanceledRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForCanceledRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) NoAlertForSkippedRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) NoAlertForSkippedRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAlertForSkippedRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetAlertOnLastAttempt(val interface{}) {
	if err := j.validateSetAlertOnLastAttemptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertOnLastAttempt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetInternalValue(val *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetNoAlertForCanceledRuns(val interface{}) {
	if err := j.validateSetNoAlertForCanceledRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAlertForCanceledRuns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetNoAlertForSkippedRuns(val interface{}) {
	if err := j.validateSetNoAlertForSkippedRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAlertForSkippedRuns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ResetAlertOnLastAttempt() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertOnLastAttempt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ResetNoAlertForCanceledRuns() {
	_jsii_.InvokeVoid(
		d,
		"resetNoAlertForCanceledRuns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ResetNoAlertForSkippedRuns() {
	_jsii_.InvokeVoid(
		d,
		"resetNoAlertForSkippedRuns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNotificationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

