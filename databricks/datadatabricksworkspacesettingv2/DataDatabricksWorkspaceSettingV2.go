// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksworkspacesettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksworkspacesettingv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2 databricks_workspace_setting_v2}.
type DataDatabricksWorkspaceSettingV2 interface {
	cdktf.TerraformDataSource
	AibiDashboardEmbeddingAccessPolicy() DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	AibiDashboardEmbeddingAccessPolicyInput() interface{}
	AibiDashboardEmbeddingApprovedDomains() DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	AibiDashboardEmbeddingApprovedDomainsInput() interface{}
	AutomaticClusterUpdateWorkspace() DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	AutomaticClusterUpdateWorkspaceInput() interface{}
	BooleanVal() DataDatabricksWorkspaceSettingV2BooleanValOutputReference
	BooleanValInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultDataSecurityMode() DataDatabricksWorkspaceSettingV2DefaultDataSecurityModeOutputReference
	DefaultDataSecurityModeInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveAibiDashboardEmbeddingAccessPolicy() DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{}
	EffectiveAibiDashboardEmbeddingApprovedDomains() DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{}
	EffectiveAutomaticClusterUpdateWorkspace() DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	EffectiveAutomaticClusterUpdateWorkspaceInput() interface{}
	EffectiveBooleanVal() DataDatabricksWorkspaceSettingV2EffectiveBooleanValOutputReference
	EffectiveDefaultDataSecurityMode() DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityModeOutputReference
	EffectiveDefaultDataSecurityModeInput() interface{}
	EffectiveIntegerVal() DataDatabricksWorkspaceSettingV2EffectiveIntegerValOutputReference
	EffectivePersonalCompute() DataDatabricksWorkspaceSettingV2EffectivePersonalComputeOutputReference
	EffectivePersonalComputeInput() interface{}
	EffectiveRestrictWorkspaceAdmins() DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	EffectiveRestrictWorkspaceAdminsInput() interface{}
	EffectiveStringVal() DataDatabricksWorkspaceSettingV2EffectiveStringValOutputReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IntegerVal() DataDatabricksWorkspaceSettingV2IntegerValOutputReference
	IntegerValInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PersonalCompute() DataDatabricksWorkspaceSettingV2PersonalComputeOutputReference
	PersonalComputeInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RestrictWorkspaceAdmins() DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdminsOutputReference
	RestrictWorkspaceAdminsInput() interface{}
	StringVal() DataDatabricksWorkspaceSettingV2StringValOutputReference
	StringValInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy)
	PutAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains)
	PutAutomaticClusterUpdateWorkspace(value *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspace)
	PutBooleanVal(value *DataDatabricksWorkspaceSettingV2BooleanVal)
	PutDefaultDataSecurityMode(value *DataDatabricksWorkspaceSettingV2DefaultDataSecurityMode)
	PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy)
	PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains)
	PutEffectiveAutomaticClusterUpdateWorkspace(value *DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace)
	PutEffectiveDefaultDataSecurityMode(value *DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityMode)
	PutEffectivePersonalCompute(value *DataDatabricksWorkspaceSettingV2EffectivePersonalCompute)
	PutEffectiveRestrictWorkspaceAdmins(value *DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdmins)
	PutIntegerVal(value *DataDatabricksWorkspaceSettingV2IntegerVal)
	PutPersonalCompute(value *DataDatabricksWorkspaceSettingV2PersonalCompute)
	PutRestrictWorkspaceAdmins(value *DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdmins)
	PutStringVal(value *DataDatabricksWorkspaceSettingV2StringVal)
	ResetAibiDashboardEmbeddingAccessPolicy()
	ResetAibiDashboardEmbeddingApprovedDomains()
	ResetAutomaticClusterUpdateWorkspace()
	ResetBooleanVal()
	ResetDefaultDataSecurityMode()
	ResetEffectiveAibiDashboardEmbeddingAccessPolicy()
	ResetEffectiveAibiDashboardEmbeddingApprovedDomains()
	ResetEffectiveAutomaticClusterUpdateWorkspace()
	ResetEffectiveDefaultDataSecurityMode()
	ResetEffectivePersonalCompute()
	ResetEffectiveRestrictWorkspaceAdmins()
	ResetIntegerVal()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersonalCompute()
	ResetRestrictWorkspaceAdmins()
	ResetStringVal()
	ResetWorkspaceId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatabricksWorkspaceSettingV2
type jsiiProxy_DataDatabricksWorkspaceSettingV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AibiDashboardEmbeddingAccessPolicy() DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AibiDashboardEmbeddingApprovedDomains() DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AutomaticClusterUpdateWorkspace() DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference {
	var returns DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) AutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) BooleanVal() DataDatabricksWorkspaceSettingV2BooleanValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2BooleanValOutputReference
	_jsii_.Get(
		j,
		"booleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) BooleanValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) DefaultDataSecurityMode() DataDatabricksWorkspaceSettingV2DefaultDataSecurityModeOutputReference {
	var returns DataDatabricksWorkspaceSettingV2DefaultDataSecurityModeOutputReference
	_jsii_.Get(
		j,
		"defaultDataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) DefaultDataSecurityModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultDataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicy() DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomains() DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAutomaticClusterUpdateWorkspace() DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveAutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveBooleanVal() DataDatabricksWorkspaceSettingV2EffectiveBooleanValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveBooleanValOutputReference
	_jsii_.Get(
		j,
		"effectiveBooleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveDefaultDataSecurityMode() DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityModeOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityModeOutputReference
	_jsii_.Get(
		j,
		"effectiveDefaultDataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveDefaultDataSecurityModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveDefaultDataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveIntegerVal() DataDatabricksWorkspaceSettingV2EffectiveIntegerValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveIntegerValOutputReference
	_jsii_.Get(
		j,
		"effectiveIntegerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectivePersonalCompute() DataDatabricksWorkspaceSettingV2EffectivePersonalComputeOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectivePersonalComputeOutputReference
	_jsii_.Get(
		j,
		"effectivePersonalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectivePersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectivePersonalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveRestrictWorkspaceAdmins() DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveRestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) EffectiveStringVal() DataDatabricksWorkspaceSettingV2EffectiveStringValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2EffectiveStringValOutputReference
	_jsii_.Get(
		j,
		"effectiveStringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) IntegerVal() DataDatabricksWorkspaceSettingV2IntegerValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2IntegerValOutputReference
	_jsii_.Get(
		j,
		"integerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) IntegerValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) PersonalCompute() DataDatabricksWorkspaceSettingV2PersonalComputeOutputReference {
	var returns DataDatabricksWorkspaceSettingV2PersonalComputeOutputReference
	_jsii_.Get(
		j,
		"personalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) PersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) RestrictWorkspaceAdmins() DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdminsOutputReference {
	var returns DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"restrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) RestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) StringVal() DataDatabricksWorkspaceSettingV2StringValOutputReference {
	var returns DataDatabricksWorkspaceSettingV2StringValOutputReference
	_jsii_.Get(
		j,
		"stringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) StringValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2 databricks_workspace_setting_v2} Data Source.
func NewDataDatabricksWorkspaceSettingV2(scope constructs.Construct, id *string, config *DataDatabricksWorkspaceSettingV2Config) DataDatabricksWorkspaceSettingV2 {
	_init_.Initialize()

	if err := validateNewDataDatabricksWorkspaceSettingV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksWorkspaceSettingV2{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/workspace_setting_v2 databricks_workspace_setting_v2} Data Source.
func NewDataDatabricksWorkspaceSettingV2_Override(d DataDatabricksWorkspaceSettingV2, scope constructs.Construct, id *string, config *DataDatabricksWorkspaceSettingV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksWorkspaceSettingV2)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksWorkspaceSettingV2 resource upon running "cdktf plan <stack-name>".
func DataDatabricksWorkspaceSettingV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksWorkspaceSettingV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataDatabricksWorkspaceSettingV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksWorkspaceSettingV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksWorkspaceSettingV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksWorkspaceSettingV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksWorkspaceSettingV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksWorkspaceSettingV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksWorkspaceSettingV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksWorkspaceSettingV2.DataDatabricksWorkspaceSettingV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingAccessPolicy) {
	if err := d.validatePutAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksWorkspaceSettingV2AibiDashboardEmbeddingApprovedDomains) {
	if err := d.validatePutAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutAutomaticClusterUpdateWorkspace(value *DataDatabricksWorkspaceSettingV2AutomaticClusterUpdateWorkspace) {
	if err := d.validatePutAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutBooleanVal(value *DataDatabricksWorkspaceSettingV2BooleanVal) {
	if err := d.validatePutBooleanValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBooleanVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutDefaultDataSecurityMode(value *DataDatabricksWorkspaceSettingV2DefaultDataSecurityMode) {
	if err := d.validatePutDefaultDataSecurityModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultDataSecurityMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy) {
	if err := d.validatePutEffectiveAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksWorkspaceSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains) {
	if err := d.validatePutEffectiveAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectiveAutomaticClusterUpdateWorkspace(value *DataDatabricksWorkspaceSettingV2EffectiveAutomaticClusterUpdateWorkspace) {
	if err := d.validatePutEffectiveAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectiveDefaultDataSecurityMode(value *DataDatabricksWorkspaceSettingV2EffectiveDefaultDataSecurityMode) {
	if err := d.validatePutEffectiveDefaultDataSecurityModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveDefaultDataSecurityMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectivePersonalCompute(value *DataDatabricksWorkspaceSettingV2EffectivePersonalCompute) {
	if err := d.validatePutEffectivePersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectivePersonalCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutEffectiveRestrictWorkspaceAdmins(value *DataDatabricksWorkspaceSettingV2EffectiveRestrictWorkspaceAdmins) {
	if err := d.validatePutEffectiveRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutIntegerVal(value *DataDatabricksWorkspaceSettingV2IntegerVal) {
	if err := d.validatePutIntegerValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIntegerVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutPersonalCompute(value *DataDatabricksWorkspaceSettingV2PersonalCompute) {
	if err := d.validatePutPersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersonalCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutRestrictWorkspaceAdmins(value *DataDatabricksWorkspaceSettingV2RestrictWorkspaceAdmins) {
	if err := d.validatePutRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) PutStringVal(value *DataDatabricksWorkspaceSettingV2StringVal) {
	if err := d.validatePutStringValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStringVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetBooleanVal() {
	_jsii_.InvokeVoid(
		d,
		"resetBooleanVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetDefaultDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectiveAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectiveAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectiveAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectiveDefaultDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveDefaultDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectivePersonalCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectivePersonalCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetEffectiveRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetIntegerVal() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegerVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetPersonalCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetPersonalCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		d,
		"resetRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetStringVal() {
	_jsii_.InvokeVoid(
		d,
		"resetStringVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksWorkspaceSettingV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

