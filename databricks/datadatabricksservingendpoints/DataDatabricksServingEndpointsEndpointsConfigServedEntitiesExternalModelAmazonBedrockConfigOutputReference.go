// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksservingendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/datadatabricksservingendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference interface {
	cdktf.ComplexObject
	AwsAccessKeyId() *string
	SetAwsAccessKeyId(val *string)
	AwsAccessKeyIdInput() *string
	AwsAccessKeyIdPlaintext() *string
	SetAwsAccessKeyIdPlaintext(val *string)
	AwsAccessKeyIdPlaintextInput() *string
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	AwsSecretAccessKey() *string
	SetAwsSecretAccessKey(val *string)
	AwsSecretAccessKeyInput() *string
	AwsSecretAccessKeyPlaintext() *string
	SetAwsSecretAccessKeyPlaintext(val *string)
	AwsSecretAccessKeyPlaintextInput() *string
	BedrockProvider() *string
	SetBedrockProvider(val *string)
	BedrockProviderInput() *string
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
	// Experimental.
	Fqn() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetAwsAccessKeyId()
	ResetAwsAccessKeyIdPlaintext()
	ResetAwsSecretAccessKey()
	ResetAwsSecretAccessKeyPlaintext()
	ResetInstanceProfileArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference
type jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsAccessKeyIdPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyIdPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsAccessKeyIdPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyIdPlaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsSecretAccessKeyPlaintext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKeyPlaintext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) AwsSecretAccessKeyPlaintextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKeyPlaintextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) BedrockProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) BedrockProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference_Override(d DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksServingEndpoints.DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetAwsAccessKeyId(val *string) {
	if err := j.validateSetAwsAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetAwsAccessKeyIdPlaintext(val *string) {
	if err := j.validateSetAwsAccessKeyIdPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccessKeyIdPlaintext",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetAwsSecretAccessKey(val *string) {
	if err := j.validateSetAwsSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetAwsSecretAccessKeyPlaintext(val *string) {
	if err := j.validateSetAwsSecretAccessKeyPlaintextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretAccessKeyPlaintext",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetBedrockProvider(val *string) {
	if err := j.validateSetBedrockProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bedrockProvider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ResetAwsAccessKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAccessKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ResetAwsAccessKeyIdPlaintext() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsAccessKeyIdPlaintext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ResetAwsSecretAccessKey() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsSecretAccessKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ResetAwsSecretAccessKeyPlaintext() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsSecretAccessKeyPlaintext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksServingEndpointsEndpointsConfigServedEntitiesExternalModelAmazonBedrockConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

