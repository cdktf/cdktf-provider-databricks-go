// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
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
	InstancePoolsToUseCount() *float64
	SetInstancePoolsToUseCount(val *float64)
	InstancePoolsToUseCountInput() *float64
	InternalValue() *InstancePoolInstancePoolFleetAttributesFleetSpotOption
	SetInternalValue(val *InstancePoolInstancePoolFleetAttributesFleetSpotOption)
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
	ResetInstancePoolsToUseCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference
type jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) InternalValue() *InstancePoolInstancePoolFleetAttributesFleetSpotOption {
	var returns *InstancePoolInstancePoolFleetAttributesFleetSpotOption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference_Override(i InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetAllocationStrategy(val *string) {
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetInstancePoolsToUseCount(val *float64) {
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetInternalValue(val *InstancePoolInstancePoolFleetAttributesFleetSpotOption) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		i,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
