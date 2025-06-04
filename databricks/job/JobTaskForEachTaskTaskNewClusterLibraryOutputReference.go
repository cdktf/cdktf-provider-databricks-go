// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobTaskForEachTaskTaskNewClusterLibraryOutputReference interface {
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
	Cran() JobTaskForEachTaskTaskNewClusterLibraryCranOutputReference
	CranInput() *JobTaskForEachTaskTaskNewClusterLibraryCran
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
	Maven() JobTaskForEachTaskTaskNewClusterLibraryMavenOutputReference
	MavenInput() *JobTaskForEachTaskTaskNewClusterLibraryMaven
	Pypi() JobTaskForEachTaskTaskNewClusterLibraryPypiOutputReference
	PypiInput() *JobTaskForEachTaskTaskNewClusterLibraryPypi
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
	PutCran(value *JobTaskForEachTaskTaskNewClusterLibraryCran)
	PutMaven(value *JobTaskForEachTaskTaskNewClusterLibraryMaven)
	PutPypi(value *JobTaskForEachTaskTaskNewClusterLibraryPypi)
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

// The jsii proxy struct for JobTaskForEachTaskTaskNewClusterLibraryOutputReference
type jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Cran() JobTaskForEachTaskTaskNewClusterLibraryCranOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterLibraryCranOutputReference
	_jsii_.Get(
		j,
		"cran",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) CranInput() *JobTaskForEachTaskTaskNewClusterLibraryCran {
	var returns *JobTaskForEachTaskTaskNewClusterLibraryCran
	_jsii_.Get(
		j,
		"cranInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Egg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) EggInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eggInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Maven() JobTaskForEachTaskTaskNewClusterLibraryMavenOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterLibraryMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) MavenInput() *JobTaskForEachTaskTaskNewClusterLibraryMaven {
	var returns *JobTaskForEachTaskTaskNewClusterLibraryMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Pypi() JobTaskForEachTaskTaskNewClusterLibraryPypiOutputReference {
	var returns JobTaskForEachTaskTaskNewClusterLibraryPypiOutputReference
	_jsii_.Get(
		j,
		"pypi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) PypiInput() *JobTaskForEachTaskTaskNewClusterLibraryPypi {
	var returns *JobTaskForEachTaskTaskNewClusterLibraryPypi
	_jsii_.Get(
		j,
		"pypiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Requirements() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) RequirementsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewJobTaskForEachTaskTaskNewClusterLibraryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobTaskForEachTaskTaskNewClusterLibraryOutputReference {
	_init_.Initialize()

	if err := validateNewJobTaskForEachTaskTaskNewClusterLibraryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobTaskForEachTaskTaskNewClusterLibraryOutputReference_Override(j JobTaskForEachTaskTaskNewClusterLibraryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobTaskForEachTaskTaskNewClusterLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetEgg(val *string) {
	if err := j.validateSetEggParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egg",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetRequirements(val *string) {
	if err := j.validateSetRequirementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirements",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) PutCran(value *JobTaskForEachTaskTaskNewClusterLibraryCran) {
	if err := j.validatePutCranParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCran",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) PutMaven(value *JobTaskForEachTaskTaskNewClusterLibraryMaven) {
	if err := j.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putMaven",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) PutPypi(value *JobTaskForEachTaskTaskNewClusterLibraryPypi) {
	if err := j.validatePutPypiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPypi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetCran() {
	_jsii_.InvokeVoid(
		j,
		"resetCran",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetEgg() {
	_jsii_.InvokeVoid(
		j,
		"resetEgg",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		j,
		"resetJar",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		j,
		"resetMaven",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetPypi() {
	_jsii_.InvokeVoid(
		j,
		"resetPypi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetRequirements() {
	_jsii_.InvokeVoid(
		j,
		"resetRequirements",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		j,
		"resetWhl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobTaskForEachTaskTaskNewClusterLibraryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

