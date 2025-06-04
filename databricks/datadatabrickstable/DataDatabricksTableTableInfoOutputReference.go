// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickstable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksTableTableInfoOutputReference interface {
	cdktf.ComplexObject
	AccessPoint() *string
	SetAccessPoint(val *string)
	AccessPointInput() *string
	BrowseOnly() interface{}
	SetBrowseOnly(val interface{})
	BrowseOnlyInput() interface{}
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
	Columns() DataDatabricksTableTableInfoColumnsList
	ColumnsInput() interface{}
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
	DataAccessConfigurationId() *string
	SetDataAccessConfigurationId(val *string)
	DataAccessConfigurationIdInput() *string
	DataSourceFormat() *string
	SetDataSourceFormat(val *string)
	DataSourceFormatInput() *string
	DeletedAt() *float64
	SetDeletedAt(val *float64)
	DeletedAtInput() *float64
	DeltaRuntimePropertiesKvpairs() DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairsOutputReference
	DeltaRuntimePropertiesKvpairsInput() *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs
	EffectivePredictiveOptimizationFlag() DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlagOutputReference
	EffectivePredictiveOptimizationFlagInput() *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag
	EnablePredictiveOptimization() *string
	SetEnablePredictiveOptimization(val *string)
	EnablePredictiveOptimizationInput() *string
	EncryptionDetails() DataDatabricksTableTableInfoEncryptionDetailsOutputReference
	EncryptionDetailsInput() *DataDatabricksTableTableInfoEncryptionDetails
	// Experimental.
	Fqn() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	InternalValue() *DataDatabricksTableTableInfo
	SetInternalValue(val *DataDatabricksTableTableInfo)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	PipelineId() *string
	SetPipelineId(val *string)
	PipelineIdInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	RowFilter() DataDatabricksTableTableInfoRowFilterOutputReference
	RowFilterInput() *DataDatabricksTableTableInfoRowFilter
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	SqlPath() *string
	SetSqlPath(val *string)
	SqlPathInput() *string
	StorageCredentialName() *string
	SetStorageCredentialName(val *string)
	StorageCredentialNameInput() *string
	StorageLocation() *string
	SetStorageLocation(val *string)
	StorageLocationInput() *string
	TableConstraints() DataDatabricksTableTableInfoTableConstraintsList
	TableConstraintsInput() interface{}
	TableId() *string
	SetTableId(val *string)
	TableIdInput() *string
	TableType() *string
	SetTableType(val *string)
	TableTypeInput() *string
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
	ViewDefinition() *string
	SetViewDefinition(val *string)
	ViewDefinitionInput() *string
	ViewDependencies() DataDatabricksTableTableInfoViewDependenciesOutputReference
	ViewDependenciesInput() *DataDatabricksTableTableInfoViewDependencies
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
	PutColumns(value interface{})
	PutDeltaRuntimePropertiesKvpairs(value *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs)
	PutEffectivePredictiveOptimizationFlag(value *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag)
	PutEncryptionDetails(value *DataDatabricksTableTableInfoEncryptionDetails)
	PutRowFilter(value *DataDatabricksTableTableInfoRowFilter)
	PutTableConstraints(value interface{})
	PutViewDependencies(value *DataDatabricksTableTableInfoViewDependencies)
	ResetAccessPoint()
	ResetBrowseOnly()
	ResetCatalogName()
	ResetColumns()
	ResetComment()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetDataAccessConfigurationId()
	ResetDataSourceFormat()
	ResetDeletedAt()
	ResetDeltaRuntimePropertiesKvpairs()
	ResetEffectivePredictiveOptimizationFlag()
	ResetEnablePredictiveOptimization()
	ResetEncryptionDetails()
	ResetFullName()
	ResetMetastoreId()
	ResetName()
	ResetOwner()
	ResetPipelineId()
	ResetProperties()
	ResetRowFilter()
	ResetSchemaName()
	ResetSqlPath()
	ResetStorageCredentialName()
	ResetStorageLocation()
	ResetTableConstraints()
	ResetTableId()
	ResetTableType()
	ResetUpdatedAt()
	ResetUpdatedBy()
	ResetViewDefinition()
	ResetViewDependencies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksTableTableInfoOutputReference
type jsiiProxy_DataDatabricksTableTableInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) AccessPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) AccessPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) BrowseOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) BrowseOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Columns() DataDatabricksTableTableInfoColumnsList {
	var returns DataDatabricksTableTableInfoColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DataAccessConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DataAccessConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DataSourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DataSourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DeletedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deletedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DeletedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deletedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DeltaRuntimePropertiesKvpairs() DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairsOutputReference {
	var returns DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairsOutputReference
	_jsii_.Get(
		j,
		"deltaRuntimePropertiesKvpairs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) DeltaRuntimePropertiesKvpairsInput() *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs {
	var returns *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs
	_jsii_.Get(
		j,
		"deltaRuntimePropertiesKvpairsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EffectivePredictiveOptimizationFlag() DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlagOutputReference {
	var returns DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlagOutputReference
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EffectivePredictiveOptimizationFlagInput() *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag {
	var returns *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EnablePredictiveOptimization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EnablePredictiveOptimizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EncryptionDetails() DataDatabricksTableTableInfoEncryptionDetailsOutputReference {
	var returns DataDatabricksTableTableInfoEncryptionDetailsOutputReference
	_jsii_.Get(
		j,
		"encryptionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) EncryptionDetailsInput() *DataDatabricksTableTableInfoEncryptionDetails {
	var returns *DataDatabricksTableTableInfoEncryptionDetails
	_jsii_.Get(
		j,
		"encryptionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) InternalValue() *DataDatabricksTableTableInfo {
	var returns *DataDatabricksTableTableInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PipelineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PipelineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) RowFilter() DataDatabricksTableTableInfoRowFilterOutputReference {
	var returns DataDatabricksTableTableInfoRowFilterOutputReference
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) RowFilterInput() *DataDatabricksTableTableInfoRowFilter {
	var returns *DataDatabricksTableTableInfoRowFilter
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) SqlPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) SqlPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) StorageCredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCredentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) StorageCredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCredentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) StorageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) StorageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableConstraints() DataDatabricksTableTableInfoTableConstraintsList {
	var returns DataDatabricksTableTableInfoTableConstraintsList
	_jsii_.Get(
		j,
		"tableConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ViewDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ViewDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ViewDependencies() DataDatabricksTableTableInfoViewDependenciesOutputReference {
	var returns DataDatabricksTableTableInfoViewDependenciesOutputReference
	_jsii_.Get(
		j,
		"viewDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ViewDependenciesInput() *DataDatabricksTableTableInfoViewDependencies {
	var returns *DataDatabricksTableTableInfoViewDependencies
	_jsii_.Get(
		j,
		"viewDependenciesInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksTableTableInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksTableTableInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksTableTableInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksTableTableInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksTableTableInfoOutputReference_Override(d DataDatabricksTableTableInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksTable.DataDatabricksTableTableInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetAccessPoint(val *string) {
	if err := j.validateSetAccessPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPoint",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetBrowseOnly(val interface{}) {
	if err := j.validateSetBrowseOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browseOnly",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetDataAccessConfigurationId(val *string) {
	if err := j.validateSetDataAccessConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessConfigurationId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetDataSourceFormat(val *string) {
	if err := j.validateSetDataSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceFormat",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetDeletedAt(val *float64) {
	if err := j.validateSetDeletedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetEnablePredictiveOptimization(val *string) {
	if err := j.validateSetEnablePredictiveOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePredictiveOptimization",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetInternalValue(val *DataDatabricksTableTableInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetPipelineId(val *string) {
	if err := j.validateSetPipelineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetSqlPath(val *string) {
	if err := j.validateSetSqlPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetStorageCredentialName(val *string) {
	if err := j.validateSetStorageCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCredentialName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetStorageLocation(val *string) {
	if err := j.validateSetStorageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetTableId(val *string) {
	if err := j.validateSetTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetTableType(val *string) {
	if err := j.validateSetTableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksTableTableInfoOutputReference)SetViewDefinition(val *string) {
	if err := j.validateSetViewDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewDefinition",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutColumns(value interface{}) {
	if err := d.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumns",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutDeltaRuntimePropertiesKvpairs(value *DataDatabricksTableTableInfoDeltaRuntimePropertiesKvpairs) {
	if err := d.validatePutDeltaRuntimePropertiesKvpairsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeltaRuntimePropertiesKvpairs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutEffectivePredictiveOptimizationFlag(value *DataDatabricksTableTableInfoEffectivePredictiveOptimizationFlag) {
	if err := d.validatePutEffectivePredictiveOptimizationFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectivePredictiveOptimizationFlag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutEncryptionDetails(value *DataDatabricksTableTableInfoEncryptionDetails) {
	if err := d.validatePutEncryptionDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEncryptionDetails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutRowFilter(value *DataDatabricksTableTableInfoRowFilter) {
	if err := d.validatePutRowFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRowFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutTableConstraints(value interface{}) {
	if err := d.validatePutTableConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableConstraints",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) PutViewDependencies(value *DataDatabricksTableTableInfoViewDependencies) {
	if err := d.validatePutViewDependenciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putViewDependencies",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetAccessPoint() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessPoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetBrowseOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowseOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetCatalogName() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetDataAccessConfigurationId() {
	_jsii_.InvokeVoid(
		d,
		"resetDataAccessConfigurationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetDataSourceFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSourceFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetDeletedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetDeltaRuntimePropertiesKvpairs() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaRuntimePropertiesKvpairs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetEffectivePredictiveOptimizationFlag() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectivePredictiveOptimizationFlag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetEnablePredictiveOptimization() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePredictiveOptimization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetEncryptionDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetPipelineId() {
	_jsii_.InvokeVoid(
		d,
		"resetPipelineId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetRowFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetSqlPath() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetStorageCredentialName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageCredentialName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetStorageLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetTableConstraints() {
	_jsii_.InvokeVoid(
		d,
		"resetTableConstraints",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetTableId() {
	_jsii_.InvokeVoid(
		d,
		"resetTableId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetTableType() {
	_jsii_.InvokeVoid(
		d,
		"resetTableType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetViewDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetViewDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ResetViewDependencies() {
	_jsii_.InvokeVoid(
		d,
		"resetViewDependencies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksTableTableInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

