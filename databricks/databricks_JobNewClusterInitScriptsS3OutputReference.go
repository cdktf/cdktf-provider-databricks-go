// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobNewClusterInitScriptsS3OutputReference interface {
	cdktf.ComplexObject
	CannedAcl() *string
	SetCannedAcl(val *string)
	CannedAclInput() *string
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
	Destination() *string
	SetDestination(val *string)
	DestinationInput() *string
	EnableEncryption() interface{}
	SetEnableEncryption(val interface{})
	EnableEncryptionInput() interface{}
	EncryptionType() *string
	SetEncryptionType(val *string)
	EncryptionTypeInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *JobNewClusterInitScriptsS3
	SetInternalValue(val *JobNewClusterInitScriptsS3)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetCannedAcl()
	ResetEnableEncryption()
	ResetEncryptionType()
	ResetEndpoint()
	ResetKmsKey()
	ResetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobNewClusterInitScriptsS3OutputReference
type jsiiProxy_JobNewClusterInitScriptsS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) CannedAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) CannedAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) EnableEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) EnableEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) InternalValue() *JobNewClusterInitScriptsS3 {
	var returns *JobNewClusterInitScriptsS3
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobNewClusterInitScriptsS3OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobNewClusterInitScriptsS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_JobNewClusterInitScriptsS3OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.JobNewClusterInitScriptsS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobNewClusterInitScriptsS3OutputReference_Override(j JobNewClusterInitScriptsS3OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.JobNewClusterInitScriptsS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetCannedAcl(val *string) {
	_jsii_.Set(
		j,
		"cannedAcl",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetDestination(val *string) {
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetEnableEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"enableEncryption",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetEncryptionType(val *string) {
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetInternalValue(val *JobNewClusterInitScriptsS3) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetCannedAcl() {
	_jsii_.InvokeVoid(
		j,
		"resetCannedAcl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetEnableEncryption() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableEncryption",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		j,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		j,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		j,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		j,
		"resetRegion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterInitScriptsS3OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
