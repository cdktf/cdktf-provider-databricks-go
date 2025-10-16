// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksregisteredmodelversions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksregisteredmodelversions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksRegisteredModelVersionsModelVersionsOutputReference interface {
	cdktf.ComplexObject
	Aliases() DataDatabricksRegisteredModelVersionsModelVersionsAliasesList
	AliasesInput() interface{}
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
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
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
	ModelVersionDependencies() DataDatabricksRegisteredModelVersionsModelVersionsModelVersionDependenciesList
	ModelVersionDependenciesInput() interface{}
	RunId() *string
	SetRunId(val *string)
	RunIdInput() *string
	RunWorkspaceId() *float64
	SetRunWorkspaceId(val *float64)
	RunWorkspaceIdInput() *float64
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StorageLocation() *string
	SetStorageLocation(val *string)
	StorageLocationInput() *string
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
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
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
	PutAliases(value interface{})
	PutModelVersionDependencies(value interface{})
	ResetAliases()
	ResetCatalogName()
	ResetComment()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetId()
	ResetMetastoreId()
	ResetModelName()
	ResetModelVersionDependencies()
	ResetRunId()
	ResetRunWorkspaceId()
	ResetSchemaName()
	ResetSource()
	ResetStatus()
	ResetStorageLocation()
	ResetUpdatedAt()
	ResetUpdatedBy()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksRegisteredModelVersionsModelVersionsOutputReference
type jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Aliases() DataDatabricksRegisteredModelVersionsModelVersionsAliasesList {
	var returns DataDatabricksRegisteredModelVersionsModelVersionsAliasesList
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) AliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ModelVersionDependencies() DataDatabricksRegisteredModelVersionsModelVersionsModelVersionDependenciesList {
	var returns DataDatabricksRegisteredModelVersionsModelVersionsModelVersionDependenciesList
	_jsii_.Get(
		j,
		"modelVersionDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ModelVersionDependenciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelVersionDependenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) RunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) RunIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) RunWorkspaceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) RunWorkspaceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) StorageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) StorageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksRegisteredModelVersionsModelVersionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksRegisteredModelVersionsModelVersionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksRegisteredModelVersionsModelVersionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksRegisteredModelVersions.DataDatabricksRegisteredModelVersionsModelVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksRegisteredModelVersionsModelVersionsOutputReference_Override(d DataDatabricksRegisteredModelVersionsModelVersionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksRegisteredModelVersions.DataDatabricksRegisteredModelVersionsModelVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetRunId(val *string) {
	if err := j.validateSetRunIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetRunWorkspaceId(val *float64) {
	if err := j.validateSetRunWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetStorageLocation(val *string) {
	if err := j.validateSetStorageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference)SetVersion(val *float64) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) PutAliases(value interface{}) {
	if err := d.validatePutAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAliases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) PutModelVersionDependencies(value interface{}) {
	if err := d.validatePutModelVersionDependenciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putModelVersionDependencies",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetAliases() {
	_jsii_.InvokeVoid(
		d,
		"resetAliases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetCatalogName() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		d,
		"resetModelName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetModelVersionDependencies() {
	_jsii_.InvokeVoid(
		d,
		"resetModelVersionDependencies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetRunId() {
	_jsii_.InvokeVoid(
		d,
		"resetRunId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetRunWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetRunWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetStorageLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksRegisteredModelVersionsModelVersionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

