// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskNewClusterOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() JobTaskNewClusterAutoscaleOutputReference
	AutoscaleInput() *JobTaskNewClusterAutoscale
	AwsAttributes() JobTaskNewClusterAwsAttributesOutputReference
	AwsAttributesInput() *JobTaskNewClusterAwsAttributes
	AzureAttributes() JobTaskNewClusterAzureAttributesOutputReference
	AzureAttributesInput() *JobTaskNewClusterAzureAttributes
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() JobTaskNewClusterClusterLogConfOutputReference
	ClusterLogConfInput() *JobTaskNewClusterClusterLogConf
	ClusterMountInfo() JobTaskNewClusterClusterMountInfoList
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
	DockerImage() JobTaskNewClusterDockerImageOutputReference
	DockerImageInput() *JobTaskNewClusterDockerImage
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
	GcpAttributes() JobTaskNewClusterGcpAttributesOutputReference
	GcpAttributesInput() *JobTaskNewClusterGcpAttributes
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	InitScripts() JobTaskNewClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *JobTaskNewCluster
	SetInternalValue(val *JobTaskNewCluster)
	IsSingleNode() interface{}
	SetIsSingleNode(val interface{})
	IsSingleNodeInput() interface{}
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	Library() JobTaskNewClusterLibraryList
	LibraryInput() interface{}
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
	UseMlRuntime() interface{}
	SetUseMlRuntime(val interface{})
	UseMlRuntimeInput() interface{}
	WorkloadType() JobTaskNewClusterWorkloadTypeOutputReference
	WorkloadTypeInput() *JobTaskNewClusterWorkloadType
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
	PutAutoscale(value *JobTaskNewClusterAutoscale)
	PutAwsAttributes(value *JobTaskNewClusterAwsAttributes)
	PutAzureAttributes(value *JobTaskNewClusterAzureAttributes)
	PutClusterLogConf(value *JobTaskNewClusterClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *JobTaskNewClusterDockerImage)
	PutGcpAttributes(value *JobTaskNewClusterGcpAttributes)
	PutInitScripts(value interface{})
	PutLibrary(value interface{})
	PutWorkloadType(value *JobTaskNewClusterWorkloadType)
	ResetApplyPolicyDefaultValues()
	ResetAutoscale()
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
	ResetIsSingleNode()
	ResetKind()
	ResetLibrary()
	ResetNodeTypeId()
	ResetNumWorkers()
	ResetPolicyId()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkEnvVars()
	ResetSparkVersion()
	ResetSshPublicKeys()
	ResetUseMlRuntime()
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

// The jsii proxy struct for JobTaskNewClusterOutputReference
type jsiiProxy_JobTaskNewClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) Autoscale() JobTaskNewClusterAutoscaleOutputReference {
	var returns JobTaskNewClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) AutoscaleInput() *JobTaskNewClusterAutoscale {
	var returns *JobTaskNewClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) AwsAttributes() JobTaskNewClusterAwsAttributesOutputReference {
	var returns JobTaskNewClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) AwsAttributesInput() *JobTaskNewClusterAwsAttributes {
	var returns *JobTaskNewClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) AzureAttributes() JobTaskNewClusterAzureAttributesOutputReference {
	var returns JobTaskNewClusterAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) AzureAttributesInput() *JobTaskNewClusterAzureAttributes {
	var returns *JobTaskNewClusterAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterLogConf() JobTaskNewClusterClusterLogConfOutputReference {
	var returns JobTaskNewClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterLogConfInput() *JobTaskNewClusterClusterLogConf {
	var returns *JobTaskNewClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterMountInfo() JobTaskNewClusterClusterMountInfoList {
	var returns JobTaskNewClusterClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DockerImage() JobTaskNewClusterDockerImageOutputReference {
	var returns JobTaskNewClusterDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DockerImageInput() *JobTaskNewClusterDockerImage {
	var returns *JobTaskNewClusterDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GcpAttributes() JobTaskNewClusterGcpAttributesOutputReference {
	var returns JobTaskNewClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GcpAttributesInput() *JobTaskNewClusterGcpAttributes {
	var returns *JobTaskNewClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InitScripts() JobTaskNewClusterInitScriptsList {
	var returns JobTaskNewClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InternalValue() *JobTaskNewCluster {
	var returns *JobTaskNewCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) IsSingleNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) IsSingleNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) Library() JobTaskNewClusterLibraryList {
	var returns JobTaskNewClusterLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) UseMlRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) UseMlRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) WorkloadType() JobTaskNewClusterWorkloadTypeOutputReference {
	var returns JobTaskNewClusterWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) WorkloadTypeInput() *JobTaskNewClusterWorkloadType {
	var returns *JobTaskNewClusterWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewJobTaskNewClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskNewClusterOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskNewClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskNewClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskNewClusterOutputReference_Override(j JobTaskNewClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetInternalValue(val *JobTaskNewCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetIsSingleNode(val interface{}) {
	if err := j.validateSetIsSingleNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleNode",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference)SetUseMlRuntime(val interface{}) {
	if err := j.validateSetUseMlRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMlRuntime",
		val,
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutAutoscale(value *JobTaskNewClusterAutoscale) {
	if err := j.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutAwsAttributes(value *JobTaskNewClusterAwsAttributes) {
	if err := j.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutAzureAttributes(value *JobTaskNewClusterAzureAttributes) {
	if err := j.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutClusterLogConf(value *JobTaskNewClusterClusterLogConf) {
	if err := j.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutClusterMountInfo(value interface{}) {
	if err := j.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutDockerImage(value *JobTaskNewClusterDockerImage) {
	if err := j.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutGcpAttributes(value *JobTaskNewClusterGcpAttributes) {
	if err := j.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutInitScripts(value interface{}) {
	if err := j.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutLibrary(value interface{}) {
	if err := j.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) PutWorkloadType(value *JobTaskNewClusterWorkloadType) {
	if err := j.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		j,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		j,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		j,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		j,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		j,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		j,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetIsSingleNode() {
	_jsii_.InvokeVoid(
		j,
		"resetIsSingleNode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		j,
		"resetKind",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		j,
		"resetLibrary",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		j,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		j,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		j,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		j,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetSparkVersion() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkVersion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetUseMlRuntime() {
	_jsii_.InvokeVoid(
		j,
		"resetUseMlRuntime",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskNewClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskNewClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

