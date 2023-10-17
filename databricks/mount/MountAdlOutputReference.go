// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/mount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MountAdlOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecretKey() *string
	SetClientSecretKey(val *string)
	ClientSecretKeyInput() *string
	ClientSecretScope() *string
	SetClientSecretScope(val *string)
	ClientSecretScopeInput() *string
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
	Directory() *string
	SetDirectory(val *string)
	DirectoryInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MountAdl
	SetInternalValue(val *MountAdl)
	SparkConfPrefix() *string
	SetSparkConfPrefix(val *string)
	SparkConfPrefixInput() *string
	StorageResourceName() *string
	SetStorageResourceName(val *string)
	StorageResourceNameInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetDirectory()
	ResetSparkConfPrefix()
	ResetStorageResourceName()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MountAdlOutputReference
type jsiiProxy_MountAdlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MountAdlOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ClientSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ClientSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ClientSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ClientSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) DirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) InternalValue() *MountAdl {
	var returns *MountAdl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) SparkConfPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkConfPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) SparkConfPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkConfPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) StorageResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) StorageResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAdlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMountAdlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MountAdlOutputReference {
	_init_.Initialize()

	if err := validateNewMountAdlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MountAdlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mount.MountAdlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMountAdlOutputReference_Override(m MountAdlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mount.MountAdlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetClientSecretKey(val *string) {
	if err := j.validateSetClientSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretKey",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetClientSecretScope(val *string) {
	if err := j.validateSetClientSecretScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretScope",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetDirectory(val *string) {
	if err := j.validateSetDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directory",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetInternalValue(val *MountAdl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetSparkConfPrefix(val *string) {
	if err := j.validateSetSparkConfPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConfPrefix",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetStorageResourceName(val *string) {
	if err := j.validateSetStorageResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageResourceName",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MountAdlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MountAdlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) ResetDirectory() {
	_jsii_.InvokeVoid(
		m,
		"resetDirectory",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAdlOutputReference) ResetSparkConfPrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetSparkConfPrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAdlOutputReference) ResetStorageResourceName() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageResourceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAdlOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		m,
		"resetTenantId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAdlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAdlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

