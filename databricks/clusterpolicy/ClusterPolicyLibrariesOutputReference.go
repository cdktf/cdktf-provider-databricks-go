// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusterpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/clusterpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterPolicyLibrariesOutputReference interface {
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
	Cran() ClusterPolicyLibrariesCranOutputReference
	CranInput() *ClusterPolicyLibrariesCran
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
	Maven() ClusterPolicyLibrariesMavenOutputReference
	MavenInput() *ClusterPolicyLibrariesMaven
	Pypi() ClusterPolicyLibrariesPypiOutputReference
	PypiInput() *ClusterPolicyLibrariesPypi
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
	PutCran(value *ClusterPolicyLibrariesCran)
	PutMaven(value *ClusterPolicyLibrariesMaven)
	PutPypi(value *ClusterPolicyLibrariesPypi)
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

// The jsii proxy struct for ClusterPolicyLibrariesOutputReference
type jsiiProxy_ClusterPolicyLibrariesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Cran() ClusterPolicyLibrariesCranOutputReference {
	var returns ClusterPolicyLibrariesCranOutputReference
	_jsii_.Get(
		j,
		"cran",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) CranInput() *ClusterPolicyLibrariesCran {
	var returns *ClusterPolicyLibrariesCran
	_jsii_.Get(
		j,
		"cranInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Egg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) EggInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eggInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Jar() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) JarInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Maven() ClusterPolicyLibrariesMavenOutputReference {
	var returns ClusterPolicyLibrariesMavenOutputReference
	_jsii_.Get(
		j,
		"maven",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) MavenInput() *ClusterPolicyLibrariesMaven {
	var returns *ClusterPolicyLibrariesMaven
	_jsii_.Get(
		j,
		"mavenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Pypi() ClusterPolicyLibrariesPypiOutputReference {
	var returns ClusterPolicyLibrariesPypiOutputReference
	_jsii_.Get(
		j,
		"pypi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) PypiInput() *ClusterPolicyLibrariesPypi {
	var returns *ClusterPolicyLibrariesPypi
	_jsii_.Get(
		j,
		"pypiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Requirements() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) RequirementsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) Whl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference) WhlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whlInput",
		&returns,
	)
	return returns
}


func NewClusterPolicyLibrariesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClusterPolicyLibrariesOutputReference {
	_init_.Initialize()

	if err := validateNewClusterPolicyLibrariesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterPolicyLibrariesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.clusterPolicy.ClusterPolicyLibrariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClusterPolicyLibrariesOutputReference_Override(c ClusterPolicyLibrariesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.clusterPolicy.ClusterPolicyLibrariesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetEgg(val *string) {
	if err := j.validateSetEggParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egg",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetJar(val *string) {
	if err := j.validateSetJarParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jar",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetRequirements(val *string) {
	if err := j.validateSetRequirementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirements",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClusterPolicyLibrariesOutputReference)SetWhl(val *string) {
	if err := j.validateSetWhlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whl",
		val,
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) PutCran(value *ClusterPolicyLibrariesCran) {
	if err := c.validatePutCranParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCran",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) PutMaven(value *ClusterPolicyLibrariesMaven) {
	if err := c.validatePutMavenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaven",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) PutPypi(value *ClusterPolicyLibrariesPypi) {
	if err := c.validatePutPypiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPypi",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetCran() {
	_jsii_.InvokeVoid(
		c,
		"resetCran",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetEgg() {
	_jsii_.InvokeVoid(
		c,
		"resetEgg",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetJar() {
	_jsii_.InvokeVoid(
		c,
		"resetJar",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetMaven() {
	_jsii_.InvokeVoid(
		c,
		"resetMaven",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetPypi() {
	_jsii_.InvokeVoid(
		c,
		"resetPypi",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetRequirements() {
	_jsii_.InvokeVoid(
		c,
		"resetRequirements",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ResetWhl() {
	_jsii_.InvokeVoid(
		c,
		"resetWhl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterPolicyLibrariesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

