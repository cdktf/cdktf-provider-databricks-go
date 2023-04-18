package mwsvpcendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/mwsvpcendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsVpcEndpointGcpVpcEndpointInfoOutputReference interface {
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
	EndpointRegion() *string
	SetEndpointRegion(val *string)
	EndpointRegionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MwsVpcEndpointGcpVpcEndpointInfo
	SetInternalValue(val *MwsVpcEndpointGcpVpcEndpointInfo)
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	PscConnectionId() *string
	SetPscConnectionId(val *string)
	PscConnectionIdInput() *string
	PscEndpointName() *string
	SetPscEndpointName(val *string)
	PscEndpointNameInput() *string
	ServiceAttachmentId() *string
	SetServiceAttachmentId(val *string)
	ServiceAttachmentIdInput() *string
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
	ResetPscConnectionId()
	ResetServiceAttachmentId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsVpcEndpointGcpVpcEndpointInfoOutputReference
type jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) EndpointRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) EndpointRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) InternalValue() *MwsVpcEndpointGcpVpcEndpointInfo {
	var returns *MwsVpcEndpointGcpVpcEndpointInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) PscConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) PscConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) PscEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) PscEndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscEndpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ServiceAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ServiceAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsVpcEndpointGcpVpcEndpointInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsVpcEndpointGcpVpcEndpointInfoOutputReference {
	_init_.Initialize()

	if err := validateNewMwsVpcEndpointGcpVpcEndpointInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpointGcpVpcEndpointInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsVpcEndpointGcpVpcEndpointInfoOutputReference_Override(m MwsVpcEndpointGcpVpcEndpointInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.mwsVpcEndpoint.MwsVpcEndpointGcpVpcEndpointInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetEndpointRegion(val *string) {
	if err := j.validateSetEndpointRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointRegion",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetInternalValue(val *MwsVpcEndpointGcpVpcEndpointInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetPscConnectionId(val *string) {
	if err := j.validateSetPscConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pscConnectionId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetPscEndpointName(val *string) {
	if err := j.validateSetPscEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pscEndpointName",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetServiceAttachmentId(val *string) {
	if err := j.validateSetServiceAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAttachmentId",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ResetPscConnectionId() {
	_jsii_.InvokeVoid(
		m,
		"resetPscConnectionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ResetServiceAttachmentId() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceAttachmentId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MwsVpcEndpointGcpVpcEndpointInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

