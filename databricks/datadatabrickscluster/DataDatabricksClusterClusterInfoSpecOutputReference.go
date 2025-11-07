// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabrickscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterClusterInfoSpecOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() DataDatabricksClusterClusterInfoSpecAutoscaleOutputReference
	AutoscaleInput() *DataDatabricksClusterClusterInfoSpecAutoscale
	AwsAttributes() DataDatabricksClusterClusterInfoSpecAwsAttributesOutputReference
	AwsAttributesInput() *DataDatabricksClusterClusterInfoSpecAwsAttributes
	AzureAttributes() DataDatabricksClusterClusterInfoSpecAzureAttributesOutputReference
	AzureAttributesInput() *DataDatabricksClusterClusterInfoSpecAzureAttributes
	ClusterId() *string
	ClusterLogConf() DataDatabricksClusterClusterInfoSpecClusterLogConfOutputReference
	ClusterLogConfInput() *DataDatabricksClusterClusterInfoSpecClusterLogConf
	ClusterMountInfo() DataDatabricksClusterClusterInfoSpecClusterMountInfoList
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
	DockerImage() DataDatabricksClusterClusterInfoSpecDockerImageOutputReference
	DockerImageInput() *DataDatabricksClusterClusterInfoSpecDockerImage
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
	GcpAttributes() DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference
	GcpAttributesInput() *DataDatabricksClusterClusterInfoSpecGcpAttributes
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	InitScripts() DataDatabricksClusterClusterInfoSpecInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *DataDatabricksClusterClusterInfoSpec
	SetInternalValue(val *DataDatabricksClusterClusterInfoSpec)
	IsSingleNode() interface{}
	SetIsSingleNode(val interface{})
	IsSingleNodeInput() interface{}
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	Library() DataDatabricksClusterClusterInfoSpecLibraryList
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
	ProviderConfig() DataDatabricksClusterClusterInfoSpecProviderConfigOutputReference
	ProviderConfigInput() *DataDatabricksClusterClusterInfoSpecProviderConfig
	RemoteDiskThroughput() *float64
	SetRemoteDiskThroughput(val *float64)
	RemoteDiskThroughputInput() *float64
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
	TotalInitialRemoteDiskSize() *float64
	SetTotalInitialRemoteDiskSize(val *float64)
	TotalInitialRemoteDiskSizeInput() *float64
	UseMlRuntime() interface{}
	SetUseMlRuntime(val interface{})
	UseMlRuntimeInput() interface{}
	WorkloadType() DataDatabricksClusterClusterInfoSpecWorkloadTypeOutputReference
	WorkloadTypeInput() *DataDatabricksClusterClusterInfoSpecWorkloadType
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
	PutAutoscale(value *DataDatabricksClusterClusterInfoSpecAutoscale)
	PutAwsAttributes(value *DataDatabricksClusterClusterInfoSpecAwsAttributes)
	PutAzureAttributes(value *DataDatabricksClusterClusterInfoSpecAzureAttributes)
	PutClusterLogConf(value *DataDatabricksClusterClusterInfoSpecClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *DataDatabricksClusterClusterInfoSpecDockerImage)
	PutGcpAttributes(value *DataDatabricksClusterClusterInfoSpecGcpAttributes)
	PutInitScripts(value interface{})
	PutLibrary(value interface{})
	PutProviderConfig(value *DataDatabricksClusterClusterInfoSpecProviderConfig)
	PutWorkloadType(value *DataDatabricksClusterClusterInfoSpecWorkloadType)
	ResetApplyPolicyDefaultValues()
	ResetAutoscale()
	ResetAwsAttributes()
	ResetAzureAttributes()
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
	ResetProviderConfig()
	ResetRemoteDiskThroughput()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkEnvVars()
	ResetSparkVersion()
	ResetSshPublicKeys()
	ResetTotalInitialRemoteDiskSize()
	ResetUseMlRuntime()
	ResetWorkloadType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksClusterClusterInfoSpecOutputReference
type jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) Autoscale() DataDatabricksClusterClusterInfoSpecAutoscaleOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) AutoscaleInput() *DataDatabricksClusterClusterInfoSpecAutoscale {
	var returns *DataDatabricksClusterClusterInfoSpecAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) AwsAttributes() DataDatabricksClusterClusterInfoSpecAwsAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) AwsAttributesInput() *DataDatabricksClusterClusterInfoSpecAwsAttributes {
	var returns *DataDatabricksClusterClusterInfoSpecAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) AzureAttributes() DataDatabricksClusterClusterInfoSpecAzureAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) AzureAttributesInput() *DataDatabricksClusterClusterInfoSpecAzureAttributes {
	var returns *DataDatabricksClusterClusterInfoSpecAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterLogConf() DataDatabricksClusterClusterInfoSpecClusterLogConfOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterLogConfInput() *DataDatabricksClusterClusterInfoSpecClusterLogConf {
	var returns *DataDatabricksClusterClusterInfoSpecClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterMountInfo() DataDatabricksClusterClusterInfoSpecClusterMountInfoList {
	var returns DataDatabricksClusterClusterInfoSpecClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DockerImage() DataDatabricksClusterClusterInfoSpecDockerImageOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DockerImageInput() *DataDatabricksClusterClusterInfoSpecDockerImage {
	var returns *DataDatabricksClusterClusterInfoSpecDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GcpAttributes() DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GcpAttributesInput() *DataDatabricksClusterClusterInfoSpecGcpAttributes {
	var returns *DataDatabricksClusterClusterInfoSpecGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InitScripts() DataDatabricksClusterClusterInfoSpecInitScriptsList {
	var returns DataDatabricksClusterClusterInfoSpecInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InternalValue() *DataDatabricksClusterClusterInfoSpec {
	var returns *DataDatabricksClusterClusterInfoSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) IsSingleNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) IsSingleNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) Library() DataDatabricksClusterClusterInfoSpecLibraryList {
	var returns DataDatabricksClusterClusterInfoSpecLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ProviderConfig() DataDatabricksClusterClusterInfoSpecProviderConfigOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecProviderConfigOutputReference
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ProviderConfigInput() *DataDatabricksClusterClusterInfoSpecProviderConfig {
	var returns *DataDatabricksClusterClusterInfoSpecProviderConfig
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) RemoteDiskThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) RemoteDiskThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) TotalInitialRemoteDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) TotalInitialRemoteDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) UseMlRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) UseMlRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) WorkloadType() DataDatabricksClusterClusterInfoSpecWorkloadTypeOutputReference {
	var returns DataDatabricksClusterClusterInfoSpecWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) WorkloadTypeInput() *DataDatabricksClusterClusterInfoSpecWorkloadType {
	var returns *DataDatabricksClusterClusterInfoSpecWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterClusterInfoSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksClusterClusterInfoSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterClusterInfoSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterClusterInfoSpecOutputReference_Override(d DataDatabricksClusterClusterInfoSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetInternalValue(val *DataDatabricksClusterClusterInfoSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetIsSingleNode(val interface{}) {
	if err := j.validateSetIsSingleNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleNode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetRemoteDiskThroughput(val *float64) {
	if err := j.validateSetRemoteDiskThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDiskThroughput",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetTotalInitialRemoteDiskSize(val *float64) {
	if err := j.validateSetTotalInitialRemoteDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalInitialRemoteDiskSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference)SetUseMlRuntime(val interface{}) {
	if err := j.validateSetUseMlRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMlRuntime",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutAutoscale(value *DataDatabricksClusterClusterInfoSpecAutoscale) {
	if err := d.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutAwsAttributes(value *DataDatabricksClusterClusterInfoSpecAwsAttributes) {
	if err := d.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutAzureAttributes(value *DataDatabricksClusterClusterInfoSpecAzureAttributes) {
	if err := d.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutClusterLogConf(value *DataDatabricksClusterClusterInfoSpecClusterLogConf) {
	if err := d.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutClusterMountInfo(value interface{}) {
	if err := d.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutDockerImage(value *DataDatabricksClusterClusterInfoSpecDockerImage) {
	if err := d.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutGcpAttributes(value *DataDatabricksClusterClusterInfoSpecGcpAttributes) {
	if err := d.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutInitScripts(value interface{}) {
	if err := d.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutLibrary(value interface{}) {
	if err := d.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLibrary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutProviderConfig(value *DataDatabricksClusterClusterInfoSpecProviderConfig) {
	if err := d.validatePutProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProviderConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) PutWorkloadType(value *DataDatabricksClusterClusterInfoSpecWorkloadType) {
	if err := d.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		d,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		d,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		d,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetIsSingleNode() {
	_jsii_.InvokeVoid(
		d,
		"resetIsSingleNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		d,
		"resetKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetLibrary() {
	_jsii_.InvokeVoid(
		d,
		"resetLibrary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetRemoteDiskThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoteDiskThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetSparkVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetTotalInitialRemoteDiskSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTotalInitialRemoteDiskSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetUseMlRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetUseMlRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

