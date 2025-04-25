// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskNewClusterOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() JobTaskForEachTaskTaskNewClusterAutoscaleOutputReference
	AutoscaleInput() *JobTaskForEachTaskTaskNewClusterAutoscale
	AwsAttributes() JobTaskForEachTaskTaskNewClusterAwsAttributesOutputReference
	AwsAttributesInput() *JobTaskForEachTaskTaskNewClusterAwsAttributes
	AzureAttributes() JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference
	AzureAttributesInput() *JobTaskForEachTaskTaskNewClusterAzureAttributes
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() JobTaskForEachTaskTaskNewClusterClusterLogConfOutputReference
	ClusterLogConfInput() *JobTaskForEachTaskTaskNewClusterClusterLogConf
	ClusterMountInfo() JobTaskForEachTaskTaskNewClusterClusterMountInfoList
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
	DockerImage() JobTaskForEachTaskTaskNewClusterDockerImageOutputReference
	DockerImageInput() *JobTaskForEachTaskTaskNewClusterDockerImage
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
	GcpAttributes() JobTaskForEachTaskTaskNewClusterGcpAttributesOutputReference
	GcpAttributesInput() *JobTaskForEachTaskTaskNewClusterGcpAttributes
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	InitScripts() JobTaskForEachTaskTaskNewClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *JobTaskForEachTaskTaskNewCluster
	SetInternalValue(val *JobTaskForEachTaskTaskNewCluster)
	IsSingleNode() interface{}
	SetIsSingleNode(val interface{})
	IsSingleNodeInput() interface{}
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	Library() JobTaskForEachTaskTaskNewClusterLibraryList
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
	WorkloadType() JobTaskForEachTaskTaskNewClusterWorkloadTypeOutputReference
	WorkloadTypeInput() *JobTaskForEachTaskTaskNewClusterWorkloadType
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
	PutAutoscale(value *JobTaskForEachTaskTaskNewClusterAutoscale)
	PutAwsAttributes(value *JobTaskForEachTaskTaskNewClusterAwsAttributes)
	PutAzureAttributes(value *JobTaskForEachTaskTaskNewClusterAzureAttributes)
	PutClusterLogConf(value *JobTaskForEachTaskTaskNewClusterClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *JobTaskForEachTaskTaskNewClusterDockerImage)
	PutGcpAttributes(value *JobTaskForEachTaskTaskNewClusterGcpAttributes)
	PutInitScripts(value interface{})
	PutLibrary(value interface{})
	PutWorkloadType(value *JobTaskForEachTaskTaskNewClusterWorkloadType)
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

// The jsii proxy struct for JobTaskForEachTaskTaskNewClusterOutputReference
type jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) Autoscale() JobTaskForEachTaskTaskNewClusterAutoscaleOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) AutoscaleInput() *JobTaskForEachTaskTaskNewClusterAutoscale {
	var returns *JobTaskForEachTaskTaskNewClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) AwsAttributes() JobTaskForEachTaskTaskNewClusterAwsAttributesOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) AwsAttributesInput() *JobTaskForEachTaskTaskNewClusterAwsAttributes {
	var returns *JobTaskForEachTaskTaskNewClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) AzureAttributes() JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) AzureAttributesInput() *JobTaskForEachTaskTaskNewClusterAzureAttributes {
	var returns *JobTaskForEachTaskTaskNewClusterAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterLogConf() JobTaskForEachTaskTaskNewClusterClusterLogConfOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterLogConfInput() *JobTaskForEachTaskTaskNewClusterClusterLogConf {
	var returns *JobTaskForEachTaskTaskNewClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterMountInfo() JobTaskForEachTaskTaskNewClusterClusterMountInfoList {
	var returns JobTaskForEachTaskTaskNewClusterClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DockerImage() JobTaskForEachTaskTaskNewClusterDockerImageOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DockerImageInput() *JobTaskForEachTaskTaskNewClusterDockerImage {
	var returns *JobTaskForEachTaskTaskNewClusterDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GcpAttributes() JobTaskForEachTaskTaskNewClusterGcpAttributesOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GcpAttributesInput() *JobTaskForEachTaskTaskNewClusterGcpAttributes {
	var returns *JobTaskForEachTaskTaskNewClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InitScripts() JobTaskForEachTaskTaskNewClusterInitScriptsList {
	var returns JobTaskForEachTaskTaskNewClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InternalValue() *JobTaskForEachTaskTaskNewCluster {
	var returns *JobTaskForEachTaskTaskNewCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) IsSingleNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) IsSingleNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) Library() JobTaskForEachTaskTaskNewClusterLibraryList {
	var returns JobTaskForEachTaskTaskNewClusterLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) UseMlRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) UseMlRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) WorkloadType() JobTaskForEachTaskTaskNewClusterWorkloadTypeOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) WorkloadTypeInput() *JobTaskForEachTaskTaskNewClusterWorkloadType {
	var returns *JobTaskForEachTaskTaskNewClusterWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskNewClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobTaskForEachTaskTaskNewClusterOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskNewClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskNewClusterOutputReference_Override(j JobTaskForEachTaskTaskNewClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetInternalValue(val *JobTaskForEachTaskTaskNewCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetIsSingleNode(val interface{}) {
	if err := j.validateSetIsSingleNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleNode",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference)SetUseMlRuntime(val interface{}) {
	if err := j.validateSetUseMlRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMlRuntime",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutAutoscale(value *JobTaskForEachTaskTaskNewClusterAutoscale) {
	if err := j.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutAwsAttributes(value *JobTaskForEachTaskTaskNewClusterAwsAttributes) {
	if err := j.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutAzureAttributes(value *JobTaskForEachTaskTaskNewClusterAzureAttributes) {
	if err := j.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutClusterLogConf(value *JobTaskForEachTaskTaskNewClusterClusterLogConf) {
	if err := j.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutClusterMountInfo(value interface{}) {
	if err := j.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutDockerImage(value *JobTaskForEachTaskTaskNewClusterDockerImage) {
	if err := j.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutGcpAttributes(value *JobTaskForEachTaskTaskNewClusterGcpAttributes) {
	if err := j.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutInitScripts(value interface{}) {
	if err := j.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutLibrary(value interface{}) {
	if err := j.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLibrary",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) PutWorkloadType(value *JobTaskForEachTaskTaskNewClusterWorkloadType) {
	if err := j.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		j,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		j,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		j,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		j,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		j,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		j,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetIsSingleNode() {
	_jsii_.InvokeVoid(
		j,
		"resetIsSingleNode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		j,
		"resetKind",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		j,
		"resetLibrary",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		j,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		j,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		j,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		j,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetSparkVersion() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkVersion",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetUseMlRuntime() {
	_jsii_.InvokeVoid(
		j,
		"resetUseMlRuntime",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

