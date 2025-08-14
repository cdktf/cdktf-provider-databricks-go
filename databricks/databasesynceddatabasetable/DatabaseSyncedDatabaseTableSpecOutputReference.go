// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesynceddatabasetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/databasesynceddatabasetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSyncedDatabaseTableSpecOutputReference interface {
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
	NewPipelineSpec() DatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference
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
	PutNewPipelineSpec(value *DatabaseSyncedDatabaseTableSpecNewPipelineSpec)
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

// The jsii proxy struct for DatabaseSyncedDatabaseTableSpecOutputReference
type jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) CreateDatabaseObjectsIfMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) CreateDatabaseObjectsIfMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseObjectsIfMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ExistingPipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ExistingPipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingPipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) NewPipelineSpec() DatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference {
	var returns DatabaseSyncedDatabaseTableSpecNewPipelineSpecOutputReference
	_jsii_.Get(
		j,
		"newPipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) NewPipelineSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newPipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) PrimaryKeyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) PrimaryKeyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"primaryKeyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) SchedulingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) SchedulingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) SourceTableFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) SourceTableFullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTableFullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) TimeseriesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) TimeseriesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeseriesKeyInput",
		&returns,
	)
	return returns
}


func NewDatabaseSyncedDatabaseTableSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseSyncedDatabaseTableSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSyncedDatabaseTableSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseSyncedDatabaseTable.DatabaseSyncedDatabaseTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSyncedDatabaseTableSpecOutputReference_Override(d DatabaseSyncedDatabaseTableSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.databaseSyncedDatabaseTable.DatabaseSyncedDatabaseTableSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetCreateDatabaseObjectsIfMissing(val interface{}) {
	if err := j.validateSetCreateDatabaseObjectsIfMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseObjectsIfMissing",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetExistingPipelineId(val *string) {
	if err := j.validateSetExistingPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingPipelineId",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetPrimaryKeyColumns(val *[]*string) {
	if err := j.validateSetPrimaryKeyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryKeyColumns",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetSchedulingPolicy(val *string) {
	if err := j.validateSetSchedulingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetSourceTableFullName(val *string) {
	if err := j.validateSetSourceTableFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTableFullName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference)SetTimeseriesKey(val *string) {
	if err := j.validateSetTimeseriesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeseriesKey",
		val,
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) PutNewPipelineSpec(value *DatabaseSyncedDatabaseTableSpecNewPipelineSpec) {
	if err := d.validatePutNewPipelineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewPipelineSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetCreateDatabaseObjectsIfMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateDatabaseObjectsIfMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetExistingPipelineId() {
	_jsii_.InvokeVoid(
		d,
		"resetExistingPipelineId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetNewPipelineSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetNewPipelineSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetPrimaryKeyColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryKeyColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetSchedulingPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedulingPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetSourceTableFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceTableFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ResetTimeseriesKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSyncedDatabaseTableSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

