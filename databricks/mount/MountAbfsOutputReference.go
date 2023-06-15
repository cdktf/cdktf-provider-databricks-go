package mount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v8/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v8/mount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MountAbfsOutputReference interface {
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
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
	InitializeFileSystem() interface{}
	SetInitializeFileSystem(val interface{})
	InitializeFileSystemInput() interface{}
	InternalValue() *MountAbfs
	SetInternalValue(val *MountAbfs)
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
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
	ResetContainerName()
	ResetDirectory()
	ResetStorageAccountName()
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

// The jsii proxy struct for MountAbfsOutputReference
type jsiiProxy_MountAbfsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientSecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientSecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientSecretScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ClientSecretScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) DirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) InitializeFileSystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializeFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) InitializeFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializeFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) InternalValue() *MountAbfs {
	var returns *MountAbfs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MountAbfsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMountAbfsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MountAbfsOutputReference {
	_init_.Initialize()

	if err := validateNewMountAbfsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MountAbfsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mount.MountAbfsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMountAbfsOutputReference_Override(m MountAbfsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mount.MountAbfsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetClientSecretKey(val *string) {
	if err := j.validateSetClientSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretKey",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetClientSecretScope(val *string) {
	if err := j.validateSetClientSecretScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretScope",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetDirectory(val *string) {
	if err := j.validateSetDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directory",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetInitializeFileSystem(val interface{}) {
	if err := j.validateSetInitializeFileSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializeFileSystem",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetInternalValue(val *MountAbfs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MountAbfsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MountAbfsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAbfsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MountAbfsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MountAbfsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MountAbfsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MountAbfsOutputReference) ResetContainerName() {
	_jsii_.InvokeVoid(
		m,
		"resetContainerName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAbfsOutputReference) ResetDirectory() {
	_jsii_.InvokeVoid(
		m,
		"resetDirectory",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAbfsOutputReference) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAbfsOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		m,
		"resetTenantId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MountAbfsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MountAbfsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

