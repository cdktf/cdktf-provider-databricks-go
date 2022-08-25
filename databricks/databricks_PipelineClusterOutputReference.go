// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/jsii"

	"github.com/hashicorp/cdktf-provider-databricks-go/databricks/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineClusterOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() PipelineClusterAutoscaleOutputReference
	AutoscaleInput() *PipelineClusterAutoscale
	AwsAttributes() PipelineClusterAwsAttributesOutputReference
	AwsAttributesInput() *PipelineClusterAwsAttributes
	ClusterLogConf() PipelineClusterClusterLogConfOutputReference
	ClusterLogConfInput() *PipelineClusterClusterLogConf
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
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	DriverInstancePoolId() *string
	SetDriverInstancePoolId(val *string)
	DriverInstancePoolIdInput() *string
	DriverNodeTypeId() *string
	SetDriverNodeTypeId(val *string)
	DriverNodeTypeIdInput() *string
	// Experimental.
	Fqn() *string
	GcpAttributes() PipelineClusterGcpAttributesOutputReference
	GcpAttributesInput() *PipelineClusterGcpAttributes
	InitScripts() PipelineClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	SparkConf() *map[string]*string
	SetSparkConf(val *map[string]*string)
	SparkConfInput() *map[string]*string
	SparkEnvVars() *map[string]*string
	SetSparkEnvVars(val *map[string]*string)
	SparkEnvVarsInput() *map[string]*string
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
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
	PutAutoscale(value *PipelineClusterAutoscale)
	PutAwsAttributes(value *PipelineClusterAwsAttributes)
	PutClusterLogConf(value *PipelineClusterClusterLogConf)
	PutGcpAttributes(value *PipelineClusterGcpAttributes)
	PutInitScripts(value interface{})
	ResetApplyPolicyDefaultValues()
	ResetAutoscale()
	ResetAwsAttributes()
	ResetClusterLogConf()
	ResetCustomTags()
	ResetDriverInstancePoolId()
	ResetDriverNodeTypeId()
	ResetGcpAttributes()
	ResetInitScripts()
	ResetInstancePoolId()
	ResetLabel()
	ResetNodeTypeId()
	ResetNumWorkers()
	ResetPolicyId()
	ResetSparkConf()
	ResetSparkEnvVars()
	ResetSshPublicKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineClusterOutputReference
type jsiiProxy_PipelineClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineClusterOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) Autoscale() PipelineClusterAutoscaleOutputReference {
	var returns PipelineClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) AutoscaleInput() *PipelineClusterAutoscale {
	var returns *PipelineClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) AwsAttributes() PipelineClusterAwsAttributesOutputReference {
	var returns PipelineClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) AwsAttributesInput() *PipelineClusterAwsAttributes {
	var returns *PipelineClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) ClusterLogConf() PipelineClusterClusterLogConfOutputReference {
	var returns PipelineClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) ClusterLogConfInput() *PipelineClusterClusterLogConf {
	var returns *PipelineClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) GcpAttributes() PipelineClusterGcpAttributesOutputReference {
	var returns PipelineClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) GcpAttributesInput() *PipelineClusterGcpAttributes {
	var returns *PipelineClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) InitScripts() PipelineClusterInitScriptsList {
	var returns PipelineClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPipelineClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PipelineClusterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_PipelineClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.PipelineClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPipelineClusterOutputReference_Override(p PipelineClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.PipelineClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetApplyPolicyDefaultValues(val interface{}) {
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetCustomTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetDriverInstancePoolId(val *string) {
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetDriverNodeTypeId(val *string) {
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetInstancePoolId(val *string) {
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetNodeTypeId(val *string) {
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetNumWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetPolicyId(val *string) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetSparkConf(val *map[string]*string) {
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetSparkEnvVars(val *map[string]*string) {
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetSshPublicKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineClusterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) PutAutoscale(value *PipelineClusterAutoscale) {
	_jsii_.InvokeVoid(
		p,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) PutAwsAttributes(value *PipelineClusterAwsAttributes) {
	_jsii_.InvokeVoid(
		p,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) PutClusterLogConf(value *PipelineClusterClusterLogConf) {
	_jsii_.InvokeVoid(
		p,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) PutGcpAttributes(value *PipelineClusterGcpAttributes) {
	_jsii_.InvokeVoid(
		p,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) PutInitScripts(value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		p,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		p,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		p,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		p,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		p,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		p,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		p,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		p,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		p,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		p,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

