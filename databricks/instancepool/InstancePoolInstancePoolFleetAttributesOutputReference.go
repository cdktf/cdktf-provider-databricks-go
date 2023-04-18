package instancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/instancepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstancePoolInstancePoolFleetAttributesOutputReference interface {
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
	FleetOnDemandOption() InstancePoolInstancePoolFleetAttributesFleetOnDemandOptionOutputReference
	FleetOnDemandOptionInput() *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption
	FleetSpotOption() InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference
	FleetSpotOptionInput() *InstancePoolInstancePoolFleetAttributesFleetSpotOption
	// Experimental.
	Fqn() *string
	InternalValue() *InstancePoolInstancePoolFleetAttributes
	SetInternalValue(val *InstancePoolInstancePoolFleetAttributes)
	LaunchTemplateOverride() InstancePoolInstancePoolFleetAttributesLaunchTemplateOverrideList
	LaunchTemplateOverrideInput() interface{}
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
	PutFleetOnDemandOption(value *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption)
	PutFleetSpotOption(value *InstancePoolInstancePoolFleetAttributesFleetSpotOption)
	PutLaunchTemplateOverride(value interface{})
	ResetFleetOnDemandOption()
	ResetFleetSpotOption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InstancePoolInstancePoolFleetAttributesOutputReference
type jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) FleetOnDemandOption() InstancePoolInstancePoolFleetAttributesFleetOnDemandOptionOutputReference {
	var returns InstancePoolInstancePoolFleetAttributesFleetOnDemandOptionOutputReference
	_jsii_.Get(
		j,
		"fleetOnDemandOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) FleetOnDemandOptionInput() *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption {
	var returns *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption
	_jsii_.Get(
		j,
		"fleetOnDemandOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) FleetSpotOption() InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference {
	var returns InstancePoolInstancePoolFleetAttributesFleetSpotOptionOutputReference
	_jsii_.Get(
		j,
		"fleetSpotOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) FleetSpotOptionInput() *InstancePoolInstancePoolFleetAttributesFleetSpotOption {
	var returns *InstancePoolInstancePoolFleetAttributesFleetSpotOption
	_jsii_.Get(
		j,
		"fleetSpotOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) InternalValue() *InstancePoolInstancePoolFleetAttributes {
	var returns *InstancePoolInstancePoolFleetAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) LaunchTemplateOverride() InstancePoolInstancePoolFleetAttributesLaunchTemplateOverrideList {
	var returns InstancePoolInstancePoolFleetAttributesLaunchTemplateOverrideList
	_jsii_.Get(
		j,
		"launchTemplateOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) LaunchTemplateOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInstancePoolInstancePoolFleetAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InstancePoolInstancePoolFleetAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewInstancePoolInstancePoolFleetAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.instancePool.InstancePoolInstancePoolFleetAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInstancePoolInstancePoolFleetAttributesOutputReference_Override(i InstancePoolInstancePoolFleetAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.instancePool.InstancePoolInstancePoolFleetAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference)SetInternalValue(val *InstancePoolInstancePoolFleetAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) PutFleetOnDemandOption(value *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption) {
	if err := i.validatePutFleetOnDemandOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFleetOnDemandOption",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) PutFleetSpotOption(value *InstancePoolInstancePoolFleetAttributesFleetSpotOption) {
	if err := i.validatePutFleetSpotOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFleetSpotOption",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) PutLaunchTemplateOverride(value interface{}) {
	if err := i.validatePutLaunchTemplateOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLaunchTemplateOverride",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ResetFleetOnDemandOption() {
	_jsii_.InvokeVoid(
		i,
		"resetFleetOnDemandOption",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ResetFleetSpotOption() {
	_jsii_.InvokeVoid(
		i,
		"resetFleetSpotOption",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstancePoolInstancePoolFleetAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

