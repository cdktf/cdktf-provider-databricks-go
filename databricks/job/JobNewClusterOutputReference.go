// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v10/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobNewClusterOutputReference interface {
	cdktf.ComplexObject
	ApplyPolicyDefaultValues() interface{}
	SetApplyPolicyDefaultValues(val interface{})
	ApplyPolicyDefaultValuesInput() interface{}
	Autoscale() JobNewClusterAutoscaleOutputReference
	AutoscaleInput() *JobNewClusterAutoscale
	AutoterminationMinutes() *float64
	SetAutoterminationMinutes(val *float64)
	AutoterminationMinutesInput() *float64
	AwsAttributes() JobNewClusterAwsAttributesOutputReference
	AwsAttributesInput() *JobNewClusterAwsAttributes
	AzureAttributes() JobNewClusterAzureAttributesOutputReference
	AzureAttributesInput() *JobNewClusterAzureAttributes
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterLogConf() JobNewClusterClusterLogConfOutputReference
	ClusterLogConfInput() *JobNewClusterClusterLogConf
	ClusterMountInfo() JobNewClusterClusterMountInfoList
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
	DockerImage() JobNewClusterDockerImageOutputReference
	DockerImageInput() *JobNewClusterDockerImage
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
	GcpAttributes() JobNewClusterGcpAttributesOutputReference
	GcpAttributesInput() *JobNewClusterGcpAttributes
	IdempotencyToken() *string
	SetIdempotencyToken(val *string)
	IdempotencyTokenInput() *string
	InitScripts() JobNewClusterInitScriptsList
	InitScriptsInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InternalValue() *JobNewCluster
	SetInternalValue(val *JobNewCluster)
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
	WorkloadType() JobNewClusterWorkloadTypeOutputReference
	WorkloadTypeInput() *JobNewClusterWorkloadType
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
	PutAutoscale(value *JobNewClusterAutoscale)
	PutAwsAttributes(value *JobNewClusterAwsAttributes)
	PutAzureAttributes(value *JobNewClusterAzureAttributes)
	PutClusterLogConf(value *JobNewClusterClusterLogConf)
	PutClusterMountInfo(value interface{})
	PutDockerImage(value *JobNewClusterDockerImage)
	PutGcpAttributes(value *JobNewClusterGcpAttributes)
	PutInitScripts(value interface{})
	PutWorkloadType(value *JobNewClusterWorkloadType)
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

// The jsii proxy struct for JobNewClusterOutputReference
type jsiiProxy_JobNewClusterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobNewClusterOutputReference) ApplyPolicyDefaultValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ApplyPolicyDefaultValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyPolicyDefaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) Autoscale() JobNewClusterAutoscaleOutputReference {
	var returns JobNewClusterAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AutoscaleInput() *JobNewClusterAutoscale {
	var returns *JobNewClusterAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AwsAttributes() JobNewClusterAwsAttributesOutputReference {
	var returns JobNewClusterAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AwsAttributesInput() *JobNewClusterAwsAttributes {
	var returns *JobNewClusterAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AzureAttributes() JobNewClusterAzureAttributesOutputReference {
	var returns JobNewClusterAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) AzureAttributesInput() *JobNewClusterAzureAttributes {
	var returns *JobNewClusterAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterLogConf() JobNewClusterClusterLogConfOutputReference {
	var returns JobNewClusterClusterLogConfOutputReference
	_jsii_.Get(
		j,
		"clusterLogConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterLogConfInput() *JobNewClusterClusterLogConf {
	var returns *JobNewClusterClusterLogConf
	_jsii_.Get(
		j,
		"clusterLogConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterMountInfo() JobNewClusterClusterMountInfoList {
	var returns JobNewClusterClusterMountInfoList
	_jsii_.Get(
		j,
		"clusterMountInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterMountInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterMountInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DataSecurityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DataSecurityModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DockerImage() JobNewClusterDockerImageOutputReference {
	var returns JobNewClusterDockerImageOutputReference
	_jsii_.Get(
		j,
		"dockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DockerImageInput() *JobNewClusterDockerImage {
	var returns *JobNewClusterDockerImage
	_jsii_.Get(
		j,
		"dockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DriverInstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DriverInstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInstancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DriverNodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) DriverNodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverNodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) EnableLocalDiskEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) EnableLocalDiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalDiskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) GcpAttributes() JobNewClusterGcpAttributesOutputReference {
	var returns JobNewClusterGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) GcpAttributesInput() *JobNewClusterGcpAttributes {
	var returns *JobNewClusterGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) IdempotencyToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) IdempotencyTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idempotencyTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InitScripts() JobNewClusterInitScriptsList {
	var returns JobNewClusterInitScriptsList
	_jsii_.Get(
		j,
		"initScripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InitScriptsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initScriptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InternalValue() *JobNewCluster {
	var returns *JobNewCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) RuntimeEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) RuntimeEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SingleUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SingleUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkConf() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkConfInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkEnvVars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkEnvVarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SparkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) WorkloadType() JobNewClusterWorkloadTypeOutputReference {
	var returns JobNewClusterWorkloadTypeOutputReference
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) WorkloadTypeInput() *JobNewClusterWorkloadType {
	var returns *JobNewClusterWorkloadType
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewJobNewClusterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobNewClusterOutputReference {
	_init_.Initialize()

	if err := validateNewJobNewClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobNewClusterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobNewClusterOutputReference_Override(j JobNewClusterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobNewClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetApplyPolicyDefaultValues(val interface{}) {
	if err := j.validateSetApplyPolicyDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyPolicyDefaultValues",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetAutoterminationMinutes(val *float64) {
	if err := j.validateSetAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetDataSecurityMode(val *string) {
	if err := j.validateSetDataSecurityModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSecurityMode",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetDriverInstancePoolId(val *string) {
	if err := j.validateSetDriverInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverInstancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetDriverNodeTypeId(val *string) {
	if err := j.validateSetDriverNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverNodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetEnableLocalDiskEncryption(val interface{}) {
	if err := j.validateSetEnableLocalDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalDiskEncryption",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetIdempotencyToken(val *string) {
	if err := j.validateSetIdempotencyTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idempotencyToken",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetInternalValue(val *JobNewCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetRuntimeEngine(val *string) {
	if err := j.validateSetRuntimeEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEngine",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetSingleUserName(val *string) {
	if err := j.validateSetSingleUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleUserName",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetSparkConf(val *map[string]*string) {
	if err := j.validateSetSparkConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkConf",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetSparkEnvVars(val *map[string]*string) {
	if err := j.validateSetSparkEnvVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkEnvVars",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetSparkVersion(val *string) {
	if err := j.validateSetSparkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkVersion",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobNewClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobNewClusterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobNewClusterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobNewClusterOutputReference) PutAutoscale(value *JobNewClusterAutoscale) {
	if err := j.validatePutAutoscaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutAwsAttributes(value *JobNewClusterAwsAttributes) {
	if err := j.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutAzureAttributes(value *JobNewClusterAzureAttributes) {
	if err := j.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutClusterLogConf(value *JobNewClusterClusterLogConf) {
	if err := j.validatePutClusterLogConfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterLogConf",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutClusterMountInfo(value interface{}) {
	if err := j.validatePutClusterMountInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putClusterMountInfo",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutDockerImage(value *JobNewClusterDockerImage) {
	if err := j.validatePutDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDockerImage",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutGcpAttributes(value *JobNewClusterGcpAttributes) {
	if err := j.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutInitScripts(value interface{}) {
	if err := j.validatePutInitScriptsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitScripts",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) PutWorkloadType(value *JobNewClusterWorkloadType) {
	if err := j.validatePutWorkloadTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putWorkloadType",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetApplyPolicyDefaultValues() {
	_jsii_.InvokeVoid(
		j,
		"resetApplyPolicyDefaultValues",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetAutoterminationMinutes() {
	_jsii_.InvokeVoid(
		j,
		"resetAutoterminationMinutes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetClusterLogConf() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterLogConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetClusterMountInfo() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterMountInfo",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		j,
		"resetClusterName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		j,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetDataSecurityMode() {
	_jsii_.InvokeVoid(
		j,
		"resetDataSecurityMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetDockerImage() {
	_jsii_.InvokeVoid(
		j,
		"resetDockerImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetDriverInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetDriverNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetDriverNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetEnableLocalDiskEncryption() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableLocalDiskEncryption",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		j,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetIdempotencyToken() {
	_jsii_.InvokeVoid(
		j,
		"resetIdempotencyToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetInitScripts() {
	_jsii_.InvokeVoid(
		j,
		"resetInitScripts",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		j,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		j,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetPolicyId() {
	_jsii_.InvokeVoid(
		j,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetRuntimeEngine() {
	_jsii_.InvokeVoid(
		j,
		"resetRuntimeEngine",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetSingleUserName() {
	_jsii_.InvokeVoid(
		j,
		"resetSingleUserName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetSparkConf() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkConf",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetSparkEnvVars() {
	_jsii_.InvokeVoid(
		j,
		"resetSparkEnvVars",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobNewClusterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobNewClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

