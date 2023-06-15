package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v8/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v8/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobJobClusterNewClusterOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() JobJobClusterNewClusterAutoscaleOutputReference
	AutoscaleInput() *JobJobClusterNewClusterAutoscale
	AutoterminationMinutes() *float64
	SetAutoterminationMinutes(val *float64)
	AutoterminationMinutesInput() *float64
	AwsAttributes() JobJobClusterNewClusterAwsAttributesOutputReference
	AwsAttributesInput() *JobJobClusterNewClusterAwsAttributes
	AzureAttributes() JobJobClusterNewClusterAzureAttributesOutputReference
	AzureAttributesInput() *JobJobClusterNewClusterAzureAttributes
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() JobJobClusterNewClusterClusterLogConfOutputReference
	ClusterLogConfInput() *JobJobClusterNewClusterClusterLogConf
	ClusterMountInfo() JobJobClusterNewClusterClusterMountInfoList
	ClusterMountInfoInput() interface{}
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	DataSecurityMode() *string
	SetDataSecurityMode(val *string)
	DataSecurityModeInput() *string
	DockerImage() JobJobClusterNewClusterDockerImageOutputReference
	DockerImageInput() *JobJobClusterNewClusterDockerImage
	DriverInstancePoolId() *string
	SetDriverInstancePoolId(val *string)
	DriverInstancePoolIdInput() *string
	DriverNodeTypeId() *string
	SetDriverNodeTypeId(val *string)
	DriverNodeTypeIdInput() *string
	EnableElasticDisk() interface{}
	SetEnableElasticDisk(val interface{})
	EnableElasticDiskInput() interface{}
	EnableLocalDiskEncryption() interface{}
	SetEnableLocalDiskEncryption(val interface{})
	EnableLocalDiskEncryptionInput() interface{}
	// Experimental.
	Fqn() *string
	GcpAttributes() JobJobClusterNewClusterGcpAttributesOutputReference
	GcpAttributesInput() *JobJobClusterNewClusterGcpAttributes
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	InitScripts() JobJobClusterNewClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *JobJobClusterNewCluster
	SetInternalValue(val *JobJobClusterNewCluster)
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	RuntimeEngine() *string
	SetRuntimeEngine(val *string)
	RuntimeEngineInput() *string
	SingleUserName() *string
	SetSingleUserName(val *string)
	SingleUserNameInput() *string
	SparkConf() *map[string]*string
	SetSparkConf(val *map[string]*string)
	SparkConfInput() *map[string]*string
	SparkEnvVars() *map[string]*string
	SetSparkEnvVars(val *map[string]*string)
	SparkEnvVarsInput() *map[string]*string
	SparkVersion() *string
	SetSparkVersion(val *string)
	SparkVersionInput() *string
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
	WorkloadType() JobJobClusterNewClusterWorkloadTypeOutputReference
	WorkloadTypeInput() *JobJobClusterNewClusterWorkloadType
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
	PutAutoscale(value *JobJobClusterNewClusterAutoscale)
	PutAwsAttributes(value *JobJobClusterNewClusterAwsAttributes)
	PutAzureAttributes(value *JobJobClusterNewClusterAzureAttributes)
	PutClusterLogConf(value *JobJobClusterNewClusterClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *JobJobClusterNewClusterDockerImage)
	PutGcpAttributes(value *JobJobClusterNewClusterGcpAttributes)
	PutInitScripts(value interface{})
	PutWorkloadType(value *JobJobClusterNewClusterWorkloadType)
	ResetApplyPolicyDefaultValues()
	ResetAutoscale()
	ResetAutoterminationMinutes()
	ResetAwsAttributes()
	ResetAzureAttributes()
	ResetClusterId()
	ResetClusterLogConf()
	ResetClusterMountInfo()
	ResetClusterName()
	ResetCustomTags()
	ResetDataSecurityMode()
	ResetDockerImage()
	ResetDriverInstancePoolId()
	ResetDriverNodeTypeId()
	ResetEnableElasticDisk()
	ResetEnableLocalDiskEncryption()
	ResetGcpAttributes()
	ResetIdempotencyToken()
	ResetInitScripts()
	ResetInstancePoolId()
	ResetNodeTypeId()
	ResetNumWorkers()
	ResetPolicyId()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkEnvVars()
	ResetSshPublicKeys()
	ResetWorkloadType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobJobClusterNewClusterOutputReference
type jsiiProxy_JobJobClusterNewClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) Autoscale() JobJobClusterNewClusterAutoscaleOutputReference {
	var returns JobJobClusterNewClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AutoscaleInput() *JobJobClusterNewClusterAutoscale {
	var returns *JobJobClusterNewClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AwsAttributes() JobJobClusterNewClusterAwsAttributesOutputReference {
	var returns JobJobClusterNewClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AwsAttributesInput() *JobJobClusterNewClusterAwsAttributes {
	var returns *JobJobClusterNewClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AzureAttributes() JobJobClusterNewClusterAzureAttributesOutputReference {
	var returns JobJobClusterNewClusterAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) AzureAttributesInput() *JobJobClusterNewClusterAzureAttributes {
	var returns *JobJobClusterNewClusterAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterLogConf() JobJobClusterNewClusterClusterLogConfOutputReference {
	var returns JobJobClusterNewClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterLogConfInput() *JobJobClusterNewClusterClusterLogConf {
	var returns *JobJobClusterNewClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterMountInfo() JobJobClusterNewClusterClusterMountInfoList {
	var returns JobJobClusterNewClusterClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DockerImage() JobJobClusterNewClusterDockerImageOutputReference {
	var returns JobJobClusterNewClusterDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DockerImageInput() *JobJobClusterNewClusterDockerImage {
	var returns *JobJobClusterNewClusterDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GcpAttributes() JobJobClusterNewClusterGcpAttributesOutputReference {
	var returns JobJobClusterNewClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GcpAttributesInput() *JobJobClusterNewClusterGcpAttributes {
	var returns *JobJobClusterNewClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InitScripts() JobJobClusterNewClusterInitScriptsList {
	var returns JobJobClusterNewClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InternalValue() *JobJobClusterNewCluster {
	var returns *JobJobClusterNewCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) WorkloadType() JobJobClusterNewClusterWorkloadTypeOutputReference {
	var returns JobJobClusterNewClusterWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) WorkloadTypeInput() *JobJobClusterNewClusterWorkloadType {
	var returns *JobJobClusterNewClusterWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewJobJobClusterNewClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobJobClusterNewClusterOutputReference {
	_init_.Initialize()

	if err := validateNewJobJobClusterNewClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobJobClusterNewClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobJobClusterNewClusterOutputReference_Override(j JobJobClusterNewClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetAutoterminationMinutes(val *float64) {
	if err := j.validateSetAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetInternalValue(val *JobJobClusterNewCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutAutoscale(value *JobJobClusterNewClusterAutoscale) {
	if err := j.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutAwsAttributes(value *JobJobClusterNewClusterAwsAttributes) {
	if err := j.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutAzureAttributes(value *JobJobClusterNewClusterAzureAttributes) {
	if err := j.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutClusterLogConf(value *JobJobClusterNewClusterClusterLogConf) {
	if err := j.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutClusterMountInfo(value interface{}) {
	if err := j.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutDockerImage(value *JobJobClusterNewClusterDockerImage) {
	if err := j.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutGcpAttributes(value *JobJobClusterNewClusterGcpAttributes) {
	if err := j.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutInitScripts(value interface{}) {
	if err := j.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) PutWorkloadType(value *JobJobClusterNewClusterWorkloadType) {
	if err := j.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		j,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetAutoterminationMinutes() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoterminationMinutes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		j,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		j,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		j,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		j,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		j,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		j,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		j,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		j,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		j,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

