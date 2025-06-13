// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedEntitiesOutputReference interface {
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
	EntityName() *string
	SetEntityName(val *string)
	EntityNameInput() *string
	EntityVersion() *string
	SetEntityVersion(val *string)
	EntityVersionInput() *string
	EnvironmentVars() *map[string]*string
	SetEnvironmentVars(val *map[string]*string)
	EnvironmentVarsInput() *map[string]*string
	ExternalModel() ModelServingConfigServedEntitiesExternalModelOutputReference
	ExternalModelInput() *ModelServingConfigServedEntitiesExternalModel
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
	PutExternalModel(value *ModelServingConfigServedEntitiesExternalModel)
	ResetEntityName()
	ResetEntityVersion()
	ResetEnvironmentVars()
	ResetExternalModel()
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

// The jsii proxy struct for ModelServingConfigServedEntitiesOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EntityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EntityVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EntityVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EnvironmentVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) EnvironmentVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ExternalModel() ModelServingConfigServedEntitiesExternalModelOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelOutputReference
	_jsii_.Get(
		j,
		"externalModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ExternalModelInput() *ModelServingConfigServedEntitiesExternalModel {
	var returns *ModelServingConfigServedEntitiesExternalModel
	_jsii_.Get(
		j,
		"externalModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MaxProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MaxProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MaxProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MaxProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MinProvisionedConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MinProvisionedConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MinProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) MinProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ProvisionedModelUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedModelUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ProvisionedModelUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedModelUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ScaleToZeroEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleToZeroEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ScaleToZeroEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleToZeroEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) WorkloadSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) WorkloadSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) WorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) WorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ModelServingConfigServedEntitiesOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesOutputReference_Override(m ModelServingConfigServedEntitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetEntityName(val *string) {
	if err := j.validateSetEntityNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityName",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetEntityVersion(val *string) {
	if err := j.validateSetEntityVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityVersion",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetEnvironmentVars(val *map[string]*string) {
	if err := j.validateSetEnvironmentVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVars",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetMaxProvisionedConcurrency(val *float64) {
	if err := j.validateSetMaxProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProvisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetMaxProvisionedThroughput(val *float64) {
	if err := j.validateSetMaxProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetMinProvisionedConcurrency(val *float64) {
	if err := j.validateSetMinProvisionedConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProvisionedConcurrency",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetMinProvisionedThroughput(val *float64) {
	if err := j.validateSetMinProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetProvisionedModelUnits(val *float64) {
	if err := j.validateSetProvisionedModelUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedModelUnits",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetScaleToZeroEnabled(val interface{}) {
	if err := j.validateSetScaleToZeroEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleToZeroEnabled",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetWorkloadSize(val *string) {
	if err := j.validateSetWorkloadSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadSize",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesOutputReference)SetWorkloadType(val *string) {
	if err := j.validateSetWorkloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadType",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) PutExternalModel(value *ModelServingConfigServedEntitiesExternalModel) {
	if err := m.validatePutExternalModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putExternalModel",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetEntityName() {
	_jsii_.InvokeVoid(
		m,
		"resetEntityName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetEntityVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetEntityVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetEnvironmentVars() {
	_jsii_.InvokeVoid(
		m,
		"resetEnvironmentVars",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetExternalModel() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalModel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		m,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetMaxProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxProvisionedConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetMaxProvisionedThroughput() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxProvisionedThroughput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetMinProvisionedConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinProvisionedConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetMinProvisionedThroughput() {
	_jsii_.InvokeVoid(
		m,
		"resetMinProvisionedThroughput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetProvisionedModelUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetProvisionedModelUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetScaleToZeroEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetScaleToZeroEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetWorkloadSize() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkloadSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelServingConfigServedEntitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

