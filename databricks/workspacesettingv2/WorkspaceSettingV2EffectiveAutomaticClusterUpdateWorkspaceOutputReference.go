// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/workspacesettingv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference interface {
	cdktf.ComplexObject
	CanToggle() interface{}
	SetCanToggle(val interface{})
	CanToggleInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnablementDetails() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference
	EnablementDetailsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaintenanceWindow() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference
	MaintenanceWindowInput() interface{}
	RestartEvenIfNoUpdatesAvailable() interface{}
	SetRestartEvenIfNoUpdatesAvailable(val interface{})
	RestartEvenIfNoUpdatesAvailableInput() interface{}
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
	PutEnablementDetails(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetails)
	PutMaintenanceWindow(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindow)
	ResetCanToggle()
	ResetEnabled()
	ResetEnablementDetails()
	ResetMaintenanceWindow()
	ResetRestartEvenIfNoUpdatesAvailable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
type jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) CanToggle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) CanToggleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canToggleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) EnablementDetails() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference {
	var returns WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference
	_jsii_.Get(
		j,
		"enablementDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) EnablementDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablementDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) MaintenanceWindow() WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference {
	var returns WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) RestartEvenIfNoUpdatesAvailable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartEvenIfNoUpdatesAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) RestartEvenIfNoUpdatesAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartEvenIfNoUpdatesAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference_Override(w WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.workspaceSettingV2.WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetCanToggle(val interface{}) {
	if err := j.validateSetCanToggleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canToggle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetRestartEvenIfNoUpdatesAvailable(val interface{}) {
	if err := j.validateSetRestartEvenIfNoUpdatesAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartEvenIfNoUpdatesAvailable",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) PutEnablementDetails(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceEnablementDetails) {
	if err := w.validatePutEnablementDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEnablementDetails",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) PutMaintenanceWindow(value *WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceMaintenanceWindow) {
	if err := w.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ResetCanToggle() {
	_jsii_.InvokeVoid(
		w,
		"resetCanToggle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ResetEnablementDetails() {
	_jsii_.InvokeVoid(
		w,
		"resetEnablementDetails",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		w,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ResetRestartEvenIfNoUpdatesAvailable() {
	_jsii_.InvokeVoid(
		w,
		"resetRestartEvenIfNoUpdatesAvailable",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

