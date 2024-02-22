// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksstoragecredential

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksstoragecredential/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksStorageCredentialStorageCredentialInfoOutputReference interface {
	cdktf.ComplexObject
	AwsIamRole() DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRoleOutputReference
	AwsIamRoleInput() *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole
	AzureManagedIdentity() DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference
	AzureManagedIdentityInput() *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity
	AzureServicePrincipal() DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipalOutputReference
	AzureServicePrincipalInput() *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal
	CloudflareApiToken() DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiTokenOutputReference
	CloudflareApiTokenInput() *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken
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
	DatabricksGcpServiceAccount() DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccountOutputReference
	DatabricksGcpServiceAccountInput() *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *DataDatabricksStorageCredentialStorageCredentialInfo
	SetInternalValue(val *DataDatabricksStorageCredentialStorageCredentialInfo)
	MetastoreId() *string
	SetMetastoreId(val *string)
	MetastoreIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
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
	UsedForManagedStorage() interface{}
	SetUsedForManagedStorage(val interface{})
	UsedForManagedStorageInput() interface{}
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
	PutAwsIamRole(value *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole)
	PutAzureManagedIdentity(value *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity)
	PutAzureServicePrincipal(value *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal)
	PutCloudflareApiToken(value *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken)
	PutDatabricksGcpServiceAccount(value *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount)
	ResetAwsIamRole()
	ResetAzureManagedIdentity()
	ResetAzureServicePrincipal()
	ResetCloudflareApiToken()
	ResetComment()
	ResetCreatedAt()
	ResetCreatedBy()
	ResetDatabricksGcpServiceAccount()
	ResetId()
	ResetMetastoreId()
	ResetName()
	ResetOwner()
	ResetReadOnly()
	ResetUpdatedAt()
	ResetUpdatedBy()
	ResetUsedForManagedStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksStorageCredentialStorageCredentialInfoOutputReference
type jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AwsIamRole() DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRoleOutputReference {
	var returns DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRoleOutputReference
	_jsii_.Get(
		j,
		"awsIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AwsIamRoleInput() *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole
	_jsii_.Get(
		j,
		"awsIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AzureManagedIdentity() DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference {
	var returns DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"azureManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AzureManagedIdentityInput() *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity
	_jsii_.Get(
		j,
		"azureManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AzureServicePrincipal() DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipalOutputReference {
	var returns DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"azureServicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) AzureServicePrincipalInput() *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal
	_jsii_.Get(
		j,
		"azureServicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CloudflareApiToken() DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiTokenOutputReference {
	var returns DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiTokenOutputReference
	_jsii_.Get(
		j,
		"cloudflareApiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CloudflareApiTokenInput() *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken
	_jsii_.Get(
		j,
		"cloudflareApiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CreatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) DatabricksGcpServiceAccount() DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccountOutputReference {
	var returns DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccountOutputReference
	_jsii_.Get(
		j,
		"databricksGcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) DatabricksGcpServiceAccountInput() *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount
	_jsii_.Get(
		j,
		"databricksGcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) InternalValue() *DataDatabricksStorageCredentialStorageCredentialInfo {
	var returns *DataDatabricksStorageCredentialStorageCredentialInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) MetastoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) MetastoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UpdatedAtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UpdatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UpdatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UsedForManagedStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usedForManagedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) UsedForManagedStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usedForManagedStorageInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksStorageCredentialStorageCredentialInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksStorageCredentialStorageCredentialInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksStorageCredentialStorageCredentialInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksStorageCredential.DataDatabricksStorageCredentialStorageCredentialInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksStorageCredentialStorageCredentialInfoOutputReference_Override(d DataDatabricksStorageCredentialStorageCredentialInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksStorageCredential.DataDatabricksStorageCredentialStorageCredentialInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetCreatedAt(val *float64) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetInternalValue(val *DataDatabricksStorageCredentialStorageCredentialInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetMetastoreId(val *string) {
	if err := j.validateSetMetastoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetUpdatedAt(val *float64) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetUpdatedBy(val *string) {
	if err := j.validateSetUpdatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedBy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference)SetUsedForManagedStorage(val interface{}) {
	if err := j.validateSetUsedForManagedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usedForManagedStorage",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) PutAwsIamRole(value *DataDatabricksStorageCredentialStorageCredentialInfoAwsIamRole) {
	if err := d.validatePutAwsIamRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsIamRole",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) PutAzureManagedIdentity(value *DataDatabricksStorageCredentialStorageCredentialInfoAzureManagedIdentity) {
	if err := d.validatePutAzureManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureManagedIdentity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) PutAzureServicePrincipal(value *DataDatabricksStorageCredentialStorageCredentialInfoAzureServicePrincipal) {
	if err := d.validatePutAzureServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureServicePrincipal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) PutCloudflareApiToken(value *DataDatabricksStorageCredentialStorageCredentialInfoCloudflareApiToken) {
	if err := d.validatePutCloudflareApiTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudflareApiToken",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) PutDatabricksGcpServiceAccount(value *DataDatabricksStorageCredentialStorageCredentialInfoDatabricksGcpServiceAccount) {
	if err := d.validatePutDatabricksGcpServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatabricksGcpServiceAccount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetAwsIamRole() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsIamRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetAzureManagedIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureManagedIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetAzureServicePrincipal() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureServicePrincipal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetCloudflareApiToken() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudflareApiToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetDatabricksGcpServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabricksGcpServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetMetastoreId() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetUpdatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ResetUsedForManagedStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetUsedForManagedStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksStorageCredentialStorageCredentialInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

