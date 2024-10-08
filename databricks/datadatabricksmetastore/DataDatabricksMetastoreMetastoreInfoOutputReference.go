// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmetastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksmetastore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksMetastoreMetastoreInfoOutputReference interface {
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
	InternalValue() *DataDatabricksMetastoreMetastoreInfo
	SetInternalValue(val *DataDatabricksMetastoreMetastoreInfo)
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

// The jsii proxy struct for DataDatabricksMetastoreMetastoreInfoOutputReference
type jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CloudInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DefaultDataAccessConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DefaultDataAccessConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDataAccessConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingOrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingOrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingOrganizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingRecipientTokenLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingRecipientTokenLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSharingRecipientTokenLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) DeltaSharingScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSharingScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ExternalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ExternalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GlobalMetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GlobalMetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalMetastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) InternalValue() *DataDatabricksMetastoreMetastoreInfo {
	var returns *DataDatabricksMetastoreMetastoreInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) PrivilegeModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) PrivilegeModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegeModelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRoot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRootCredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRootCredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRootCredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRootCredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootCredentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) StorageRootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksMetastoreMetastoreInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksMetastoreMetastoreInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksMetastoreMetastoreInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastoreMetastoreInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksMetastoreMetastoreInfoOutputReference_Override(d DataDatabricksMetastoreMetastoreInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMetastore.DataDatabricksMetastoreMetastoreInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetCloud(val *string) {
	if err := j.validateSetCloudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloud",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetDefaultDataAccessConfigId(val *string) {
	if err := j.validateSetDefaultDataAccessConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDataAccessConfigId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetDeltaSharingOrganizationName(val *string) {
	if err := j.validateSetDeltaSharingOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingOrganizationName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetDeltaSharingRecipientTokenLifetimeInSeconds(val *float64) {
	if err := j.validateSetDeltaSharingRecipientTokenLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingRecipientTokenLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetDeltaSharingScope(val *string) {
	if err := j.validateSetDeltaSharingScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSharingScope",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetExternalAccessEnabled(val interface{}) {
	if err := j.validateSetExternalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetGlobalMetastoreId(val *string) {
	if err := j.validateSetGlobalMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalMetastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetInternalValue(val *DataDatabricksMetastoreMetastoreInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetPrivilegeModelVersion(val *string) {
	if err := j.validateSetPrivilegeModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegeModelVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetStorageRoot(val *string) {
	if err := j.validateSetStorageRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRoot",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetStorageRootCredentialId(val *string) {
	if err := j.validateSetStorageRootCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetStorageRootCredentialName(val *string) {
	if err := j.validateSetStorageRootCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageRootCredentialName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetCloud() {
	_jsii_.InvokeVoid(
		d,
		"resetCloud",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetDefaultDataAccessConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDataAccessConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetDeltaSharingOrganizationName() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingOrganizationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetDeltaSharingRecipientTokenLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingRecipientTokenLifetimeInSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetDeltaSharingScope() {
	_jsii_.InvokeVoid(
		d,
		"resetDeltaSharingScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetExternalAccessEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalAccessEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetGlobalMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetPrivilegeModelVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivilegeModelVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetStorageRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetStorageRootCredentialId() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRootCredentialId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetStorageRootCredentialName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageRootCredentialName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksMetastoreMetastoreInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

