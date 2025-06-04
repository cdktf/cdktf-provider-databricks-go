// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v14/automaticclusterupdateworkspacesetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference interface {
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
	EnablementDetails() AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsList
	// Experimental.
	Fqn() *string
	InternalValue() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace
	SetInternalValue(val *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace)
	MaintenanceWindow() AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference
	MaintenanceWindowInput() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow
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
	PutMaintenanceWindow(value *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow)
	ResetCanToggle()
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

// The jsii proxy struct for AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference
type jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) CanToggle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canToggle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) CanToggleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canToggleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) EnablementDetails() AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsList {
	var returns AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsList
	_jsii_.Get(
		j,
		"enablementDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) InternalValue() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace {
	var returns *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) MaintenanceWindow() AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference {
	var returns AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) MaintenanceWindowInput() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow {
	var returns *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) RestartEvenIfNoUpdatesAvailable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartEvenIfNoUpdatesAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) RestartEvenIfNoUpdatesAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartEvenIfNoUpdatesAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference {
	_init_.Initialize()

	if err := validateNewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.automaticClusterUpdateWorkspaceSetting.AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference_Override(a AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.automaticClusterUpdateWorkspaceSetting.AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetCanToggle(val interface{}) {
	if err := j.validateSetCanToggleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canToggle",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetInternalValue(val *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspace) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetRestartEvenIfNoUpdatesAvailable(val interface{}) {
	if err := j.validateSetRestartEvenIfNoUpdatesAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartEvenIfNoUpdatesAvailable",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) PutMaintenanceWindow(value *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceMaintenanceWindow) {
	if err := a.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ResetCanToggle() {
	_jsii_.InvokeVoid(
		a,
		"resetCanToggle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ResetRestartEvenIfNoUpdatesAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartEvenIfNoUpdatesAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

