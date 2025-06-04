// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference interface {
	cdktf.ComplexObject
	Abfss() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference
	AbfssInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss
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
	Dbfs() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference
	DbfsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs
	File() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference
	FileInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile
	// Experimental.
	Fqn() *string
	Gcs() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference
	GcsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference
	S3Input() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference
	VolumesInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes
	Workspace() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference
	WorkspaceInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace
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
	PutAbfss(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss)
	PutDbfs(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs)
	PutFile(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile)
	PutGcs(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs)
	PutS3(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3)
	PutVolumes(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes)
	PutWorkspace(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace)
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

// The jsii proxy struct for DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference
type jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Abfss() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfssOutputReference
	_jsii_.Get(
		j,
		"abfss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) AbfssInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss
	_jsii_.Get(
		j,
		"abfssInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Dbfs() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfsOutputReference
	_jsii_.Get(
		j,
		"dbfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) DbfsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs
	_jsii_.Get(
		j,
		"dbfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) File() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFileOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) FileInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Gcs() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GcsInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) S3() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) S3Input() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3 {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Volumes() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumesOutputReference
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) VolumesInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Workspace() DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference {
	var returns DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspaceOutputReference
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) WorkspaceInput() *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace {
	var returns *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference_Override(d DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksJob.DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutAbfss(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsAbfss) {
	if err := d.validatePutAbfssParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAbfss",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutDbfs(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsDbfs) {
	if err := d.validatePutDbfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDbfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutFile(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsFile) {
	if err := d.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutGcs(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsGcs) {
	if err := d.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutS3(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsS3) {
	if err := d.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutVolumes(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsVolumes) {
	if err := d.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) PutWorkspace(value *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsWorkspace) {
	if err := d.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetAbfss() {
	_jsii_.InvokeVoid(
		d,
		"resetAbfss",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetDbfs() {
	_jsii_.InvokeVoid(
		d,
		"resetDbfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		d,
		"resetFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		d,
		"resetGcs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		d,
		"resetS3",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterInitScriptsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

