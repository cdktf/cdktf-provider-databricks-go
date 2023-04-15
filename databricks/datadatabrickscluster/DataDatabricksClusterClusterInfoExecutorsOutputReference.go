package datadatabrickscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/datadatabrickscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterClusterInfoExecutorsOutputReference interface {
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
	HostPrivateIp() *string
	SetHostPrivateIp(val *string)
	HostPrivateIpInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NodeAwsAttributes() DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributesOutputReference
	NodeAwsAttributesInput() *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes
	NodeId() *string
	SetNodeId(val *string)
	NodeIdInput() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	PublicDns() *string
	SetPublicDns(val *string)
	PublicDnsInput() *string
	StartTimestamp() *float64
	SetStartTimestamp(val *float64)
	StartTimestampInput() *float64
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
	PutNodeAwsAttributes(value *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes)
	ResetHostPrivateIp()
	ResetInstanceId()
	ResetNodeAwsAttributes()
	ResetNodeId()
	ResetPrivateIp()
	ResetPublicDns()
	ResetStartTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksClusterClusterInfoExecutorsOutputReference
type jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) HostPrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) HostPrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPrivateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) NodeAwsAttributes() DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"nodeAwsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) NodeAwsAttributesInput() *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes {
	var returns *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes
	_jsii_.Get(
		j,
		"nodeAwsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) NodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) NodeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) PublicDnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) StartTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) StartTimestampInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterClusterInfoExecutorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksClusterClusterInfoExecutorsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterClusterInfoExecutorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoExecutorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterClusterInfoExecutorsOutputReference_Override(d DataDatabricksClusterClusterInfoExecutorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoExecutorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetHostPrivateIp(val *string) {
	if err := j.validateSetHostPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPrivateIp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetNodeId(val *string) {
	if err := j.validateSetNodeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetPublicDns(val *string) {
	if err := j.validateSetPublicDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicDns",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetStartTimestamp(val *float64) {
	if err := j.validateSetStartTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimestamp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) PutNodeAwsAttributes(value *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes) {
	if err := d.validatePutNodeAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNodeAwsAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetHostPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetInstanceId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetNodeAwsAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeAwsAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetNodeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetPublicDns() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicDns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ResetStartTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetStartTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoExecutorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

