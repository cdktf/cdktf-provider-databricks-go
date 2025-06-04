// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfssOutputReference
	AbfssInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss
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
	Dbfs() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfsOutputReference
	DbfsInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs
	File() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFileOutputReference
	FileInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcsOutputReference
	GcsInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3OutputReference
	S3Input() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumesOutputReference
	VolumesInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes
	Workspace() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspaceOutputReference
	WorkspaceInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace
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
	PutAbfss(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss)
	PutDbfs(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs)
	PutFile(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile)
	PutGcs(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs)
	PutS3(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3)
	PutVolumes(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes)
	PutWorkspace(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace)
	ResetAbfss()
	ResetDbfs()
	ResetFile()
	ResetGcs()
	ResetS3()
	ResetVolumes()
	ResetWorkspace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Abfss() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfssOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) AbfssInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Dbfs() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) DbfsInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) File() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFileOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) FileInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Gcs() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GcsInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) S3() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3OutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) S3Input() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3 {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Volumes() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumesOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumesOutputReference
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) VolumesInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Workspace() DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspaceOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspaceOutputReference
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) WorkspaceInput() *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace {
	var returns *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference_Override(d DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutAbfss(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsAbfss) {
	if err := d.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAbfss",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutDbfs(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsDbfs) {
	if err := d.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutFile(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsFile) {
	if err := d.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutGcs(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsGcs) {
	if err := d.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutS3(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutVolumes(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsVolumes) {
	if err := d.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) PutWorkspace(value *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsWorkspace) {
	if err := d.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		d,
		"resetAbfss",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		d,
		"resetDbfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		d,
		"resetFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		d,
		"resetGcs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

