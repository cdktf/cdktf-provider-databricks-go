package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() ClusterInitScriptsAbfssOutputReference
	AbfssInput() *ClusterInitScriptsAbfss
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
	Dbfs() ClusterInitScriptsDbfsOutputReference
	DbfsInput() *ClusterInitScriptsDbfs
	File() ClusterInitScriptsFileOutputReference
	FileInput() *ClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() ClusterInitScriptsGcsOutputReference
	GcsInput() *ClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() ClusterInitScriptsS3OutputReference
	S3Input() *ClusterInitScriptsS3
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
	PutAbfss(value *ClusterInitScriptsAbfss)
	PutDbfs(value *ClusterInitScriptsDbfs)
	PutFile(value *ClusterInitScriptsFile)
	PutGcs(value *ClusterInitScriptsGcs)
	PutS3(value *ClusterInitScriptsS3)
	ResetAbfss()
	ResetDbfs()
	ResetFile()
	ResetGcs()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterInitScriptsOutputReference
type jsiiProxy_ClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) Abfss() ClusterInitScriptsAbfssOutputReference {
	var returns ClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) AbfssInput() *ClusterInitScriptsAbfss {
	var returns *ClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) Dbfs() ClusterInitScriptsDbfsOutputReference {
	var returns ClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) DbfsInput() *ClusterInitScriptsDbfs {
	var returns *ClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) File() ClusterInitScriptsFileOutputReference {
	var returns ClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) FileInput() *ClusterInitScriptsFile {
	var returns *ClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) Gcs() ClusterInitScriptsGcsOutputReference {
	var returns ClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) GcsInput() *ClusterInitScriptsGcs {
	var returns *ClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) S3() ClusterInitScriptsS3OutputReference {
	var returns ClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) S3Input() *ClusterInitScriptsS3 {
	var returns *ClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClusterInitScriptsOutputReference_Override(c ClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.cluster.ClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) PutAbfss(value *ClusterInitScriptsAbfss) {
	if err := c.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAbfss",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) PutDbfs(value *ClusterInitScriptsDbfs) {
	if err := c.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDbfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) PutFile(value *ClusterInitScriptsFile) {
	if err := c.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) PutGcs(value *ClusterInitScriptsGcs) {
	if err := c.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) PutS3(value *ClusterInitScriptsS3) {
	if err := c.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		c,
		"resetAbfss",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		c,
		"resetDbfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		c,
		"resetFile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		c,
		"resetGcs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		c,
		"resetS3",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

