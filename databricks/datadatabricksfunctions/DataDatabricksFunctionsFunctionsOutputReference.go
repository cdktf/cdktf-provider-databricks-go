// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksfunctions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksFunctionsFunctionsOutputReference interface {
	cdktf.ComplexObject
	BrowseOnly() interface{}
	SetBrowseOnly(val interface{})
	BrowseOnlyInput() interface{}
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
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	ExternalLanguage() *string
	SetExternalLanguage(val *string)
	ExternalLanguageInput() *string
	ExternalName() *string
	SetExternalName(val *string)
	ExternalNameInput() *string
	// Experimental.
	Fqn() *string
	FullDataType() *string
	SetFullDataType(val *string)
	FullDataTypeInput() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	FunctionId() *string
	SetFunctionId(val *string)
	FunctionIdInput() *string
	InputParams() DataDatabricksFunctionsFunctionsInputParamsOutputReference
	InputParamsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsDeterministic() interface{}
	SetIsDeterministic(val interface{})
	IsDeterministicInput() interface{}
	IsNullCall() interface{}
	SetIsNullCall(val interface{})
	IsNullCallInput() interface{}
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	ParameterStyle() *string
	SetParameterStyle(val *string)
	ParameterStyleInput() *string
	Properties() *string
	SetProperties(val *string)
	PropertiesInput() *string
	ReturnParams() DataDatabricksFunctionsFunctionsReturnParamsOutputReference
	ReturnParamsInput() interface{}
	RoutineBody() *string
	SetRoutineBody(val *string)
	RoutineBodyInput() *string
	RoutineDefinition() *string
	SetRoutineDefinition(val *string)
	RoutineDefinitionInput() *string
	RoutineDependencies() DataDatabricksFunctionsFunctionsRoutineDependenciesOutputReference
	RoutineDependenciesInput() interface{}
	SchemaName() *string
	SetSchemaName(val *string)
	SchemaNameInput() *string
	SecurityType() *string
	SetSecurityType(val *string)
	SecurityTypeInput() *string
	SpecificName() *string
	SetSpecificName(val *string)
	SpecificNameInput() *string
	SqlDataAccess() *string
	SetSqlDataAccess(val *string)
	SqlDataAccessInput() *string
	SqlPath() *string
	SetSqlPath(val *string)
	SqlPathInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutInputParams(value *DataDatabricksFunctionsFunctionsInputParams)
	PutReturnParams(value *DataDatabricksFunctionsFunctionsReturnParams)
	PutRoutineDependencies(value *DataDatabricksFunctionsFunctionsRoutineDependencies)
	ResetBrowseOnly()
	ResetCatalogName()
	ResetComment()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetDataType()
	ResetExternalLanguage()
	ResetExternalName()
	ResetFullDataType()
	ResetFullName()
	ResetFunctionId()
	ResetInputParams()
	ResetIsDeterministic()
	ResetIsNullCall()
	ResetMetastoreId()
	ResetName()
	ResetOwner()
	ResetParameterStyle()
	ResetProperties()
	ResetReturnParams()
	ResetRoutineBody()
	ResetRoutineDefinition()
	ResetRoutineDependencies()
	ResetSchemaName()
	ResetSecurityType()
	ResetSpecificName()
	ResetSqlDataAccess()
	ResetSqlPath()
	ResetUpdatedAt()
	ResetUpdatedBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksFunctionsFunctionsOutputReference
type jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) BrowseOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) BrowseOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ExternalLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ExternalLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FullDataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullDataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FullDataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullDataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) FunctionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) InputParams() DataDatabricksFunctionsFunctionsInputParamsOutputReference {
	var returns DataDatabricksFunctionsFunctionsInputParamsOutputReference
	_jsii_.Get(
		j,
		"inputParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) InputParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) IsDeterministic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeterministic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) IsDeterministicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeterministicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) IsNullCall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) IsNullCallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ParameterStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ParameterStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) PropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ReturnParams() DataDatabricksFunctionsFunctionsReturnParamsOutputReference {
	var returns DataDatabricksFunctionsFunctionsReturnParamsOutputReference
	_jsii_.Get(
		j,
		"returnParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ReturnParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineDependencies() DataDatabricksFunctionsFunctionsRoutineDependenciesOutputReference {
	var returns DataDatabricksFunctionsFunctionsRoutineDependenciesOutputReference
	_jsii_.Get(
		j,
		"routineDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) RoutineDependenciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routineDependenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SchemaNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SecurityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SecurityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SpecificName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specificName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SpecificNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specificNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SqlDataAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SqlDataAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlDataAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SqlPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) SqlPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksFunctionsFunctionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksFunctionsFunctionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksFunctionsFunctionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksFunctions.DataDatabricksFunctionsFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksFunctionsFunctionsOutputReference_Override(d DataDatabricksFunctionsFunctionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksFunctions.DataDatabricksFunctionsFunctionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetBrowseOnly(val interface{}) {
	if err := j.validateSetBrowseOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browseOnly",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetDataType(val *string) {
	if err := j.validateSetDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetExternalLanguage(val *string) {
	if err := j.validateSetExternalLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalLanguage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetFullDataType(val *string) {
	if err := j.validateSetFullDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullDataType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetFunctionId(val *string) {
	if err := j.validateSetFunctionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetIsDeterministic(val interface{}) {
	if err := j.validateSetIsDeterministicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDeterministic",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetIsNullCall(val interface{}) {
	if err := j.validateSetIsNullCallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNullCall",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetParameterStyle(val *string) {
	if err := j.validateSetParameterStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterStyle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetProperties(val *string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetRoutineBody(val *string) {
	if err := j.validateSetRoutineBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineBody",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetRoutineDefinition(val *string) {
	if err := j.validateSetRoutineDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routineDefinition",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetSchemaName(val *string) {
	if err := j.validateSetSchemaNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetSecurityType(val *string) {
	if err := j.validateSetSecurityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetSpecificName(val *string) {
	if err := j.validateSetSpecificNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specificName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetSqlDataAccess(val *string) {
	if err := j.validateSetSqlDataAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlDataAccess",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetSqlPath(val *string) {
	if err := j.validateSetSqlPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) PutInputParams(value *DataDatabricksFunctionsFunctionsInputParams) {
	if err := d.validatePutInputParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInputParams",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) PutReturnParams(value *DataDatabricksFunctionsFunctionsReturnParams) {
	if err := d.validatePutReturnParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReturnParams",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) PutRoutineDependencies(value *DataDatabricksFunctionsFunctionsRoutineDependencies) {
	if err := d.validatePutRoutineDependenciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRoutineDependencies",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetBrowseOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowseOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetCatalogName() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetDataType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetExternalLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalLanguage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetExternalName() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetFullDataType() {
	_jsii_.InvokeVoid(
		d,
		"resetFullDataType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetFunctionId() {
	_jsii_.InvokeVoid(
		d,
		"resetFunctionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetInputParams() {
	_jsii_.InvokeVoid(
		d,
		"resetInputParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetIsDeterministic() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDeterministic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetIsNullCall() {
	_jsii_.InvokeVoid(
		d,
		"resetIsNullCall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetParameterStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetReturnParams() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetRoutineBody() {
	_jsii_.InvokeVoid(
		d,
		"resetRoutineBody",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetRoutineDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRoutineDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetRoutineDependencies() {
	_jsii_.InvokeVoid(
		d,
		"resetRoutineDependencies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetSchemaName() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetSecurityType() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetSpecificName() {
	_jsii_.InvokeVoid(
		d,
		"resetSpecificName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetSqlDataAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlDataAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetSqlPath() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksFunctionsFunctionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

