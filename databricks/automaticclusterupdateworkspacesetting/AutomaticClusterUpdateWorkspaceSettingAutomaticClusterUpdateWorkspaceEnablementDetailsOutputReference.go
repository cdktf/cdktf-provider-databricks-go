// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automaticclusterupdateworkspacesetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/automaticclusterupdateworkspacesetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference interface {
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
	InternalValue() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails
	SetInternalValue(val *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails)
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

// The jsii proxy struct for AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference
type jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ForcedForComplianceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcedForComplianceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ForcedForComplianceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcedForComplianceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InternalValue() *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails {
	var returns *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForDisabledEntitlement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForDisabledEntitlement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForDisabledEntitlementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForDisabledEntitlementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForNonEnterpriseTier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForNonEnterpriseTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) UnavailableForNonEnterpriseTierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unavailableForNonEnterpriseTierInput",
		&returns,
	)
	return returns
}


func NewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-databricks.automaticClusterUpdateWorkspaceSetting.AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference_Override(a AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.automaticClusterUpdateWorkspaceSetting.AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetForcedForComplianceMode(val interface{}) {
	if err := j.validateSetForcedForComplianceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forcedForComplianceMode",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetInternalValue(val *AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetUnavailableForDisabledEntitlement(val interface{}) {
	if err := j.validateSetUnavailableForDisabledEntitlementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unavailableForDisabledEntitlement",
		val,
	)
}

func (j *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference)SetUnavailableForNonEnterpriseTier(val interface{}) {
	if err := j.validateSetUnavailableForNonEnterpriseTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unavailableForNonEnterpriseTier",
		val,
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetForcedForComplianceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetForcedForComplianceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetUnavailableForDisabledEntitlement() {
	_jsii_.InvokeVoid(
		a,
		"resetUnavailableForDisabledEntitlement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ResetUnavailableForNonEnterpriseTier() {
	_jsii_.InvokeVoid(
		a,
		"resetUnavailableForNonEnterpriseTier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutomaticClusterUpdateWorkspaceSettingAutomaticClusterUpdateWorkspaceEnablementDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

