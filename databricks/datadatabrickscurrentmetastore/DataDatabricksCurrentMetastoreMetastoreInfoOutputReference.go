// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscurrentmetastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabrickscurrentmetastore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksCurrentMetastoreMetastoreInfoOutputReference interface {
	cdktf.ComplexObject
	Cloud() *string
	SetCloud(val *string)
	CloudInput() *string
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
	DefaultDataAccessConfigId() *string
	SetDefaultDataAccessConfigId(val *string)
	DefaultDataAccessConfigIdInput() *string
	DeltaSharingOrganizationName() *string
	SetDeltaSharingOrganizationName(val *string)
	DeltaSharingOrganizationNameInput() *string
	DeltaSharingRecipientTokenLifetimeInSeconds() *float64
	SetDeltaSharingRecipientTokenLifetimeInSeconds(val *float64)
	DeltaSharingRecipientTokenLifetimeInSecondsInput() *float64
	DeltaSharingScope() *string
	SetDeltaSharingScope(val *string)
	DeltaSharingScopeInput() *string
	ExternalAccessEnabled() interface{}
	SetExternalAccessEnabled(val interface{})
	ExternalAccessEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	GlobalMetastoreId() *string
	SetGlobalMetastoreId(val *string)
	GlobalMetastoreIdInput() *string
	InternalValue() *DataDatabricksCurrentMetastoreMetastoreInfo
	SetInternalValue(val *DataDatabricksCurrentMetastoreMetastoreInfo)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	PrivilegeModelVersion() *string
	SetPrivilegeModelVersion(val *string)
	PrivilegeModelVersionInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	StorageRoot() *string
	SetStorageRoot(val *string)
	StorageRootCredentialId() *string
	SetStorageRootCredentialId(val *string)
	StorageRootCredentialIdInput() *string
	StorageRootCredentialName() *string
	SetStorageRootCredentialName(val *string)
	StorageRootCredentialNameInput() *string
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
	ResetCloud()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetDefaultDataAccessConfigId()
	ResetDeltaSharingOrganizationName()
	ResetDeltaSharingRecipientTokenLifetimeInSeconds()
	ResetDeltaSharingScope()
	ResetExternalAccessEnabled()
	ResetGlobalMetastoreId()
	ResetMetastoreId()
	ResetName()
	ResetOwner()
	ResetPrivilegeModelVersion()
	ResetRegion()
	ResetStorageRoot()
	ResetStorageRootCredentialId()
	ResetStorageRootCredentialName()
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

// The jsii proxy struct for DataDatabricksCurrentMetastoreMetastoreInfoOutputReference
type jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CloudInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DefaultDataAccessConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DefaultDataAccessConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingOrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingOrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingRecipientTokenLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingRecipientTokenLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) DeltaSharingScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ExternalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ExternalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GlobalMetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GlobalMetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) InternalValue() *DataDatabricksCurrentMetastoreMetastoreInfo {
	var returns *DataDatabricksCurrentMetastoreMetastoreInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) PrivilegeModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) PrivilegeModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRootCredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRootCredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRootCredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRootCredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksCurrentMetastoreMetastoreInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksCurrentMetastoreMetastoreInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksCurrentMetastoreMetastoreInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCurrentMetastore.DataDatabricksCurrentMetastoreMetastoreInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksCurrentMetastoreMetastoreInfoOutputReference_Override(d DataDatabricksCurrentMetastoreMetastoreInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCurrentMetastore.DataDatabricksCurrentMetastoreMetastoreInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetCloud(val *string) {
	if err := j.validateSetCloudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloud",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetDefaultDataAccessConfigId(val *string) {
	if err := j.validateSetDefaultDataAccessConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDataAccessConfigId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetDeltaSharingOrganizationName(val *string) {
	if err := j.validateSetDeltaSharingOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingOrganizationName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetDeltaSharingRecipientTokenLifetimeInSeconds(val *float64) {
	if err := j.validateSetDeltaSharingRecipientTokenLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetDeltaSharingScope(val *string) {
	if err := j.validateSetDeltaSharingScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingScope",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetExternalAccessEnabled(val interface{}) {
	if err := j.validateSetExternalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetGlobalMetastoreId(val *string) {
	if err := j.validateSetGlobalMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalMetastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetInternalValue(val *DataDatabricksCurrentMetastoreMetastoreInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetPrivilegeModelVersion(val *string) {
	if err := j.validateSetPrivilegeModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegeModelVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetStorageRootCredentialId(val *string) {
	if err := j.validateSetStorageRootCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetStorageRootCredentialName(val *string) {
	if err := j.validateSetStorageRootCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetCloud() {
	_jsii_.InvokeVoid(
		d,
		"resetCloud",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetDefaultDataAccessConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDataAccessConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetDeltaSharingOrganizationName() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingOrganizationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetDeltaSharingRecipientTokenLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingRecipientTokenLifetimeInSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetDeltaSharingScope() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetExternalAccessEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalAccessEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetGlobalMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetPrivilegeModelVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivilegeModelVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetStorageRootCredentialId() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRootCredentialId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetStorageRootCredentialName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRootCredentialName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksCurrentMetastoreMetastoreInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

