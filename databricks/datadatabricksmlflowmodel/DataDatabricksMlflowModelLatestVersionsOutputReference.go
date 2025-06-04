// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksmlflowmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/datadatabricksmlflowmodel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksMlflowModelLatestVersionsOutputReference interface {
	cdktf.ComplexObject
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
	CreationTimestamp() *float64
	SetCreationTimestamp(val *float64)
	CreationTimestampInput() *float64
	CurrentStage() *string
	SetCurrentStage(val *string)
	CurrentStageInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastUpdatedTimestamp() *float64
	SetLastUpdatedTimestamp(val *float64)
	LastUpdatedTimestampInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	RunId() *string
	SetRunId(val *string)
	RunIdInput() *string
	RunLink() *string
	SetRunLink(val *string)
	RunLinkInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StatusMessage() *string
	SetStatusMessage(val *string)
	StatusMessageInput() *string
	Tags() DataDatabricksMlflowModelLatestVersionsTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutTags(value interface{})
	ResetCreationTimestamp()
	ResetCurrentStage()
	ResetDescription()
	ResetLastUpdatedTimestamp()
	ResetName()
	ResetRunId()
	ResetRunLink()
	ResetSource()
	ResetStatus()
	ResetStatusMessage()
	ResetTags()
	ResetUserId()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksMlflowModelLatestVersionsOutputReference
type jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) CreationTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) CreationTimestampInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) CurrentStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) CurrentStageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) LastUpdatedTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) LastUpdatedTimestampInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastUpdatedTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) RunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) RunIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) RunLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) RunLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) StatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) StatusMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Tags() DataDatabricksMlflowModelLatestVersionsTagsList {
	var returns DataDatabricksMlflowModelLatestVersionsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksMlflowModelLatestVersionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksMlflowModelLatestVersionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksMlflowModelLatestVersionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMlflowModel.DataDatabricksMlflowModelLatestVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksMlflowModelLatestVersionsOutputReference_Override(d DataDatabricksMlflowModelLatestVersionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksMlflowModel.DataDatabricksMlflowModelLatestVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetCreationTimestamp(val *float64) {
	if err := j.validateSetCreationTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTimestamp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetCurrentStage(val *string) {
	if err := j.validateSetCurrentStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentStage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetLastUpdatedTimestamp(val *float64) {
	if err := j.validateSetLastUpdatedTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastUpdatedTimestamp",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetRunId(val *string) {
	if err := j.validateSetRunIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetRunLink(val *string) {
	if err := j.validateSetRunLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runLink",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetStatusMessage(val *string) {
	if err := j.validateSetStatusMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusMessage",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetCreationTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetCreationTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetCurrentStage() {
	_jsii_.InvokeVoid(
		d,
		"resetCurrentStage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetLastUpdatedTimestamp() {
	_jsii_.InvokeVoid(
		d,
		"resetLastUpdatedTimestamp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetRunId() {
	_jsii_.InvokeVoid(
		d,
		"resetRunId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetRunLink() {
	_jsii_.InvokeVoid(
		d,
		"resetRunLink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetStatusMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetStatusMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		d,
		"resetUserId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksMlflowModelLatestVersionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

