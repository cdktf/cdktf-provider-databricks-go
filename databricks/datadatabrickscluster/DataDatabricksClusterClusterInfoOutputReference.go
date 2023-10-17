// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v12/datadatabrickscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterClusterInfoOutputReference interface {
	cdktf.ComplexObject
	Autoscale() DataDatabricksClusterClusterInfoAutoscaleOutputReference
	AutoscaleInput() *DataDatabricksClusterClusterInfoAutoscale
	AutoterminationMinutes() *float64
	SetAutoterminationMinutes(val *float64)
	AutoterminationMinutesInput() *float64
	AwsAttributes() DataDatabricksClusterClusterInfoAwsAttributesOutputReference
	AwsAttributesInput() *DataDatabricksClusterClusterInfoAwsAttributes
	AzureAttributes() DataDatabricksClusterClusterInfoAzureAttributesOutputReference
	AzureAttributesInput() *DataDatabricksClusterClusterInfoAzureAttributes
	ClusterCores() *float64
	SetClusterCores(val *float64)
	ClusterCoresInput() *float64
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() DataDatabricksClusterClusterInfoClusterLogConfOutputReference
	ClusterLogConfInput() *DataDatabricksClusterClusterInfoClusterLogConf
	ClusterLogStatus() DataDatabricksClusterClusterInfoClusterLogStatusOutputReference
	ClusterLogStatusInput() *DataDatabricksClusterClusterInfoClusterLogStatus
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
	DockerImage() DataDatabricksClusterClusterInfoDockerImageOutputReference
	DockerImageInput() *DataDatabricksClusterClusterInfoDockerImage
	Driver() DataDatabricksClusterClusterInfoDriverOutputReference
	DriverInput() *DataDatabricksClusterClusterInfoDriver
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
	Executors() DataDatabricksClusterClusterInfoExecutorsList
	ExecutorsInput() interface{}
	// Experimental.
	Fqn() *string
	GcpAttributes() DataDatabricksClusterClusterInfoGcpAttributesOutputReference
	GcpAttributesInput() *DataDatabricksClusterClusterInfoGcpAttributes
	InitScripts() DataDatabricksClusterClusterInfoInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *DataDatabricksClusterClusterInfo
	SetInternalValue(val *DataDatabricksClusterClusterInfo)
	JdbcPort() *float64
	SetJdbcPort(val *float64)
	JdbcPortInput() *float64
	LastActivityTime() *float64
	SetLastActivityTime(val *float64)
	LastActivityTimeInput() *float64
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
	TerminateTime() *float64
	SetTerminateTime(val *float64)
	TerminateTimeInput() *float64
	TerminationReason() DataDatabricksClusterClusterInfoTerminationReasonOutputReference
	TerminationReasonInput() *DataDatabricksClusterClusterInfoTerminationReason
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
	PutAutoscale(value *DataDatabricksClusterClusterInfoAutoscale)
	PutAwsAttributes(value *DataDatabricksClusterClusterInfoAwsAttributes)
	PutAzureAttributes(value *DataDatabricksClusterClusterInfoAzureAttributes)
	PutClusterLogConf(value *DataDatabricksClusterClusterInfoClusterLogConf)
	PutClusterLogStatus(value *DataDatabricksClusterClusterInfoClusterLogStatus)
	PutDockerImage(value *DataDatabricksClusterClusterInfoDockerImage)
	PutDriver(value *DataDatabricksClusterClusterInfoDriver)
	PutExecutors(value interface{})
	PutGcpAttributes(value *DataDatabricksClusterClusterInfoGcpAttributes)
	PutInitScripts(value interface{})
	PutTerminationReason(value *DataDatabricksClusterClusterInfoTerminationReason)
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
	ResetJdbcPort()
	ResetLastActivityTime()
	ResetLastStateLossTime()
	ResetNodeTypeId()
	ResetNumWorkers()
	ResetPolicyId()
	ResetRuntimeEngine()
	ResetSingleUserName()
	ResetSparkConf()
	ResetSparkContextId()
	ResetSparkEnvVars()
	ResetSshPublicKeys()
	ResetStartTime()
	ResetStateMessage()
	ResetTerminateTime()
	ResetTerminationReason()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksClusterClusterInfoOutputReference
type jsiiProxy_DataDatabricksClusterClusterInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) Autoscale() DataDatabricksClusterClusterInfoAutoscaleOutputReference {
	var returns DataDatabricksClusterClusterInfoAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AutoscaleInput() *DataDatabricksClusterClusterInfoAutoscale {
	var returns *DataDatabricksClusterClusterInfoAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AwsAttributes() DataDatabricksClusterClusterInfoAwsAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AwsAttributesInput() *DataDatabricksClusterClusterInfoAwsAttributes {
	var returns *DataDatabricksClusterClusterInfoAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AzureAttributes() DataDatabricksClusterClusterInfoAzureAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) AzureAttributesInput() *DataDatabricksClusterClusterInfoAzureAttributes {
	var returns *DataDatabricksClusterClusterInfoAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterCores() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterCoresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterCoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterLogConf() DataDatabricksClusterClusterInfoClusterLogConfOutputReference {
	var returns DataDatabricksClusterClusterInfoClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterLogConfInput() *DataDatabricksClusterClusterInfoClusterLogConf {
	var returns *DataDatabricksClusterClusterInfoClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterLogStatus() DataDatabricksClusterClusterInfoClusterLogStatusOutputReference {
	var returns DataDatabricksClusterClusterInfoClusterLogStatusOutputReference
	_jsii_.Get(
		j,
		"clusterLogStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterLogStatusInput() *DataDatabricksClusterClusterInfoClusterLogStatus {
	var returns *DataDatabricksClusterClusterInfoClusterLogStatus
	_jsii_.Get(
		j,
		"clusterLogStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ClusterSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) CreatorUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) CreatorUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DefaultTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DefaultTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DockerImage() DataDatabricksClusterClusterInfoDockerImageOutputReference {
	var returns DataDatabricksClusterClusterInfoDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DockerImageInput() *DataDatabricksClusterClusterInfoDockerImage {
	var returns *DataDatabricksClusterClusterInfoDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) Driver() DataDatabricksClusterClusterInfoDriverOutputReference {
	var returns DataDatabricksClusterClusterInfoDriverOutputReference
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DriverInput() *DataDatabricksClusterClusterInfoDriver {
	var returns *DataDatabricksClusterClusterInfoDriver
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) Executors() DataDatabricksClusterClusterInfoExecutorsList {
	var returns DataDatabricksClusterClusterInfoExecutorsList
	_jsii_.Get(
		j,
		"executors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ExecutorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GcpAttributes() DataDatabricksClusterClusterInfoGcpAttributesOutputReference {
	var returns DataDatabricksClusterClusterInfoGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GcpAttributesInput() *DataDatabricksClusterClusterInfoGcpAttributes {
	var returns *DataDatabricksClusterClusterInfoGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InitScripts() DataDatabricksClusterClusterInfoInitScriptsList {
	var returns DataDatabricksClusterClusterInfoInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InternalValue() *DataDatabricksClusterClusterInfo {
	var returns *DataDatabricksClusterClusterInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) JdbcPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jdbcPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) JdbcPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jdbcPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) LastActivityTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastActivityTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) LastActivityTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastActivityTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) LastStateLossTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastStateLossTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) LastStateLossTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastStateLossTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkContextId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sparkContextId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkContextIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sparkContextIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) StartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) StartTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) StateMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerminateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerminateTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerminationReason() DataDatabricksClusterClusterInfoTerminationReasonOutputReference {
	var returns DataDatabricksClusterClusterInfoTerminationReasonOutputReference
	_jsii_.Get(
		j,
		"terminationReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerminationReasonInput() *DataDatabricksClusterClusterInfoTerminationReason {
	var returns *DataDatabricksClusterClusterInfoTerminationReason
	_jsii_.Get(
		j,
		"terminationReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterClusterInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksClusterClusterInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterClusterInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterClusterInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterClusterInfoOutputReference_Override(d DataDatabricksClusterClusterInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetAutoterminationMinutes(val *float64) {
	if err := j.validateSetAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetClusterCores(val *float64) {
	if err := j.validateSetClusterCoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterCores",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetClusterMemoryMb(val *float64) {
	if err := j.validateSetClusterMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMemoryMb",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetClusterSource(val *string) {
	if err := j.validateSetClusterSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetCreatorUserName(val *string) {
	if err := j.validateSetCreatorUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creatorUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetDefaultTags(val *map[string]*string) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetInternalValue(val *DataDatabricksClusterClusterInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetJdbcPort(val *float64) {
	if err := j.validateSetJdbcPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jdbcPort",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetLastActivityTime(val *float64) {
	if err := j.validateSetLastActivityTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastActivityTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetLastStateLossTime(val *float64) {
	if err := j.validateSetLastStateLossTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastStateLossTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSparkContextId(val *float64) {
	if err := j.validateSetSparkContextIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkContextId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetStartTime(val *float64) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetStateMessage(val *string) {
	if err := j.validateSetStateMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateMessage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetTerminateTime(val *float64) {
	if err := j.validateSetTerminateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateTime",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutAutoscale(value *DataDatabricksClusterClusterInfoAutoscale) {
	if err := d.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutAwsAttributes(value *DataDatabricksClusterClusterInfoAwsAttributes) {
	if err := d.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutAzureAttributes(value *DataDatabricksClusterClusterInfoAzureAttributes) {
	if err := d.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutClusterLogConf(value *DataDatabricksClusterClusterInfoClusterLogConf) {
	if err := d.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutClusterLogStatus(value *DataDatabricksClusterClusterInfoClusterLogStatus) {
	if err := d.validatePutClusterLogStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClusterLogStatus",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutDockerImage(value *DataDatabricksClusterClusterInfoDockerImage) {
	if err := d.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutDriver(value *DataDatabricksClusterClusterInfoDriver) {
	if err := d.validatePutDriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDriver",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutExecutors(value interface{}) {
	if err := d.validatePutExecutorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExecutors",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutGcpAttributes(value *DataDatabricksClusterClusterInfoGcpAttributes) {
	if err := d.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutInitScripts(value interface{}) {
	if err := d.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) PutTerminationReason(value *DataDatabricksClusterClusterInfoTerminationReason) {
	if err := d.validatePutTerminationReasonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTerminationReason",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetAutoterminationMinutes() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoterminationMinutes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterCores() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterCores",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterLogStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterLogStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterMemoryMb() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterMemoryMb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetClusterSource() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetCreatorUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatorUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetDriver() {
	_jsii_.InvokeVoid(
		d,
		"resetDriver",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetExecutors() {
	_jsii_.InvokeVoid(
		d,
		"resetExecutors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		d,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetJdbcPort() {
	_jsii_.InvokeVoid(
		d,
		"resetJdbcPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetLastActivityTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLastActivityTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetLastStateLossTime() {
	_jsii_.InvokeVoid(
		d,
		"resetLastStateLossTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetSparkContextId() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkContextId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		d,
		"resetStartTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetStateMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetStateMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetTerminateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ResetTerminationReason() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminationReason",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

