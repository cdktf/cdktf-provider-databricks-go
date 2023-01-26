package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalMountDirPath() *string
	SetLocalMountDirPath(val *string)
	LocalMountDirPathInput() *string
	NetworkFilesystemInfo() DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference
	NetworkFilesystemInfoInput() *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo
	RemoteMountDirPath() *string
	SetRemoteMountDirPath(val *string)
	RemoteMountDirPathInput() *string
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
	PutNetworkFilesystemInfo(value *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo)
	ResetRemoteMountDirPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) LocalMountDirPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountDirPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) LocalMountDirPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localMountDirPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) NetworkFilesystemInfo() DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfoOutputReference
	_jsii_.Get(
		j,
		"networkFilesystemInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) NetworkFilesystemInfoInput() *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo {
	var returns *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo
	_jsii_.Get(
		j,
		"networkFilesystemInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) RemoteMountDirPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteMountDirPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) RemoteMountDirPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteMountDirPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetLocalMountDirPath(val *string) {
	if err := j.validateSetLocalMountDirPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMountDirPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetRemoteMountDirPath(val *string) {
	if err := j.validateSetRemoteMountDirPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteMountDirPath",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) PutNetworkFilesystemInfo(value *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoNetworkFilesystemInfo) {
	if err := d.validatePutNetworkFilesystemInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkFilesystemInfo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) ResetRemoteMountDirPath() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoteMountDirPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterMountInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

