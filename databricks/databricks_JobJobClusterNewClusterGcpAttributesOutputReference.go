// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobJobClusterNewClusterGcpAttributesOutputReference interface {
	cdktf.ComplexObject
	Availability() *string
	SetAvailability(val *string)
	AvailabilityInput() *string
	BootDiskSize() *float64
	SetBootDiskSize(val *float64)
	BootDiskSizeInput() *float64
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
	GoogleServiceAccount() *string
	SetGoogleServiceAccount(val *string)
	GoogleServiceAccountInput() *string
	InternalValue() *JobJobClusterNewClusterGcpAttributes
	SetInternalValue(val *JobJobClusterNewClusterGcpAttributes)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsePreemptibleExecutors() interface{}
	SetUsePreemptibleExecutors(val interface{})
	UsePreemptibleExecutorsInput() interface{}
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
	ResetBootDiskSize()
	ResetGoogleServiceAccount()
	ResetUsePreemptibleExecutors()
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

// The jsii proxy struct for JobJobClusterNewClusterGcpAttributesOutputReference
type jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) BootDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) BootDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GoogleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GoogleServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) InternalValue() *JobJobClusterNewClusterGcpAttributes {
	var returns *JobJobClusterNewClusterGcpAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) UsePreemptibleExecutors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) UsePreemptibleExecutorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePreemptibleExecutorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


func NewJobJobClusterNewClusterGcpAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobJobClusterNewClusterGcpAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.JobJobClusterNewClusterGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobJobClusterNewClusterGcpAttributesOutputReference_Override(j JobJobClusterNewClusterGcpAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.JobJobClusterNewClusterGcpAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetAvailability(val *string) {
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetBootDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"bootDiskSize",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetGoogleServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"googleServiceAccount",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetInternalValue(val *JobJobClusterNewClusterGcpAttributes) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetUsePreemptibleExecutors(val interface{}) {
	_jsii_.Set(
		j,
		"usePreemptibleExecutors",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) SetZoneId(val *string) {
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		j,
		"resetAvailability",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ResetBootDiskSize() {
	_jsii_.InvokeVoid(
		j,
		"resetBootDiskSize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ResetGoogleServiceAccount() {
	_jsii_.InvokeVoid(
		j,
		"resetGoogleServiceAccount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ResetUsePreemptibleExecutors() {
	_jsii_.InvokeVoid(
		j,
		"resetUsePreemptibleExecutors",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ResetZoneId() {
	_jsii_.InvokeVoid(
		j,
		"resetZoneId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterGcpAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
