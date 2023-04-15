package datadatabricksinstancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v6/datadatabricksinstancepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference interface {
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
	InternalValue() *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption
	SetInternalValue(val *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption)
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

// The jsii proxy struct for DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference
type jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) InternalValue() *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption {
	var returns *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference_Override(d DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetInstancePoolsToUseCount(val *float64) {
	if err := j.validateSetInstancePoolsToUseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetInternalValue(val *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

