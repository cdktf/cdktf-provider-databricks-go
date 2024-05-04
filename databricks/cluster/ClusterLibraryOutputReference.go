// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterLibraryOutputReference interface {
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
	Cran() ClusterLibraryCranOutputReference
	CranInput() *ClusterLibraryCran
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
	Maven() ClusterLibraryMavenOutputReference
	MavenInput() *ClusterLibraryMaven
	Pypi() ClusterLibraryPypiOutputReference
	PypiInput() *ClusterLibraryPypi
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
	PutCran(value *ClusterLibraryCran)
	PutMaven(value *ClusterLibraryMaven)
	PutPypi(value *ClusterLibraryPypi)
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

// The jsii proxy struct for ClusterLibraryOutputReference
type jsiiProxy_ClusterLibraryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterLibraryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Cran() ClusterLibraryCranOutputReference {
	var returns ClusterLibraryCranOutputReference
	_jsii_.Get(
		j,
		"cran",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) CranInput() *ClusterLibraryCran {
	var returns *ClusterLibraryCran
	_jsii_.Get(
		j,
		"cranInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Egg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) EggInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eggInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Maven() ClusterLibraryMavenOutputReference {
	var returns ClusterLibraryMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) MavenInput() *ClusterLibraryMaven {
	var returns *ClusterLibraryMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Pypi() ClusterLibraryPypiOutputReference {
	var returns ClusterLibraryPypiOutputReference
	_jsii_.Get(
		j,
		"pypi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) PypiInput() *ClusterLibraryPypi {
	var returns *ClusterLibraryPypi
	_jsii_.Get(
		j,
		"pypiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Requirements() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) RequirementsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterLibraryOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewClusterLibraryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClusterLibraryOutputReference {
	_init_.Initialize()

	if err := validateNewClusterLibraryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterLibraryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClusterLibraryOutputReference_Override(c ClusterLibraryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterLibraryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetEgg(val *string) {
	if err := j.validateSetEggParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egg",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetRequirements(val *string) {
	if err := j.validateSetRequirementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirements",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClusterLibraryOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) PutCran(value *ClusterLibraryCran) {
	if err := c.validatePutCranParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCran",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) PutMaven(value *ClusterLibraryMaven) {
	if err := c.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaven",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) PutPypi(value *ClusterLibraryPypi) {
	if err := c.validatePutPypiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPypi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetCran() {
	_jsii_.InvokeVoid(
		c,
		"resetCran",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetEgg() {
	_jsii_.InvokeVoid(
		c,
		"resetEgg",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		c,
		"resetJar",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		c,
		"resetMaven",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetPypi() {
	_jsii_.InvokeVoid(
		c,
		"resetPypi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetRequirements() {
	_jsii_.InvokeVoid(
		c,
		"resetRequirements",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		c,
		"resetWhl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterLibraryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterLibraryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

