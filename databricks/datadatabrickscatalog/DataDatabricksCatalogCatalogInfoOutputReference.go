// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabrickscatalog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCatalogCatalogInfoOutputReference interface {
	cdktf.ComplexObject
	BrowseOnly() interface{}
	SetBrowseOnly(val interface{})
	BrowseOnlyInput() interface{}
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
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
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
	EffectivePredictiveOptimizationFlag() DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference
	EffectivePredictiveOptimizationFlagInput() *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag
	EnablePredictiveOptimization() *string
	SetEnablePredictiveOptimization(val *string)
	EnablePredictiveOptimizationInput() *string
	// Experimental.
	Fqn() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	InternalValue() *DataDatabricksCatalogCatalogInfo
	SetInternalValue(val *DataDatabricksCatalogCatalogInfo)
	IsolationMode() *string
	SetIsolationMode(val *string)
	IsolationModeInput() *string
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	ProvisioningInfo() DataDatabricksCatalogCatalogInfoProvisioningInfoOutputReference
	ProvisioningInfoInput() *DataDatabricksCatalogCatalogInfoProvisioningInfo
	SecurableKind() *string
	SetSecurableKind(val *string)
	SecurableKindInput() *string
	SecurableType() *string
	SetSecurableType(val *string)
	SecurableTypeInput() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
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
	PutEffectivePredictiveOptimizationFlag(value *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag)
	PutProvisioningInfo(value *DataDatabricksCatalogCatalogInfoProvisioningInfo)
	ResetBrowseOnly()
	ResetCatalogType()
	ResetComment()
	ResetConnectionName()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetEffectivePredictiveOptimizationFlag()
	ResetEnablePredictiveOptimization()
	ResetFullName()
	ResetIsolationMode()
	ResetMetastoreId()
	ResetName()
	ResetOptions()
	ResetOwner()
	ResetProperties()
	ResetProviderName()
	ResetProvisioningInfo()
	ResetSecurableKind()
	ResetSecurableType()
	ResetShareName()
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

// The jsii proxy struct for DataDatabricksCatalogCatalogInfoOutputReference
type jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) BrowseOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) BrowseOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browseOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CatalogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CatalogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) EffectivePredictiveOptimizationFlag() DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference {
	var returns DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlagOutputReference
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) EffectivePredictiveOptimizationFlagInput() *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag {
	var returns *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag
	_jsii_.Get(
		j,
		"effectivePredictiveOptimizationFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) EnablePredictiveOptimization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) EnablePredictiveOptimizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePredictiveOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) InternalValue() *DataDatabricksCatalogCatalogInfo {
	var returns *DataDatabricksCatalogCatalogInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) IsolationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) IsolationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ProvisioningInfo() DataDatabricksCatalogCatalogInfoProvisioningInfoOutputReference {
	var returns DataDatabricksCatalogCatalogInfoProvisioningInfoOutputReference
	_jsii_.Get(
		j,
		"provisioningInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ProvisioningInfoInput() *DataDatabricksCatalogCatalogInfoProvisioningInfo {
	var returns *DataDatabricksCatalogCatalogInfoProvisioningInfo
	_jsii_.Get(
		j,
		"provisioningInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) SecurableKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) SecurableKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) SecurableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) SecurableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) StorageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) StorageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCatalogCatalogInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCatalogCatalogInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCatalogCatalogInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCatalogCatalogInfoOutputReference_Override(d DataDatabricksCatalogCatalogInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCatalog.DataDatabricksCatalogCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetBrowseOnly(val interface{}) {
	if err := j.validateSetBrowseOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browseOnly",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetCatalogType(val *string) {
	if err := j.validateSetCatalogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetConnectionName(val *string) {
	if err := j.validateSetConnectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetEnablePredictiveOptimization(val *string) {
	if err := j.validateSetEnablePredictiveOptimizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePredictiveOptimization",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetInternalValue(val *DataDatabricksCatalogCatalogInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetIsolationMode(val *string) {
	if err := j.validateSetIsolationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolationMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetSecurableKind(val *string) {
	if err := j.validateSetSecurableKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableKind",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetSecurableType(val *string) {
	if err := j.validateSetSecurableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securableType",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetStorageLocation(val *string) {
	if err := j.validateSetStorageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocation",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) PutEffectivePredictiveOptimizationFlag(value *DataDatabricksCatalogCatalogInfoEffectivePredictiveOptimizationFlag) {
	if err := d.validatePutEffectivePredictiveOptimizationFlagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectivePredictiveOptimizationFlag",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) PutProvisioningInfo(value *DataDatabricksCatalogCatalogInfoProvisioningInfo) {
	if err := d.validatePutProvisioningInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisioningInfo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetBrowseOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetBrowseOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetCatalogType() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetConnectionName() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetEffectivePredictiveOptimizationFlag() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectivePredictiveOptimizationFlag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetEnablePredictiveOptimization() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePredictiveOptimization",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetFullName() {
	_jsii_.InvokeVoid(
		d,
		"resetFullName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetIsolationMode() {
	_jsii_.InvokeVoid(
		d,
		"resetIsolationMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetProviderName() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetProvisioningInfo() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisioningInfo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetSecurableKind() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurableKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetSecurableType() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurableType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetShareName() {
	_jsii_.InvokeVoid(
		d,
		"resetShareName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetStorageLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCatalogCatalogInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

