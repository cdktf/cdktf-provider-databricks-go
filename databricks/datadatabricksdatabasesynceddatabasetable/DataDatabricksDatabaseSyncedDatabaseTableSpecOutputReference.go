// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdatabasesynceddatabasetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference interface {
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
	CreateDatabaseObjectsIfMissing() interface{}
	SetCreateDatabaseObjectsIfMissing(val interface{})
	CreateDatabaseObjectsIfMissingInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExistingPipelineId() *string
	SetExistingPipelineId(val *string)
	ExistingPipelineIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NewPipelineSpec() DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference
	NewPipelineSpecInput() interface{}
	PrimaryKeyColumns() *[]*string
	SetPrimaryKeyColumns(val *[]*string)
	PrimaryKeyColumnsInput() *[]*string
	SchedulingPolicy() *string
	SetSchedulingPolicy(val *string)
	SchedulingPolicyInput() *string
	SourceTableFullName() *string
	SetSourceTableFullName(val *string)
	SourceTableFullNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesKey() *string
	SetTimeseriesKey(val *string)
	TimeseriesKeyInput() *string
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
	PutNewPipelineSpec(value *DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec)
	ResetCreateDatabaseObjectsIfMissing()
	ResetExistingPipelineId()
	ResetNewPipelineSpec()
	ResetPrimaryKeyColumns()
	ResetSchedulingPolicy()
	ResetSourceTableFullName()
	ResetTimeseriesKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference
type jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) CreateDatabaseObjectsIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) CreateDatabaseObjectsIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ExistingPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ExistingPipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) NewPipelineSpec() DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference
	_jsii_.Get(
		j,
		"newPipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) NewPipelineSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newPipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) PrimaryKeyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) PrimaryKeyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) SchedulingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) SchedulingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) SourceTableFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) SourceTableFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) TimeseriesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) TimeseriesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKeyInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseSyncedDatabaseTableSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference_Override(d DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTable.DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetCreateDatabaseObjectsIfMissing(val interface{}) {
	if err := j.validateSetCreateDatabaseObjectsIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseObjectsIfMissing",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetExistingPipelineId(val *string) {
	if err := j.validateSetExistingPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingPipelineId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetPrimaryKeyColumns(val *[]*string) {
	if err := j.validateSetPrimaryKeyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetSchedulingPolicy(val *string) {
	if err := j.validateSetSchedulingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingPolicy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetSourceTableFullName(val *string) {
	if err := j.validateSetSourceTableFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTableFullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference)SetTimeseriesKey(val *string) {
	if err := j.validateSetTimeseriesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesKey",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) PutNewPipelineSpec(value *DataDatabricksDatabaseSyncedDatabaseTableSpecNewPipelineSpec) {
	if err := d.validatePutNewPipelineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewPipelineSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetCreateDatabaseObjectsIfMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateDatabaseObjectsIfMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetExistingPipelineId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingPipelineId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetNewPipelineSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetNewPipelineSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetPrimaryKeyColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryKeyColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetSourceTableFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTableFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ResetTimeseriesKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTableSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

