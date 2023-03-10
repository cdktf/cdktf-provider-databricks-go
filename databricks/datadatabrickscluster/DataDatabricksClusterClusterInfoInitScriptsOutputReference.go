package datadatabrickscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/datadatabrickscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksClusterClusterInfoInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() DataDatabricksClusterClusterInfoInitScriptsAbfssOutputReference
	AbfssInput() *DataDatabricksClusterClusterInfoInitScriptsAbfss
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
	Dbfs() DataDatabricksClusterClusterInfoInitScriptsDbfsOutputReference
	DbfsInput() *DataDatabricksClusterClusterInfoInitScriptsDbfs
	File() DataDatabricksClusterClusterInfoInitScriptsFileOutputReference
	FileInput() *DataDatabricksClusterClusterInfoInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() DataDatabricksClusterClusterInfoInitScriptsGcsOutputReference
	GcsInput() *DataDatabricksClusterClusterInfoInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() DataDatabricksClusterClusterInfoInitScriptsS3OutputReference
	S3Input() *DataDatabricksClusterClusterInfoInitScriptsS3
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
	PutAbfss(value *DataDatabricksClusterClusterInfoInitScriptsAbfss)
	PutDbfs(value *DataDatabricksClusterClusterInfoInitScriptsDbfs)
	PutFile(value *DataDatabricksClusterClusterInfoInitScriptsFile)
	PutGcs(value *DataDatabricksClusterClusterInfoInitScriptsGcs)
	PutS3(value *DataDatabricksClusterClusterInfoInitScriptsS3)
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

// The jsii proxy struct for DataDatabricksClusterClusterInfoInitScriptsOutputReference
type jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) Abfss() DataDatabricksClusterClusterInfoInitScriptsAbfssOutputReference {
	var returns DataDatabricksClusterClusterInfoInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) AbfssInput() *DataDatabricksClusterClusterInfoInitScriptsAbfss {
	var returns *DataDatabricksClusterClusterInfoInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) Dbfs() DataDatabricksClusterClusterInfoInitScriptsDbfsOutputReference {
	var returns DataDatabricksClusterClusterInfoInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) DbfsInput() *DataDatabricksClusterClusterInfoInitScriptsDbfs {
	var returns *DataDatabricksClusterClusterInfoInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) File() DataDatabricksClusterClusterInfoInitScriptsFileOutputReference {
	var returns DataDatabricksClusterClusterInfoInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) FileInput() *DataDatabricksClusterClusterInfoInitScriptsFile {
	var returns *DataDatabricksClusterClusterInfoInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) Gcs() DataDatabricksClusterClusterInfoInitScriptsGcsOutputReference {
	var returns DataDatabricksClusterClusterInfoInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GcsInput() *DataDatabricksClusterClusterInfoInitScriptsGcs {
	var returns *DataDatabricksClusterClusterInfoInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) S3() DataDatabricksClusterClusterInfoInitScriptsS3OutputReference {
	var returns DataDatabricksClusterClusterInfoInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) S3Input() *DataDatabricksClusterClusterInfoInitScriptsS3 {
	var returns *DataDatabricksClusterClusterInfoInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatabricksClusterClusterInfoInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksClusterClusterInfoInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksClusterClusterInfoInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksClusterClusterInfoInitScriptsOutputReference_Override(d DataDatabricksClusterClusterInfoInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksCluster.DataDatabricksClusterClusterInfoInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) PutAbfss(value *DataDatabricksClusterClusterInfoInitScriptsAbfss) {
	if err := d.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAbfss",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) PutDbfs(value *DataDatabricksClusterClusterInfoInitScriptsDbfs) {
	if err := d.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) PutFile(value *DataDatabricksClusterClusterInfoInitScriptsFile) {
	if err := d.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) PutGcs(value *DataDatabricksClusterClusterInfoInitScriptsGcs) {
	if err := d.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) PutS3(value *DataDatabricksClusterClusterInfoInitScriptsS3) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		d,
		"resetAbfss",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		d,
		"resetDbfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		d,
		"resetFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		d,
		"resetGcs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksClusterClusterInfoInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

