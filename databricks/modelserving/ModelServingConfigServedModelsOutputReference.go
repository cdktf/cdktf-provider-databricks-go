// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedModelsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnvironmentVars() *map[string]*string
	SetEnvironmentVars(val *map[string]*string)
	EnvironmentVarsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxProvisionedConcurrency() *float64
	SetMaxProvisionedConcurrency(val *float64)
	MaxProvisionedConcurrencyInput() *float64
	MaxProvisionedThroughput() *float64
	SetMaxProvisionedThroughput(val *float64)
	MaxProvisionedThroughputInput() *float64
	MinProvisionedConcurrency() *float64
	SetMinProvisionedConcurrency(val *float64)
	MinProvisionedConcurrencyInput() *float64
	MinProvisionedThroughput() *float64
	SetMinProvisionedThroughput(val *float64)
	MinProvisionedThroughputInput() *float64
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
	ModelVersion() *string
	SetModelVersion(val *string)
	ModelVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ProvisionedModelUnits() *float64
	SetProvisionedModelUnits(val *float64)
	ProvisionedModelUnitsInput() *float64
	ScaleToZeroEnabled() interface{}
	SetScaleToZeroEnabled(val interface{})
	ScaleToZeroEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadSize() *string
	SetWorkloadSize(val *string)
	WorkloadSizeInput() *string
	WorkloadType() *string
	SetWorkloadType(val *string)
	WorkloadTypeInput() *string
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
	ResetEnvironmentVars()
	ResetInstanceProfileArn()
	ResetMaxProvisionedConcurrency()
	ResetMaxProvisionedThroughput()
	ResetMinProvisionedConcurrency()
	ResetMinProvisionedThroughput()
	ResetName()
	ResetProvisionedModelUnits()
	ResetScaleToZeroEnabled()
	ResetWorkloadSize()
	ResetWorkloadType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedModelsOutputReference
type jsiiProxy_ModelServingConfigServedModelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) EnvironmentVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) EnvironmentVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MaxProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MaxProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MaxProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MaxProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MinProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MinProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MinProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) MinProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ProvisionedModelUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedModelUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ProvisionedModelUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedModelUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ScaleToZeroEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleToZeroEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) ScaleToZeroEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleToZeroEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) WorkloadSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) WorkloadSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) WorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference) WorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedModelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ModelServingConfigServedModelsOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedModelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedModelsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedModelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewModelServingConfigServedModelsOutputReference_Override(m ModelServingConfigServedModelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedModelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetEnvironmentVars(val *map[string]*string) {
	if err := j.validateSetEnvironmentVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVars",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetMaxProvisionedConcurrency(val *float64) {
	if err := j.validateSetMaxProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProvisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetMaxProvisionedThroughput(val *float64) {
	if err := j.validateSetMaxProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetMinProvisionedConcurrency(val *float64) {
	if err := j.validateSetMinProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProvisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetMinProvisionedThroughput(val *float64) {
	if err := j.validateSetMinProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetModelVersion(val *string) {
	if err := j.validateSetModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelVersion",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetProvisionedModelUnits(val *float64) {
	if err := j.validateSetProvisionedModelUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedModelUnits",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetScaleToZeroEnabled(val interface{}) {
	if err := j.validateSetScaleToZeroEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleToZeroEnabled",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetWorkloadSize(val *string) {
	if err := j.validateSetWorkloadSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadSize",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedModelsOutputReference)SetWorkloadType(val *string) {
	if err := j.validateSetWorkloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadType",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetEnvironmentVars() {
	_jsii_.InvokeVoid(
		m,
		"resetEnvironmentVars",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		m,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetMaxProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxProvisionedConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetMaxProvisionedThroughput() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxProvisionedThroughput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetMinProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinProvisionedConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetMinProvisionedThroughput() {
	_jsii_.InvokeVoid(
		m,
		"resetMinProvisionedThroughput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetProvisionedModelUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetProvisionedModelUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetScaleToZeroEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetScaleToZeroEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetWorkloadSize() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkloadSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedModelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

