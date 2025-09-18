// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/datadatabricksaccountsettingv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_setting_v2 databricks_account_setting_v2}.
type DataDatabricksAccountSettingV2 interface {
	cdktf.TerraformDataSource
	AibiDashboardEmbeddingAccessPolicy() DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	AibiDashboardEmbeddingAccessPolicyInput() interface{}
	AibiDashboardEmbeddingApprovedDomains() DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	AibiDashboardEmbeddingApprovedDomainsInput() interface{}
	AutomaticClusterUpdateWorkspace() DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	AutomaticClusterUpdateWorkspaceInput() interface{}
	BooleanVal() DataDatabricksAccountSettingV2BooleanValOutputReference
	BooleanValInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultDataSecurityMode() DataDatabricksAccountSettingV2DefaultDataSecurityModeOutputReference
	DefaultDataSecurityModeInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveAibiDashboardEmbeddingAccessPolicy() DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{}
	EffectiveAibiDashboardEmbeddingApprovedDomains() DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{}
	EffectiveAutomaticClusterUpdateWorkspace() DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	EffectiveAutomaticClusterUpdateWorkspaceInput() interface{}
	EffectiveBooleanVal() DataDatabricksAccountSettingV2EffectiveBooleanValOutputReference
	EffectiveDefaultDataSecurityMode() DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityModeOutputReference
	EffectiveDefaultDataSecurityModeInput() interface{}
	EffectiveIntegerVal() DataDatabricksAccountSettingV2EffectiveIntegerValOutputReference
	EffectivePersonalCompute() DataDatabricksAccountSettingV2EffectivePersonalComputeOutputReference
	EffectivePersonalComputeInput() interface{}
	EffectiveRestrictWorkspaceAdmins() DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	EffectiveRestrictWorkspaceAdminsInput() interface{}
	EffectiveStringVal() DataDatabricksAccountSettingV2EffectiveStringValOutputReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IntegerVal() DataDatabricksAccountSettingV2IntegerValOutputReference
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
	PersonalCompute() DataDatabricksAccountSettingV2PersonalComputeOutputReference
	PersonalComputeInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RestrictWorkspaceAdmins() DataDatabricksAccountSettingV2RestrictWorkspaceAdminsOutputReference
	RestrictWorkspaceAdminsInput() interface{}
	StringVal() DataDatabricksAccountSettingV2StringValOutputReference
	StringValInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicy)
	PutAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomains)
	PutAutomaticClusterUpdateWorkspace(value *DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspace)
	PutBooleanVal(value *DataDatabricksAccountSettingV2BooleanVal)
	PutDefaultDataSecurityMode(value *DataDatabricksAccountSettingV2DefaultDataSecurityMode)
	PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy)
	PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains)
	PutEffectiveAutomaticClusterUpdateWorkspace(value *DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspace)
	PutEffectiveDefaultDataSecurityMode(value *DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityMode)
	PutEffectivePersonalCompute(value *DataDatabricksAccountSettingV2EffectivePersonalCompute)
	PutEffectiveRestrictWorkspaceAdmins(value *DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdmins)
	PutIntegerVal(value *DataDatabricksAccountSettingV2IntegerVal)
	PutPersonalCompute(value *DataDatabricksAccountSettingV2PersonalCompute)
	PutRestrictWorkspaceAdmins(value *DataDatabricksAccountSettingV2RestrictWorkspaceAdmins)
	PutStringVal(value *DataDatabricksAccountSettingV2StringVal)
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

// The jsii proxy struct for DataDatabricksAccountSettingV2
type jsiiProxy_DataDatabricksAccountSettingV2 struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AibiDashboardEmbeddingAccessPolicy() DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AibiDashboardEmbeddingApprovedDomains() DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AutomaticClusterUpdateWorkspace() DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference {
	var returns DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) AutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) BooleanVal() DataDatabricksAccountSettingV2BooleanValOutputReference {
	var returns DataDatabricksAccountSettingV2BooleanValOutputReference
	_jsii_.Get(
		j,
		"booleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) BooleanValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) DefaultDataSecurityMode() DataDatabricksAccountSettingV2DefaultDataSecurityModeOutputReference {
	var returns DataDatabricksAccountSettingV2DefaultDataSecurityModeOutputReference
	_jsii_.Get(
		j,
		"defaultDataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) DefaultDataSecurityModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultDataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicy() DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomains() DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAutomaticClusterUpdateWorkspace() DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveAutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveBooleanVal() DataDatabricksAccountSettingV2EffectiveBooleanValOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveBooleanValOutputReference
	_jsii_.Get(
		j,
		"effectiveBooleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveDefaultDataSecurityMode() DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityModeOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityModeOutputReference
	_jsii_.Get(
		j,
		"effectiveDefaultDataSecurityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveDefaultDataSecurityModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveDefaultDataSecurityModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveIntegerVal() DataDatabricksAccountSettingV2EffectiveIntegerValOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveIntegerValOutputReference
	_jsii_.Get(
		j,
		"effectiveIntegerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectivePersonalCompute() DataDatabricksAccountSettingV2EffectivePersonalComputeOutputReference {
	var returns DataDatabricksAccountSettingV2EffectivePersonalComputeOutputReference
	_jsii_.Get(
		j,
		"effectivePersonalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectivePersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectivePersonalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveRestrictWorkspaceAdmins() DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveRestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) EffectiveStringVal() DataDatabricksAccountSettingV2EffectiveStringValOutputReference {
	var returns DataDatabricksAccountSettingV2EffectiveStringValOutputReference
	_jsii_.Get(
		j,
		"effectiveStringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) IntegerVal() DataDatabricksAccountSettingV2IntegerValOutputReference {
	var returns DataDatabricksAccountSettingV2IntegerValOutputReference
	_jsii_.Get(
		j,
		"integerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) IntegerValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) PersonalCompute() DataDatabricksAccountSettingV2PersonalComputeOutputReference {
	var returns DataDatabricksAccountSettingV2PersonalComputeOutputReference
	_jsii_.Get(
		j,
		"personalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) PersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) RestrictWorkspaceAdmins() DataDatabricksAccountSettingV2RestrictWorkspaceAdminsOutputReference {
	var returns DataDatabricksAccountSettingV2RestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"restrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) RestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) StringVal() DataDatabricksAccountSettingV2StringValOutputReference {
	var returns DataDatabricksAccountSettingV2StringValOutputReference
	_jsii_.Get(
		j,
		"stringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) StringValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_setting_v2 databricks_account_setting_v2} Data Source.
func NewDataDatabricksAccountSettingV2(scope constructs.Construct, id *string, config *DataDatabricksAccountSettingV2Config) DataDatabricksAccountSettingV2 {
	_init_.Initialize()

	if err := validateNewDataDatabricksAccountSettingV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksAccountSettingV2{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.90.0/docs/data-sources/account_setting_v2 databricks_account_setting_v2} Data Source.
func NewDataDatabricksAccountSettingV2_Override(d DataDatabricksAccountSettingV2, scope constructs.Construct, id *string, config *DataDatabricksAccountSettingV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksAccountSettingV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksAccountSettingV2 resource upon running "cdktf plan <stack-name>".
func DataDatabricksAccountSettingV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksAccountSettingV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
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
func DataDatabricksAccountSettingV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAccountSettingV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksAccountSettingV2_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAccountSettingV2_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksAccountSettingV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksAccountSettingV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksAccountSettingV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksAccountSettingV2.DataDatabricksAccountSettingV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksAccountSettingV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksAccountSettingV2AibiDashboardEmbeddingAccessPolicy) {
	if err := d.validatePutAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksAccountSettingV2AibiDashboardEmbeddingApprovedDomains) {
	if err := d.validatePutAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutAutomaticClusterUpdateWorkspace(value *DataDatabricksAccountSettingV2AutomaticClusterUpdateWorkspace) {
	if err := d.validatePutAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutBooleanVal(value *DataDatabricksAccountSettingV2BooleanVal) {
	if err := d.validatePutBooleanValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBooleanVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutDefaultDataSecurityMode(value *DataDatabricksAccountSettingV2DefaultDataSecurityMode) {
	if err := d.validatePutDefaultDataSecurityModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultDataSecurityMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy) {
	if err := d.validatePutEffectiveAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *DataDatabricksAccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains) {
	if err := d.validatePutEffectiveAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectiveAutomaticClusterUpdateWorkspace(value *DataDatabricksAccountSettingV2EffectiveAutomaticClusterUpdateWorkspace) {
	if err := d.validatePutEffectiveAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectiveDefaultDataSecurityMode(value *DataDatabricksAccountSettingV2EffectiveDefaultDataSecurityMode) {
	if err := d.validatePutEffectiveDefaultDataSecurityModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveDefaultDataSecurityMode",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectivePersonalCompute(value *DataDatabricksAccountSettingV2EffectivePersonalCompute) {
	if err := d.validatePutEffectivePersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectivePersonalCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutEffectiveRestrictWorkspaceAdmins(value *DataDatabricksAccountSettingV2EffectiveRestrictWorkspaceAdmins) {
	if err := d.validatePutEffectiveRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEffectiveRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutIntegerVal(value *DataDatabricksAccountSettingV2IntegerVal) {
	if err := d.validatePutIntegerValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIntegerVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutPersonalCompute(value *DataDatabricksAccountSettingV2PersonalCompute) {
	if err := d.validatePutPersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersonalCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutRestrictWorkspaceAdmins(value *DataDatabricksAccountSettingV2RestrictWorkspaceAdmins) {
	if err := d.validatePutRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) PutStringVal(value *DataDatabricksAccountSettingV2StringVal) {
	if err := d.validatePutStringValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStringVal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetBooleanVal() {
	_jsii_.InvokeVoid(
		d,
		"resetBooleanVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetDefaultDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectiveAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectiveAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectiveAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectiveDefaultDataSecurityMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveDefaultDataSecurityMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectivePersonalCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectivePersonalCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetEffectiveRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectiveRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetIntegerVal() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegerVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetPersonalCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetPersonalCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		d,
		"resetRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ResetStringVal() {
	_jsii_.InvokeVoid(
		d,
		"resetStringVal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksAccountSettingV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

