// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksdatabasesynceddatabasetables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksdatabasesynceddatabasetables/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference interface {
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
	DatabaseInstanceName() *string
	DataSynchronizationStatus() DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference
	EffectiveDatabaseInstanceName() *string
	EffectiveLogicalDatabaseName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables
	SetInternalValue(val *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables)
	LogicalDatabaseName() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Spec() DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnityCatalogProvisioningState() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference
type jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) DatabaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) DataSynchronizationStatus() DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesDataSynchronizationStatusOutputReference
	_jsii_.Get(
		j,
		"dataSynchronizationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) EffectiveDatabaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDatabaseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) EffectiveLogicalDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveLogicalDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) InternalValue() *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables {
	var returns *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) LogicalDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) Spec() DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference {
	var returns DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesSpecOutputReference
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) UnityCatalogProvisioningState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unityCatalogProvisioningState",
		&returns,
	)
	return returns
}


func NewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference_Override(d DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksDatabaseSyncedDatabaseTables.DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetInternalValue(val *DataDatabricksDatabaseSyncedDatabaseTablesSyncedTables) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksDatabaseSyncedDatabaseTablesSyncedTablesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

