// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksappssettingscustomtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksappssettingscustomtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference
	JobSpecInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	SecretSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference
	SecretSpecInput() interface{}
	ServingEndpointSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference
	ServingEndpointSpecInput() interface{}
	SqlWarehouseSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference
	SqlWarehouseSpecInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UcSecurableSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference
	UcSecurableSpecInput() interface{}
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
	PutJobSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpec)
	PutSecretSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpec)
	PutServingEndpointSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec)
	PutSqlWarehouseSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec)
	PutUcSecurableSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec)
	ResetDescription()
	ResetJobSpec()
	ResetSecretSpec()
	ResetServingEndpointSpec()
	ResetSqlWarehouseSpec()
	ResetUcSecurableSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference
type jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) JobSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference {
	var returns DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpecOutputReference
	_jsii_.Get(
		j,
		"jobSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) JobSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SecretSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference {
	var returns DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpecOutputReference
	_jsii_.Get(
		j,
		"secretSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SecretSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ServingEndpointSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference {
	var returns DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpecOutputReference
	_jsii_.Get(
		j,
		"servingEndpointSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ServingEndpointSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"servingEndpointSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SqlWarehouseSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference {
	var returns DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpecOutputReference
	_jsii_.Get(
		j,
		"sqlWarehouseSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) SqlWarehouseSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlWarehouseSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) UcSecurableSpec() DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference {
	var returns DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpecOutputReference
	_jsii_.Get(
		j,
		"ucSecurableSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) UcSecurableSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ucSecurableSpecInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAppsSettingsCustomTemplate.DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference_Override(d DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAppsSettingsCustomTemplate.DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutJobSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsJobSpec) {
	if err := d.validatePutJobSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutSecretSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSecretSpec) {
	if err := d.validatePutSecretSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecretSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutServingEndpointSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsServingEndpointSpec) {
	if err := d.validatePutServingEndpointSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServingEndpointSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutSqlWarehouseSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsSqlWarehouseSpec) {
	if err := d.validatePutSqlWarehouseSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlWarehouseSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) PutUcSecurableSpec(value *DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsUcSecurableSpec) {
	if err := d.validatePutUcSecurableSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUcSecurableSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetJobSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetJobSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetSecretSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetServingEndpointSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetServingEndpointSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetSqlWarehouseSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlWarehouseSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ResetUcSecurableSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetUcSecurableSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksAppsSettingsCustomTemplateManifestResourceSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

