package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobNewClusterAwsAttributesOutputReference interface {
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
	EbsVolumeSize() *float64
	SetEbsVolumeSize(val *float64)
	EbsVolumeSizeInput() *float64
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
	InternalValue() *JobNewClusterAwsAttributes
	SetInternalValue(val *JobNewClusterAwsAttributes)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAvailability()
	ResetEbsVolumeCount()
	ResetEbsVolumeSize()
	ResetEbsVolumeType()
	ResetFirstOnDemand()
	ResetInstanceProfileArn()
	ResetSpotBidPricePercent()
	ResetZoneId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobNewClusterAwsAttributesOutputReference
type jsiiProxy_JobNewClusterAwsAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) EbsVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) FirstOnDemand() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) FirstOnDemandInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstOnDemandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) InternalValue() *JobNewClusterAwsAttributes {
	var returns *JobNewClusterAwsAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) SpotBidPricePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPricePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) SpotBidPricePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotBidPricePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewJobNewClusterAwsAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobNewClusterAwsAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewJobNewClusterAwsAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobNewClusterAwsAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobNewClusterAwsAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobNewClusterAwsAttributesOutputReference_Override(j JobNewClusterAwsAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobNewClusterAwsAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetEbsVolumeCount(val *float64) {
	if err := j.validateSetEbsVolumeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeCount",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetEbsVolumeSize(val *float64) {
	if err := j.validateSetEbsVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeSize",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetEbsVolumeType(val *string) {
	if err := j.validateSetEbsVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeType",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetFirstOnDemand(val *float64) {
	if err := j.validateSetFirstOnDemandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstOnDemand",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetInternalValue(val *JobNewClusterAwsAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetSpotBidPricePercent(val *float64) {
	if err := j.validateSetSpotBidPricePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotBidPricePercent",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		j,
		"resetAvailability",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetEbsVolumeCount() {
	_jsii_.InvokeVoid(
		j,
		"resetEbsVolumeCount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetEbsVolumeSize() {
	_jsii_.InvokeVoid(
		j,
		"resetEbsVolumeSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetEbsVolumeType() {
	_jsii_.InvokeVoid(
		j,
		"resetEbsVolumeType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetFirstOnDemand() {
	_jsii_.InvokeVoid(
		j,
		"resetFirstOnDemand",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		j,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetSpotBidPricePercent() {
	_jsii_.InvokeVoid(
		j,
		"resetSpotBidPricePercent",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		j,
		"resetZoneId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobNewClusterAwsAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

