// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksclusterpluginframework/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterPluginframeworkClusterInfoOutputReference interface {
	cdktf.ComplexObject
	Autoscale() DataDatabricksClusterPluginframeworkClusterInfoAutoscaleList
	AutoscaleInput() interface{}
	AutoterminationMinutes() *float64
	SetAutoterminationMinutes(val *float64)
	AutoterminationMinutesInput() *float64
	AwsAttributes() DataDatabricksClusterPluginframeworkClusterInfoAwsAttributesList
	AwsAttributesInput() interface{}
	AzureAttributes() DataDatabricksClusterPluginframeworkClusterInfoAzureAttributesList
	AzureAttributesInput() interface{}
	ClusterCores() *float64
	SetClusterCores(val *float64)
	ClusterCoresInput() *float64
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() DataDatabricksClusterPluginframeworkClusterInfoClusterLogConfList
	ClusterLogConfInput() interface{}
	ClusterLogStatus() DataDatabricksClusterPluginframeworkClusterInfoClusterLogStatusList
	ClusterLogStatusInput() interface{}
	ClusterMemoryMb() *float64
	SetClusterMemoryMb(val *float64)
	ClusterMemoryMbInput() *float64
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	ClusterSource() *string
	SetClusterSource(val *string)
	ClusterSourceInput() *string
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
	CreatorUserName() *string
	SetCreatorUserName(val *string)
	CreatorUserNameInput() *string
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	DataSecurityMode() *string
	SetDataSecurityMode(val *string)
	DataSecurityModeInput() *string
	DefaultTags() *map[string]*string
	SetDefaultTags(val *map[string]*string)
	DefaultTagsInput() *map[string]*string
	DockerImage() DataDatabricksClusterPluginframeworkClusterInfoDockerImageList
	DockerImageInput() interface{}
	Driver() DataDatabricksClusterPluginframeworkClusterInfoDriverList
	DriverInput() interface{}
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
	Executors() DataDatabricksClusterPluginframeworkClusterInfoExecutorsList
	ExecutorsInput() interface{}
	// Experimental.
	Fqn() *string
	GcpAttributes() DataDatabricksClusterPluginframeworkClusterInfoGcpAttributesList
	GcpAttributesInput() interface{}
	InitScripts() DataDatabricksClusterPluginframeworkClusterInfoInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsSingleNode() interface{}
	SetIsSingleNode(val interface{})
	IsSingleNodeInput() interface{}
	JdbcPort() *float64
	SetJdbcPort(val *float64)
	JdbcPortInput() *float64
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	LastRestartedTime() *float64
	SetLastRestartedTime(val *float64)
	LastRestartedTimeInput() *float64
	LastStateLossTime() *float64
	SetLastStateLossTime(val *float64)
	LastStateLossTimeInput() *float64
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
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
	SparkContextId() *float64
	SetSparkContextId(val *float64)
	SparkContextIdInput() *float64
	SparkEnvVars() *map[string]*string
	SetSparkEnvVars(val *map[string]*string)
	SparkEnvVarsInput() *map[string]*string
	SparkVersion() *string
	SetSparkVersion(val *string)
	SparkVersionInput() *string
	Spec() DataDatabricksClusterPluginframeworkClusterInfoSpecList
	SpecInput() interface{}
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	StartTime() *float64
	SetStartTime(val *float64)
	StartTimeInput() *float64
	State() *string
	SetState(val *string)
	StateInput() *string
	StateMessage() *string
	SetStateMessage(val *string)
	StateMessageInput() *string
	TerminatedTime() *float64
	SetTerminatedTime(val *float64)
	TerminatedTimeInput() *float64
	TerminationReason() DataDatabricksClusterPluginframeworkClusterInfoTerminationReasonList
	TerminationReasonInput() interface{}
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
	WorkloadType() DataDatabricksClusterPluginframeworkClusterInfoWorkloadTypeList
	WorkloadTypeInput() interface{}
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
	PutAutoscale(value interface{})
	PutAwsAttributes(value interface{})
	PutAzureAttributes(value interface{})
	PutClusterLogConf(value interface{})
	PutClusterLogStatus(value interface{})
	PutDockerImage(value interface{})
	PutDriver(value interface{})
	PutExecutors(value interface{})
	PutGcpAttributes(value interface{})
	PutInitScripts(value interface{})
	PutSpec(value interface{})
	PutTerminationReason(value interface{})
	PutWorkloadType(value interface{})
	ResetAutoscale()
	ResetAutoterminationMinutes()
	ResetAwsAttributes()
	ResetAzureAttributes()
	ResetClusterCores()
	ResetClusterId()
	ResetClusterLogConf()
	ResetClusterLogStatus()
	ResetClusterMemoryMb()
	ResetClusterName()
	ResetClusterSource()
	ResetCreatorUserName()
	ResetCustomTags()
	ResetDataSecurityMode()
	ResetDefaultTags()
	ResetDockerImage()
	ResetDriver()
	ResetDriverInstancePoolId()
	ResetDriverNodeTypeId()
	ResetEnableElasticDisk()
	ResetEnableLocalDiskEncryption()
	ResetExecutors()
	ResetGcpAttributes()
	ResetInitScripts()
	ResetInstancePoolId()
	ResetIsSingleNode()
	ResetJdbcPort()
	ResetKind()
	ResetLastRestartedTime()
	ResetLastStateLossTime()
	ResetNodeTypeId()
	ResetNumWorkers()
	ResetPolicyId()
	ResetRemoteDiskThroughput()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkContextId()
	ResetSparkEnvVars()
	ResetSparkVersion()
	ResetSpec()
	ResetSshPublicKeys()
	ResetStartTime()
	ResetState()
	ResetStateMessage()
	ResetTerminatedTime()
	ResetTerminationReason()
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

// The jsii proxy struct for DataDatabricksClusterPluginframeworkClusterInfoOutputReference
type jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Autoscale() DataDatabricksClusterPluginframeworkClusterInfoAutoscaleList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoAutoscaleList
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AutoscaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AwsAttributes() DataDatabricksClusterPluginframeworkClusterInfoAwsAttributesList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoAwsAttributesList
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AwsAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AzureAttributes() DataDatabricksClusterPluginframeworkClusterInfoAzureAttributesList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoAzureAttributesList
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) AzureAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterCoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterCoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterLogConf() DataDatabricksClusterPluginframeworkClusterInfoClusterLogConfList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoClusterLogConfList
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterLogConfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterLogStatus() DataDatabricksClusterPluginframeworkClusterInfoClusterLogStatusList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoClusterLogStatusList
	_jsii_.Get(
		j,
		"clusterLogStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterLogStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterLogStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ClusterSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) CreatorUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) CreatorUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DefaultTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DefaultTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DockerImage() DataDatabricksClusterPluginframeworkClusterInfoDockerImageList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoDockerImageList
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DockerImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Driver() DataDatabricksClusterPluginframeworkClusterInfoDriverList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoDriverList
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DriverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Executors() DataDatabricksClusterPluginframeworkClusterInfoExecutorsList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoExecutorsList
	_jsii_.Get(
		j,
		"executors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ExecutorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GcpAttributes() DataDatabricksClusterPluginframeworkClusterInfoGcpAttributesList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoGcpAttributesList
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GcpAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InitScripts() DataDatabricksClusterPluginframeworkClusterInfoInitScriptsList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) IsSingleNode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) IsSingleNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) JdbcPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jdbcPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) JdbcPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jdbcPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) LastRestartedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastRestartedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) LastRestartedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastRestartedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) LastStateLossTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastStateLossTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) LastStateLossTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastStateLossTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) RemoteDiskThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) RemoteDiskThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"remoteDiskThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkContextId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sparkContextId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkContextIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sparkContextIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Spec() DataDatabricksClusterPluginframeworkClusterInfoSpecList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoSpecList
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) StartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) StartTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) StateMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerminatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerminatedTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminatedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerminationReason() DataDatabricksClusterPluginframeworkClusterInfoTerminationReasonList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoTerminationReasonList
	_jsii_.Get(
		j,
		"terminationReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerminationReasonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TotalInitialRemoteDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) TotalInitialRemoteDiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInitialRemoteDiskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) UseMlRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) UseMlRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMlRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) WorkloadType() DataDatabricksClusterPluginframeworkClusterInfoWorkloadTypeList {
	var returns DataDatabricksClusterPluginframeworkClusterInfoWorkloadTypeList
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) WorkloadTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterPluginframeworkClusterInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksClusterPluginframeworkClusterInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterPluginframeworkClusterInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksClusterPluginframework.DataDatabricksClusterPluginframeworkClusterInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterPluginframeworkClusterInfoOutputReference_Override(d DataDatabricksClusterPluginframeworkClusterInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksClusterPluginframework.DataDatabricksClusterPluginframeworkClusterInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetAutoterminationMinutes(val *float64) {
	if err := j.validateSetAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetClusterCores(val *float64) {
	if err := j.validateSetClusterCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterCores",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetClusterMemoryMb(val *float64) {
	if err := j.validateSetClusterMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMemoryMb",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetClusterSource(val *string) {
	if err := j.validateSetClusterSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetCreatorUserName(val *string) {
	if err := j.validateSetCreatorUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creatorUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetDefaultTags(val *map[string]*string) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetIsSingleNode(val interface{}) {
	if err := j.validateSetIsSingleNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleNode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetJdbcPort(val *float64) {
	if err := j.validateSetJdbcPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcPort",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetLastRestartedTime(val *float64) {
	if err := j.validateSetLastRestartedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastRestartedTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetLastStateLossTime(val *float64) {
	if err := j.validateSetLastStateLossTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastStateLossTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetRemoteDiskThroughput(val *float64) {
	if err := j.validateSetRemoteDiskThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteDiskThroughput",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSparkContextId(val *float64) {
	if err := j.validateSetSparkContextIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkContextId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetStartTime(val *float64) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetStateMessage(val *string) {
	if err := j.validateSetStateMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateMessage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetTerminatedTime(val *float64) {
	if err := j.validateSetTerminatedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminatedTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetTotalInitialRemoteDiskSize(val *float64) {
	if err := j.validateSetTotalInitialRemoteDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalInitialRemoteDiskSize",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference)SetUseMlRuntime(val interface{}) {
	if err := j.validateSetUseMlRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMlRuntime",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutAutoscale(value interface{}) {
	if err := d.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutAwsAttributes(value interface{}) {
	if err := d.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutAzureAttributes(value interface{}) {
	if err := d.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutClusterLogConf(value interface{}) {
	if err := d.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutClusterLogStatus(value interface{}) {
	if err := d.validatePutClusterLogStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterLogStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutDockerImage(value interface{}) {
	if err := d.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutDriver(value interface{}) {
	if err := d.validatePutDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDriver",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutExecutors(value interface{}) {
	if err := d.validatePutExecutorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExecutors",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutGcpAttributes(value interface{}) {
	if err := d.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutInitScripts(value interface{}) {
	if err := d.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutSpec(value interface{}) {
	if err := d.validatePutSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutTerminationReason(value interface{}) {
	if err := d.validatePutTerminationReasonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTerminationReason",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) PutWorkloadType(value interface{}) {
	if err := d.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetAutoterminationMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoterminationMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterCores() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterCores",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterLogStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterLogStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterMemoryMb() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterMemoryMb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetClusterSource() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetCreatorUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatorUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDriver() {
	_jsii_.InvokeVoid(
		d,
		"resetDriver",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetExecutors() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		d,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetIsSingleNode() {
	_jsii_.InvokeVoid(
		d,
		"resetIsSingleNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetJdbcPort() {
	_jsii_.InvokeVoid(
		d,
		"resetJdbcPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		d,
		"resetKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetLastRestartedTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLastRestartedTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetLastStateLossTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLastStateLossTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetRemoteDiskThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetRemoteDiskThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSparkContextId() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkContextId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSparkVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		d,
		"resetStartTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetStateMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetStateMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetTerminatedTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminatedTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetTerminationReason() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminationReason",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetTotalInitialRemoteDiskSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTotalInitialRemoteDiskSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetUseMlRuntime() {
	_jsii_.InvokeVoid(
		d,
		"resetUseMlRuntime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterPluginframeworkClusterInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

