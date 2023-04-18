package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v7/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobLibraryOutputReference interface {
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
	Cran() JobLibraryCranOutputReference
	CranInput() *JobLibraryCran
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
	Maven() JobLibraryMavenOutputReference
	MavenInput() *JobLibraryMaven
	Pypi() JobLibraryPypiOutputReference
	PypiInput() *JobLibraryPypi
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
	PutCran(value *JobLibraryCran)
	PutMaven(value *JobLibraryMaven)
	PutPypi(value *JobLibraryPypi)
	ResetCran()
	ResetEgg()
	ResetJar()
	ResetMaven()
	ResetPypi()
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

// The jsii proxy struct for JobLibraryOutputReference
type jsiiProxy_JobLibraryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobLibraryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Cran() JobLibraryCranOutputReference {
	var returns JobLibraryCranOutputReference
	_jsii_.Get(
		j,
		"cran",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) CranInput() *JobLibraryCran {
	var returns *JobLibraryCran
	_jsii_.Get(
		j,
		"cranInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Egg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) EggInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eggInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Maven() JobLibraryMavenOutputReference {
	var returns JobLibraryMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) MavenInput() *JobLibraryMaven {
	var returns *JobLibraryMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Pypi() JobLibraryPypiOutputReference {
	var returns JobLibraryPypiOutputReference
	_jsii_.Get(
		j,
		"pypi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) PypiInput() *JobLibraryPypi {
	var returns *JobLibraryPypi
	_jsii_.Get(
		j,
		"pypiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewJobLibraryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobLibraryOutputReference {
	_init_.Initialize()

	if err := validateNewJobLibraryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobLibraryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobLibraryOutputReference_Override(j JobLibraryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetEgg(val *string) {
	if err := j.validateSetEggParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egg",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobLibraryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobLibraryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobLibraryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobLibraryOutputReference) PutCran(value *JobLibraryCran) {
	if err := j.validatePutCranParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putCran",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) PutMaven(value *JobLibraryMaven) {
	if err := j.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putMaven",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) PutPypi(value *JobLibraryPypi) {
	if err := j.validatePutPypiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPypi",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetCran() {
	_jsii_.InvokeVoid(
		j,
		"resetCran",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetEgg() {
	_jsii_.InvokeVoid(
		j,
		"resetEgg",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		j,
		"resetJar",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		j,
		"resetMaven",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetPypi() {
	_jsii_.InvokeVoid(
		j,
		"resetPypi",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		j,
		"resetWhl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobLibraryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobLibraryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

