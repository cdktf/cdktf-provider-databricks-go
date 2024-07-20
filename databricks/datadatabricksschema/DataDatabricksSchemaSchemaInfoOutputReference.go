// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksschema/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksSchemaSchemaInfoOutputReference interface {
	cdktf.ComplexObject
	BrowseOnly() interface{}
	SetBrowseOnly(val interface{})
	BrowseOnlyInput() interface{}
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
	CatalogType() *string
	SetCatalogType(val *string)
	CatalogTypeInput() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	CreatedAt() *float64
	SetCreatedAt(val *float64)
	CreatedAtInput() *float64
	CreatedBy() *string
	SetCreatedBy(val *string)
	CreatedByInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectivePredictiveOptimizationFlag() DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference
	EffectivePredictiveOptimizationFlagInput() *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag
	EnablePredictiveOptimization() *string
	SetEnablePredictiveOptimization(val *string)
	EnablePredictiveOptimizationInput() *string
	// Experimental.
	Fqn() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	InternalValue() *DataDatabricksSchemaSchemaInfo
	SetInternalValue(val *DataDatabricksSchemaSchemaInfo)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	SchemaId() *string
	SetSchemaId(val *string)
	SchemaIdInput() *string
	StorageLocation() *string
	SetStorageLocation(val *string)
	StorageLocationInput() *string
	StorageRoot() *string
	SetStorageRoot(val *string)
	StorageRootInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdatedAt() *float64
	SetUpdatedAt(val *float64)
	UpdatedAtInput() *float64
	UpdatedBy() *string
	SetUpdatedBy(val *string)
	UpdatedByInput() *string
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
	PutEffectivePredictiveOptimizationFlag(value *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag)
	ResetBrowseOnly()
	ResetCatalogName()
	ResetCatalogType()
	ResetComment()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetEffectivePredictiveOptimizationFlag()
	ResetEnablePredictiveOptimization()
	ResetFullName()
	ResetMetastoreId()
	ResetName()
	ResetOwner()
	ResetProperties()
	ResetSchemaId()
	ResetStorageLocation()
	ResetStorageRoot()
	ResetUpdatedAt()
	ResetUpdatedBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksSchemaSchemaInfoOutputReference
type jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) BrowseOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) BrowseOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CatalogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CatalogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) EffectivePredictiveOptimizationFlag() DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference {
	var returns DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlagOutputReference
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) EffectivePredictiveOptimizationFlagInput() *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag {
	var returns *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) EnablePredictiveOptimization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) EnablePredictiveOptimizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) InternalValue() *DataDatabricksSchemaSchemaInfo {
	var returns *DataDatabricksSchemaSchemaInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) SchemaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) SchemaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) StorageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) StorageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksSchemaSchemaInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksSchemaSchemaInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksSchemaSchemaInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSchema.DataDatabricksSchemaSchemaInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksSchemaSchemaInfoOutputReference_Override(d DataDatabricksSchemaSchemaInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksSchema.DataDatabricksSchemaSchemaInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetBrowseOnly(val interface{}) {
	if err := j.validateSetBrowseOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browseOnly",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetCatalogType(val *string) {
	if err := j.validateSetCatalogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetEnablePredictiveOptimization(val *string) {
	if err := j.validateSetEnablePredictiveOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePredictiveOptimization",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetInternalValue(val *DataDatabricksSchemaSchemaInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetSchemaId(val *string) {
	if err := j.validateSetSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetStorageLocation(val *string) {
	if err := j.validateSetStorageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) PutEffectivePredictiveOptimizationFlag(value *DataDatabricksSchemaSchemaInfoEffectivePredictiveOptimizationFlag) {
	if err := d.validatePutEffectivePredictiveOptimizationFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectivePredictiveOptimizationFlag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetBrowseOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowseOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetCatalogName() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetCatalogType() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetEffectivePredictiveOptimizationFlag() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectivePredictiveOptimizationFlag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetEnablePredictiveOptimization() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePredictiveOptimization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetStorageLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksSchemaSchemaInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

