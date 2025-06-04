// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference interface {
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
	Cran() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCranOutputReference
	CranInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCran
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Egg() *string
	SetEgg(val *string)
	EggInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jar() *string
	SetJar(val *string)
	JarInput() *string
	Maven() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMavenOutputReference
	MavenInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMaven
	Pypi() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypiOutputReference
	PypiInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypi
	Requirements() *string
	SetRequirements(val *string)
	RequirementsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Whl() *string
	SetWhl(val *string)
	WhlInput() *string
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
	PutCran(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCran)
	PutMaven(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMaven)
	PutPypi(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypi)
	ResetCran()
	ResetEgg()
	ResetJar()
	ResetMaven()
	ResetPypi()
	ResetRequirements()
	ResetWhl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Cran() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCranOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCranOutputReference
	_jsii_.Get(
		j,
		"cran",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) CranInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCran {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCran
	_jsii_.Get(
		j,
		"cranInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Egg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) EggInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eggInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Maven() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMavenOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) MavenInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMaven {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Pypi() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypiOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypiOutputReference
	_jsii_.Get(
		j,
		"pypi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) PypiInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypi {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypi
	_jsii_.Get(
		j,
		"pypiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Requirements() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) RequirementsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetEgg(val *string) {
	if err := j.validateSetEggParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egg",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetRequirements(val *string) {
	if err := j.validateSetRequirementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirements",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) PutCran(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryCran) {
	if err := d.validatePutCranParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCran",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) PutMaven(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryMaven) {
	if err := d.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMaven",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) PutPypi(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryPypi) {
	if err := d.validatePutPypiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPypi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetCran() {
	_jsii_.InvokeVoid(
		d,
		"resetCran",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetEgg() {
	_jsii_.InvokeVoid(
		d,
		"resetEgg",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		d,
		"resetJar",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		d,
		"resetMaven",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetPypi() {
	_jsii_.InvokeVoid(
		d,
		"resetPypi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetRequirements() {
	_jsii_.InvokeVoid(
		d,
		"resetRequirements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		d,
		"resetWhl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskLibraryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

