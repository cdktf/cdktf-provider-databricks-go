// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/workspacesettingv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForcedForComplianceMode() interface{}
	SetForcedForComplianceMode(val interface{})
	ForcedForComplianceModeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnavailableForDisabledEntitlement() interface{}
	SetUnavailableForDisabledEntitlement(val interface{})
	UnavailableForDisabledEntitlementInput() interface{}
	UnavailableForNonEnterpriseTier() interface{}
	SetUnavailableForNonEnterpriseTier(val interface{})
	UnavailableForNonEnterpriseTierInput() interface{}
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
	ResetForcedForComplianceMode()
	ResetUnavailableForDisabledEntitlement()
	ResetUnavailableForNonEnterpriseTier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference
type jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ForcedForComplianceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcedForComplianceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ForcedForComplianceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcedForComplianceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForDisabledEntitlement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForDisabledEntitlement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForDisabledEntitlementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForDisabledEntitlementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForNonEnterpriseTier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForNonEnterpriseTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForNonEnterpriseTierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForNonEnterpriseTierInput",
		&returns,
	)
	return returns
}


func NewWorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.workspaceSettingV2.WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference_Override(w WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.workspaceSettingV2.WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetForcedForComplianceMode(val interface{}) {
	if err := j.validateSetForcedForComplianceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forcedForComplianceMode",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetUnavailableForDisabledEntitlement(val interface{}) {
	if err := j.validateSetUnavailableForDisabledEntitlementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unavailableForDisabledEntitlement",
		val,
	)
}

func (j *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetUnavailableForNonEnterpriseTier(val interface{}) {
	if err := j.validateSetUnavailableForNonEnterpriseTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unavailableForNonEnterpriseTier",
		val,
	)
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetForcedForComplianceMode() {
	_jsii_.InvokeVoid(
		w,
		"resetForcedForComplianceMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetUnavailableForDisabledEntitlement() {
	_jsii_.InvokeVoid(
		w,
		"resetUnavailableForDisabledEntitlement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetUnavailableForNonEnterpriseTier() {
	_jsii_.InvokeVoid(
		w,
		"resetUnavailableForNonEnterpriseTier",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceSettingV2AutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

