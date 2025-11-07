// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineClusterAwsAttributesOutputReference interface {
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
	EbsVolumeCount() *float64
	SetEbsVolumeCount(val *float64)
	EbsVolumeCountInput() *float64
	EbsVolumeIops() *float64
	SetEbsVolumeIops(val *float64)
	EbsVolumeIopsInput() *float64
	EbsVolumeSize() *float64
	SetEbsVolumeSize(val *float64)
	EbsVolumeSizeInput() *float64
	EbsVolumeThroughput() *float64
	SetEbsVolumeThroughput(val *float64)
	EbsVolumeThroughputInput() *float64
	EbsVolumeType() *string
	SetEbsVolumeType(val *string)
	EbsVolumeTypeInput() *string
	FirstOnDemand() *float64
	SetFirstOnDemand(val *float64)
	FirstOnDemandInput() *float64
	// Experimental.
	Fqn() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InternalValue() *PipelineClusterAwsAttributes
	SetInternalValue(val *PipelineClusterAwsAttributes)
	SpotBidPricePercent() *float64
	SetSpotBidPricePercent(val *float64)
	SpotBidPricePercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAvailability()
	ResetEbsVolumeCount()
	ResetEbsVolumeIops()
	ResetEbsVolumeSize()
	ResetEbsVolumeThroughput()
	ResetEbsVolumeType()
	ResetFirstOnDemand()
	ResetInstanceProfileArn()
	ResetSpotBidPricePercent()
	ResetZoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineClusterAwsAttributesOutputReference
type jsiiProxy_PipelineClusterAwsAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) EbsVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) FirstOnDemand() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) FirstOnDemandInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) InternalValue() *PipelineClusterAwsAttributes {
	var returns *PipelineClusterAwsAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) SpotBidPricePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPricePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) SpotBidPricePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPricePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewPipelineClusterAwsAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipelineClusterAwsAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineClusterAwsAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineClusterAwsAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineClusterAwsAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineClusterAwsAttributesOutputReference_Override(p PipelineClusterAwsAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.PipelineClusterAwsAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetEbsVolumeCount(val *float64) {
	if err := j.validateSetEbsVolumeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeCount",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetEbsVolumeIops(val *float64) {
	if err := j.validateSetEbsVolumeIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeIops",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetEbsVolumeSize(val *float64) {
	if err := j.validateSetEbsVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeSize",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetEbsVolumeThroughput(val *float64) {
	if err := j.validateSetEbsVolumeThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeThroughput",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetEbsVolumeType(val *string) {
	if err := j.validateSetEbsVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeType",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetFirstOnDemand(val *float64) {
	if err := j.validateSetFirstOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstOnDemand",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetInternalValue(val *PipelineClusterAwsAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetSpotBidPricePercent(val *float64) {
	if err := j.validateSetSpotBidPricePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotBidPricePercent",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterAwsAttributesOutputReference)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		p,
		"resetAvailability",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetEbsVolumeCount() {
	_jsii_.InvokeVoid(
		p,
		"resetEbsVolumeCount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetEbsVolumeIops() {
	_jsii_.InvokeVoid(
		p,
		"resetEbsVolumeIops",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetEbsVolumeSize() {
	_jsii_.InvokeVoid(
		p,
		"resetEbsVolumeSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetEbsVolumeThroughput() {
	_jsii_.InvokeVoid(
		p,
		"resetEbsVolumeThroughput",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetEbsVolumeType() {
	_jsii_.InvokeVoid(
		p,
		"resetEbsVolumeType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetFirstOnDemand() {
	_jsii_.InvokeVoid(
		p,
		"resetFirstOnDemand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		p,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetSpotBidPricePercent() {
	_jsii_.InvokeVoid(
		p,
		"resetSpotBidPricePercent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		p,
		"resetZoneId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterAwsAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

