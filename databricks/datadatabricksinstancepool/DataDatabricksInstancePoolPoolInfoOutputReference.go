package datadatabricksinstancepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/datadatabricksinstancepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksInstancePoolPoolInfoOutputReference interface {
	cdktf.ComplexObject
	AwsAttributes() DataDatabricksInstancePoolPoolInfoAwsAttributesOutputReference
	AwsAttributesInput() *DataDatabricksInstancePoolPoolInfoAwsAttributes
	AzureAttributes() DataDatabricksInstancePoolPoolInfoAzureAttributesOutputReference
	AzureAttributesInput() *DataDatabricksInstancePoolPoolInfoAzureAttributes
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
	DefaultTags() *map[string]*string
	SetDefaultTags(val *map[string]*string)
	DefaultTagsInput() *map[string]*string
	DiskSpec() DataDatabricksInstancePoolPoolInfoDiskSpecOutputReference
	DiskSpecInput() *DataDatabricksInstancePoolPoolInfoDiskSpec
	EnableElasticDisk() interface{}
	SetEnableElasticDisk(val interface{})
	EnableElasticDiskInput() interface{}
	// Experimental.
	Fqn() *string
	GcpAttributes() DataDatabricksInstancePoolPoolInfoGcpAttributesOutputReference
	GcpAttributesInput() *DataDatabricksInstancePoolPoolInfoGcpAttributes
	IdleInstanceAutoterminationMinutes() *float64
	SetIdleInstanceAutoterminationMinutes(val *float64)
	IdleInstanceAutoterminationMinutesInput() *float64
	InstancePoolFleetAttributes() DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesList
	InstancePoolFleetAttributesInput() interface{}
	InstancePoolId() *string
	SetInstancePoolId(val *string)
	InstancePoolIdInput() *string
	InstancePoolName() *string
	SetInstancePoolName(val *string)
	InstancePoolNameInput() *string
	InternalValue() *DataDatabricksInstancePoolPoolInfo
	SetInternalValue(val *DataDatabricksInstancePoolPoolInfo)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	PreloadedDockerImage() DataDatabricksInstancePoolPoolInfoPreloadedDockerImageList
	PreloadedDockerImageInput() interface{}
	PreloadedSparkVersions() *[]*string
	SetPreloadedSparkVersions(val *[]*string)
	PreloadedSparkVersionsInput() *[]*string
	State() *string
	SetState(val *string)
	StateInput() *string
	Stats() DataDatabricksInstancePoolPoolInfoStatsOutputReference
	StatsInput() *DataDatabricksInstancePoolPoolInfoStats
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
	PutAwsAttributes(value *DataDatabricksInstancePoolPoolInfoAwsAttributes)
	PutAzureAttributes(value *DataDatabricksInstancePoolPoolInfoAzureAttributes)
	PutDiskSpec(value *DataDatabricksInstancePoolPoolInfoDiskSpec)
	PutGcpAttributes(value *DataDatabricksInstancePoolPoolInfoGcpAttributes)
	PutInstancePoolFleetAttributes(value interface{})
	PutPreloadedDockerImage(value interface{})
	PutStats(value *DataDatabricksInstancePoolPoolInfoStats)
	ResetAwsAttributes()
	ResetAzureAttributes()
	ResetCustomTags()
	ResetDefaultTags()
	ResetDiskSpec()
	ResetEnableElasticDisk()
	ResetGcpAttributes()
	ResetInstancePoolFleetAttributes()
	ResetInstancePoolId()
	ResetMaxCapacity()
	ResetMinIdleInstances()
	ResetNodeTypeId()
	ResetPreloadedDockerImage()
	ResetPreloadedSparkVersions()
	ResetState()
	ResetStats()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksInstancePoolPoolInfoOutputReference
type jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) AwsAttributes() DataDatabricksInstancePoolPoolInfoAwsAttributesOutputReference {
	var returns DataDatabricksInstancePoolPoolInfoAwsAttributesOutputReference
	_jsii_.Get(
		j,
		"awsAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) AwsAttributesInput() *DataDatabricksInstancePoolPoolInfoAwsAttributes {
	var returns *DataDatabricksInstancePoolPoolInfoAwsAttributes
	_jsii_.Get(
		j,
		"awsAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) AzureAttributes() DataDatabricksInstancePoolPoolInfoAzureAttributesOutputReference {
	var returns DataDatabricksInstancePoolPoolInfoAzureAttributesOutputReference
	_jsii_.Get(
		j,
		"azureAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) AzureAttributesInput() *DataDatabricksInstancePoolPoolInfoAzureAttributes {
	var returns *DataDatabricksInstancePoolPoolInfoAzureAttributes
	_jsii_.Get(
		j,
		"azureAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) DefaultTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) DefaultTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) DiskSpec() DataDatabricksInstancePoolPoolInfoDiskSpecOutputReference {
	var returns DataDatabricksInstancePoolPoolInfoDiskSpecOutputReference
	_jsii_.Get(
		j,
		"diskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) DiskSpecInput() *DataDatabricksInstancePoolPoolInfoDiskSpec {
	var returns *DataDatabricksInstancePoolPoolInfoDiskSpec
	_jsii_.Get(
		j,
		"diskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) EnableElasticDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) EnableElasticDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableElasticDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GcpAttributes() DataDatabricksInstancePoolPoolInfoGcpAttributesOutputReference {
	var returns DataDatabricksInstancePoolPoolInfoGcpAttributesOutputReference
	_jsii_.Get(
		j,
		"gcpAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GcpAttributesInput() *DataDatabricksInstancePoolPoolInfoGcpAttributes {
	var returns *DataDatabricksInstancePoolPoolInfoGcpAttributes
	_jsii_.Get(
		j,
		"gcpAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) IdleInstanceAutoterminationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInstanceAutoterminationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) IdleInstanceAutoterminationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInstanceAutoterminationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolFleetAttributes() DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesList {
	var returns DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesList
	_jsii_.Get(
		j,
		"instancePoolFleetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolFleetAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancePoolFleetAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InstancePoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InternalValue() *DataDatabricksInstancePoolPoolInfo {
	var returns *DataDatabricksInstancePoolPoolInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PreloadedDockerImage() DataDatabricksInstancePoolPoolInfoPreloadedDockerImageList {
	var returns DataDatabricksInstancePoolPoolInfoPreloadedDockerImageList
	_jsii_.Get(
		j,
		"preloadedDockerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PreloadedDockerImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadedDockerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PreloadedSparkVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preloadedSparkVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PreloadedSparkVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preloadedSparkVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) Stats() DataDatabricksInstancePoolPoolInfoStatsOutputReference {
	var returns DataDatabricksInstancePoolPoolInfoStatsOutputReference
	_jsii_.Get(
		j,
		"stats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) StatsInput() *DataDatabricksInstancePoolPoolInfoStats {
	var returns *DataDatabricksInstancePoolPoolInfoStats
	_jsii_.Get(
		j,
		"statsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksInstancePoolPoolInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatabricksInstancePoolPoolInfoOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksInstancePoolPoolInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatabricksInstancePoolPoolInfoOutputReference_Override(d DataDatabricksInstancePoolPoolInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksInstancePool.DataDatabricksInstancePoolPoolInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetDefaultTags(val *map[string]*string) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetEnableElasticDisk(val interface{}) {
	if err := j.validateSetEnableElasticDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableElasticDisk",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetIdleInstanceAutoterminationMinutes(val *float64) {
	if err := j.validateSetIdleInstanceAutoterminationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleInstanceAutoterminationMinutes",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetInstancePoolId(val *string) {
	if err := j.validateSetInstancePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetInstancePoolName(val *string) {
	if err := j.validateSetInstancePoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetInternalValue(val *DataDatabricksInstancePoolPoolInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetPreloadedSparkVersions(val *[]*string) {
	if err := j.validateSetPreloadedSparkVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preloadedSparkVersions",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutAwsAttributes(value *DataDatabricksInstancePoolPoolInfoAwsAttributes) {
	if err := d.validatePutAwsAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutAzureAttributes(value *DataDatabricksInstancePoolPoolInfoAzureAttributes) {
	if err := d.validatePutAzureAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutDiskSpec(value *DataDatabricksInstancePoolPoolInfoDiskSpec) {
	if err := d.validatePutDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDiskSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutGcpAttributes(value *DataDatabricksInstancePoolPoolInfoGcpAttributes) {
	if err := d.validatePutGcpAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcpAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutInstancePoolFleetAttributes(value interface{}) {
	if err := d.validatePutInstancePoolFleetAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstancePoolFleetAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutPreloadedDockerImage(value interface{}) {
	if err := d.validatePutPreloadedDockerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPreloadedDockerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) PutStats(value *DataDatabricksInstancePoolPoolInfoStats) {
	if err := d.validatePutStatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStats",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetAwsAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetAzureAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetCustomTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetDiskSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetEnableElasticDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableElasticDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetGcpAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetGcpAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetInstancePoolFleetAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolFleetAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetInstancePoolId() {
	_jsii_.InvokeVoid(
		d,
		"resetInstancePoolId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetNodeTypeId() {
	_jsii_.InvokeVoid(
		d,
		"resetNodeTypeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetPreloadedDockerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetPreloadedDockerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetPreloadedSparkVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreloadedSparkVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ResetStats() {
	_jsii_.InvokeVoid(
		d,
		"resetStats",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksInstancePoolPoolInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

