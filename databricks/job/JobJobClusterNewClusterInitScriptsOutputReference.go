package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v5/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobJobClusterNewClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() JobJobClusterNewClusterInitScriptsAbfssOutputReference
	AbfssInput() *JobJobClusterNewClusterInitScriptsAbfss
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
	Dbfs() JobJobClusterNewClusterInitScriptsDbfsOutputReference
	DbfsInput() *JobJobClusterNewClusterInitScriptsDbfs
	File() JobJobClusterNewClusterInitScriptsFileOutputReference
	FileInput() *JobJobClusterNewClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() JobJobClusterNewClusterInitScriptsGcsOutputReference
	GcsInput() *JobJobClusterNewClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() JobJobClusterNewClusterInitScriptsS3OutputReference
	S3Input() *JobJobClusterNewClusterInitScriptsS3
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
	PutAbfss(value *JobJobClusterNewClusterInitScriptsAbfss)
	PutDbfs(value *JobJobClusterNewClusterInitScriptsDbfs)
	PutFile(value *JobJobClusterNewClusterInitScriptsFile)
	PutGcs(value *JobJobClusterNewClusterInitScriptsGcs)
	PutS3(value *JobJobClusterNewClusterInitScriptsS3)
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

// The jsii proxy struct for JobJobClusterNewClusterInitScriptsOutputReference
type jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) Abfss() JobJobClusterNewClusterInitScriptsAbfssOutputReference {
	var returns JobJobClusterNewClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) AbfssInput() *JobJobClusterNewClusterInitScriptsAbfss {
	var returns *JobJobClusterNewClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) Dbfs() JobJobClusterNewClusterInitScriptsDbfsOutputReference {
	var returns JobJobClusterNewClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) DbfsInput() *JobJobClusterNewClusterInitScriptsDbfs {
	var returns *JobJobClusterNewClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) File() JobJobClusterNewClusterInitScriptsFileOutputReference {
	var returns JobJobClusterNewClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) FileInput() *JobJobClusterNewClusterInitScriptsFile {
	var returns *JobJobClusterNewClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) Gcs() JobJobClusterNewClusterInitScriptsGcsOutputReference {
	var returns JobJobClusterNewClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GcsInput() *JobJobClusterNewClusterInitScriptsGcs {
	var returns *JobJobClusterNewClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) S3() JobJobClusterNewClusterInitScriptsS3OutputReference {
	var returns JobJobClusterNewClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) S3Input() *JobJobClusterNewClusterInitScriptsS3 {
	var returns *JobJobClusterNewClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobJobClusterNewClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobJobClusterNewClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewJobJobClusterNewClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobJobClusterNewClusterInitScriptsOutputReference_Override(j JobJobClusterNewClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.job.JobJobClusterNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) PutAbfss(value *JobJobClusterNewClusterInitScriptsAbfss) {
	if err := j.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAbfss",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) PutDbfs(value *JobJobClusterNewClusterInitScriptsDbfs) {
	if err := j.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDbfs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) PutFile(value *JobJobClusterNewClusterInitScriptsFile) {
	if err := j.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putFile",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) PutGcs(value *JobJobClusterNewClusterInitScriptsGcs) {
	if err := j.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putGcs",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) PutS3(value *JobJobClusterNewClusterInitScriptsS3) {
	if err := j.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putS3",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		j,
		"resetAbfss",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		j,
		"resetDbfs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		j,
		"resetFile",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		j,
		"resetGcs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		j,
		"resetS3",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (j *jsiiProxy_JobJobClusterNewClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

