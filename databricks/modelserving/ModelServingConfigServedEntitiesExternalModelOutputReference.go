// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelserving

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/modelserving/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelServingConfigServedEntitiesExternalModelOutputReference interface {
	cdktf.ComplexObject
	Ai21LabsConfig() ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference
	Ai21LabsConfigInput() *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig
	AmazonBedrockConfig() ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference
	AmazonBedrockConfigInput() *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig
	AnthropicConfig() ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference
	AnthropicConfigInput() *ModelServingConfigServedEntitiesExternalModelAnthropicConfig
	CohereConfig() ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference
	CohereConfigInput() *ModelServingConfigServedEntitiesExternalModelCohereConfig
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
	CustomProviderConfig() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference
	CustomProviderConfigInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig
	DatabricksModelServingConfig() ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference
	DatabricksModelServingConfigInput() *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig
	// Experimental.
	Fqn() *string
	GoogleCloudVertexAiConfig() ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference
	GoogleCloudVertexAiConfigInput() *ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig
	InternalValue() *ModelServingConfigServedEntitiesExternalModel
	SetInternalValue(val *ModelServingConfigServedEntitiesExternalModel)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OpenaiConfig() ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference
	OpenaiConfigInput() *ModelServingConfigServedEntitiesExternalModelOpenaiConfig
	PalmConfig() ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference
	PalmConfigInput() *ModelServingConfigServedEntitiesExternalModelPalmConfig
	Provider() *string
	SetProvider(val *string)
	ProviderInput() *string
	Task() *string
	SetTask(val *string)
	TaskInput() *string
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
	PutAi21LabsConfig(value *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig)
	PutAmazonBedrockConfig(value *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig)
	PutAnthropicConfig(value *ModelServingConfigServedEntitiesExternalModelAnthropicConfig)
	PutCohereConfig(value *ModelServingConfigServedEntitiesExternalModelCohereConfig)
	PutCustomProviderConfig(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig)
	PutDatabricksModelServingConfig(value *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig)
	PutGoogleCloudVertexAiConfig(value *ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig)
	PutOpenaiConfig(value *ModelServingConfigServedEntitiesExternalModelOpenaiConfig)
	PutPalmConfig(value *ModelServingConfigServedEntitiesExternalModelPalmConfig)
	ResetAi21LabsConfig()
	ResetAmazonBedrockConfig()
	ResetAnthropicConfig()
	ResetCohereConfig()
	ResetCustomProviderConfig()
	ResetDatabricksModelServingConfig()
	ResetGoogleCloudVertexAiConfig()
	ResetOpenaiConfig()
	ResetPalmConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelServingConfigServedEntitiesExternalModelOutputReference
type jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Ai21LabsConfig() ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelAi21LabsConfigOutputReference
	_jsii_.Get(
		j,
		"ai21LabsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Ai21LabsConfigInput() *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig
	_jsii_.Get(
		j,
		"ai21LabsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) AmazonBedrockConfig() ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference
	_jsii_.Get(
		j,
		"amazonBedrockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) AmazonBedrockConfigInput() *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig
	_jsii_.Get(
		j,
		"amazonBedrockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) AnthropicConfig() ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelAnthropicConfigOutputReference
	_jsii_.Get(
		j,
		"anthropicConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) AnthropicConfigInput() *ModelServingConfigServedEntitiesExternalModelAnthropicConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelAnthropicConfig
	_jsii_.Get(
		j,
		"anthropicConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) CohereConfig() ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelCohereConfigOutputReference
	_jsii_.Get(
		j,
		"cohereConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) CohereConfigInput() *ModelServingConfigServedEntitiesExternalModelCohereConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelCohereConfig
	_jsii_.Get(
		j,
		"cohereConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) CustomProviderConfig() ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelCustomProviderConfigOutputReference
	_jsii_.Get(
		j,
		"customProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) CustomProviderConfigInput() *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig
	_jsii_.Get(
		j,
		"customProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) DatabricksModelServingConfig() ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfigOutputReference
	_jsii_.Get(
		j,
		"databricksModelServingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) DatabricksModelServingConfigInput() *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig
	_jsii_.Get(
		j,
		"databricksModelServingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GoogleCloudVertexAiConfig() ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfigOutputReference
	_jsii_.Get(
		j,
		"googleCloudVertexAiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GoogleCloudVertexAiConfigInput() *ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig
	_jsii_.Get(
		j,
		"googleCloudVertexAiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) InternalValue() *ModelServingConfigServedEntitiesExternalModel {
	var returns *ModelServingConfigServedEntitiesExternalModel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) OpenaiConfig() ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelOpenaiConfigOutputReference
	_jsii_.Get(
		j,
		"openaiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) OpenaiConfigInput() *ModelServingConfigServedEntitiesExternalModelOpenaiConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelOpenaiConfig
	_jsii_.Get(
		j,
		"openaiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PalmConfig() ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference {
	var returns ModelServingConfigServedEntitiesExternalModelPalmConfigOutputReference
	_jsii_.Get(
		j,
		"palmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PalmConfigInput() *ModelServingConfigServedEntitiesExternalModelPalmConfig {
	var returns *ModelServingConfigServedEntitiesExternalModelPalmConfig
	_jsii_.Get(
		j,
		"palmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) TaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelServingConfigServedEntitiesExternalModelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelServingConfigServedEntitiesExternalModelOutputReference {
	_init_.Initialize()

	if err := validateNewModelServingConfigServedEntitiesExternalModelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelServingConfigServedEntitiesExternalModelOutputReference_Override(m ModelServingConfigServedEntitiesExternalModelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.modelServing.ModelServingConfigServedEntitiesExternalModelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetInternalValue(val *ModelServingConfigServedEntitiesExternalModel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetTask(val *string) {
	if err := j.validateSetTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"task",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutAi21LabsConfig(value *ModelServingConfigServedEntitiesExternalModelAi21LabsConfig) {
	if err := m.validatePutAi21LabsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAi21LabsConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutAmazonBedrockConfig(value *ModelServingConfigServedEntitiesExternalModelAmazonBedrockConfig) {
	if err := m.validatePutAmazonBedrockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAmazonBedrockConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutAnthropicConfig(value *ModelServingConfigServedEntitiesExternalModelAnthropicConfig) {
	if err := m.validatePutAnthropicConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAnthropicConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutCohereConfig(value *ModelServingConfigServedEntitiesExternalModelCohereConfig) {
	if err := m.validatePutCohereConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCohereConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutCustomProviderConfig(value *ModelServingConfigServedEntitiesExternalModelCustomProviderConfig) {
	if err := m.validatePutCustomProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCustomProviderConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutDatabricksModelServingConfig(value *ModelServingConfigServedEntitiesExternalModelDatabricksModelServingConfig) {
	if err := m.validatePutDatabricksModelServingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDatabricksModelServingConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutGoogleCloudVertexAiConfig(value *ModelServingConfigServedEntitiesExternalModelGoogleCloudVertexAiConfig) {
	if err := m.validatePutGoogleCloudVertexAiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGoogleCloudVertexAiConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutOpenaiConfig(value *ModelServingConfigServedEntitiesExternalModelOpenaiConfig) {
	if err := m.validatePutOpenaiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOpenaiConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) PutPalmConfig(value *ModelServingConfigServedEntitiesExternalModelPalmConfig) {
	if err := m.validatePutPalmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPalmConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetAi21LabsConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAi21LabsConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetAmazonBedrockConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAmazonBedrockConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetAnthropicConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAnthropicConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetCohereConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetCohereConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetCustomProviderConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomProviderConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetDatabricksModelServingConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetDatabricksModelServingConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetGoogleCloudVertexAiConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetGoogleCloudVertexAiConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetOpenaiConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenaiConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ResetPalmConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetPalmConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelServingConfigServedEntitiesExternalModelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

