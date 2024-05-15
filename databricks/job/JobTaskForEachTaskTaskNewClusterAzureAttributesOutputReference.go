// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference interface {
	cdktf.ComplexObject
	Availability() *string
	SetAvailability(val *string)
	AvailabilityInput() *string
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
	FirstOnDemand() *float64
	SetFirstOnDemand(val *float64)
	FirstOnDemandInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *JobTaskForEachTaskTaskNewClusterAzureAttributes
	SetInternalValue(val *JobTaskForEachTaskTaskNewClusterAzureAttributes)
	LogAnalyticsInfo() JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfoOutputReference
	LogAnalyticsInfoInput() *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo
	SpotBidMaxPrice() *float64
	SetSpotBidMaxPrice(val *float64)
	SpotBidMaxPriceInput() *float64
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
	PutLogAnalyticsInfo(value *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo)
	ResetAvailability()
	ResetFirstOnDemand()
	ResetLogAnalyticsInfo()
	ResetSpotBidMaxPrice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference
type jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) FirstOnDemand() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) FirstOnDemandInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) InternalValue() *JobTaskForEachTaskTaskNewClusterAzureAttributes {
	var returns *JobTaskForEachTaskTaskNewClusterAzureAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) LogAnalyticsInfo() JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfoOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfoOutputReference
	_jsii_.Get(
		j,
		"logAnalyticsInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) LogAnalyticsInfoInput() *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo {
	var returns *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo
	_jsii_.Get(
		j,
		"logAnalyticsInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) SpotBidMaxPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) SpotBidMaxPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskNewClusterAzureAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference_Override(j JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetFirstOnDemand(val *float64) {
	if err := j.validateSetFirstOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstOnDemand",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskNewClusterAzureAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetSpotBidMaxPrice(val *float64) {
	if err := j.validateSetSpotBidMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotBidMaxPrice",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) PutLogAnalyticsInfo(value *JobTaskForEachTaskTaskNewClusterAzureAttributesLogAnalyticsInfo) {
	if err := j.validatePutLogAnalyticsInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLogAnalyticsInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		j,
		"resetAvailability",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ResetFirstOnDemand() {
	_jsii_.InvokeVoid(
		j,
		"resetFirstOnDemand",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ResetLogAnalyticsInfo() {
	_jsii_.InvokeVoid(
		j,
		"resetLogAnalyticsInfo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ResetSpotBidMaxPrice() {
	_jsii_.InvokeVoid(
		j,
		"resetSpotBidMaxPrice",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

