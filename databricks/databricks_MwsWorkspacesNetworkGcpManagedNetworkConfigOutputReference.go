// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference interface {
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
	GkeClusterPodIpRange() *string
	SetGkeClusterPodIpRange(val *string)
	GkeClusterPodIpRangeInput() *string
	GkeClusterServiceIpRange() *string
	SetGkeClusterServiceIpRange(val *string)
	GkeClusterServiceIpRangeInput() *string
	InternalValue() *MwsWorkspacesNetworkGcpManagedNetworkConfig
	SetInternalValue(val *MwsWorkspacesNetworkGcpManagedNetworkConfig)
	SubnetCidr() *string
	SetSubnetCidr(val *string)
	SubnetCidrInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference
type jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GkeClusterPodIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeClusterPodIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GkeClusterPodIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeClusterPodIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GkeClusterServiceIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeClusterServiceIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GkeClusterServiceIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeClusterServiceIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) InternalValue() *MwsWorkspacesNetworkGcpManagedNetworkConfig {
	var returns *MwsWorkspacesNetworkGcpManagedNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference_Override(m MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetGkeClusterPodIpRange(val *string) {
	_jsii_.Set(
		j,
		"gkeClusterPodIpRange",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetGkeClusterServiceIpRange(val *string) {
	_jsii_.Set(
		j,
		"gkeClusterServiceIpRange",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetInternalValue(val *MwsWorkspacesNetworkGcpManagedNetworkConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetSubnetCidr(val *string) {
	_jsii_.Set(
		j,
		"subnetCidr",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MwsWorkspacesNetworkGcpManagedNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
