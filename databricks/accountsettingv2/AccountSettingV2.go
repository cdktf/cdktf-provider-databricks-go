// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountsettingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/accountsettingv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/account_setting_v2 databricks_account_setting_v2}.
type AccountSettingV2 interface {
	cdktf.TerraformResource
	AibiDashboardEmbeddingAccessPolicy() AccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	AibiDashboardEmbeddingAccessPolicyInput() interface{}
	AibiDashboardEmbeddingApprovedDomains() AccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	AibiDashboardEmbeddingApprovedDomainsInput() interface{}
	AutomaticClusterUpdateWorkspace() AccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	AutomaticClusterUpdateWorkspaceInput() interface{}
	BooleanVal() AccountSettingV2BooleanValOutputReference
	BooleanValInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveAibiDashboardEmbeddingAccessPolicy() AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{}
	EffectiveAibiDashboardEmbeddingApprovedDomains() AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{}
	EffectiveAutomaticClusterUpdateWorkspace() AccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	EffectiveAutomaticClusterUpdateWorkspaceInput() interface{}
	EffectiveBooleanVal() AccountSettingV2EffectiveBooleanValOutputReference
	EffectiveIntegerVal() AccountSettingV2EffectiveIntegerValOutputReference
	EffectivePersonalCompute() AccountSettingV2EffectivePersonalComputeOutputReference
	EffectivePersonalComputeInput() interface{}
	EffectiveRestrictWorkspaceAdmins() AccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	EffectiveRestrictWorkspaceAdminsInput() interface{}
	EffectiveStringVal() AccountSettingV2EffectiveStringValOutputReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IntegerVal() AccountSettingV2IntegerValOutputReference
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
	PersonalCompute() AccountSettingV2PersonalComputeOutputReference
	PersonalComputeInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RestrictWorkspaceAdmins() AccountSettingV2RestrictWorkspaceAdminsOutputReference
	RestrictWorkspaceAdminsInput() interface{}
	StringVal() AccountSettingV2StringValOutputReference
	StringValInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAibiDashboardEmbeddingAccessPolicy(value *AccountSettingV2AibiDashboardEmbeddingAccessPolicy)
	PutAibiDashboardEmbeddingApprovedDomains(value *AccountSettingV2AibiDashboardEmbeddingApprovedDomains)
	PutAutomaticClusterUpdateWorkspace(value *AccountSettingV2AutomaticClusterUpdateWorkspace)
	PutBooleanVal(value *AccountSettingV2BooleanVal)
	PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy)
	PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains)
	PutEffectiveAutomaticClusterUpdateWorkspace(value *AccountSettingV2EffectiveAutomaticClusterUpdateWorkspace)
	PutEffectivePersonalCompute(value *AccountSettingV2EffectivePersonalCompute)
	PutEffectiveRestrictWorkspaceAdmins(value *AccountSettingV2EffectiveRestrictWorkspaceAdmins)
	PutIntegerVal(value *AccountSettingV2IntegerVal)
	PutPersonalCompute(value *AccountSettingV2PersonalCompute)
	PutRestrictWorkspaceAdmins(value *AccountSettingV2RestrictWorkspaceAdmins)
	PutStringVal(value *AccountSettingV2StringVal)
	ResetAibiDashboardEmbeddingAccessPolicy()
	ResetAibiDashboardEmbeddingApprovedDomains()
	ResetAutomaticClusterUpdateWorkspace()
	ResetBooleanVal()
	ResetEffectiveAibiDashboardEmbeddingAccessPolicy()
	ResetEffectiveAibiDashboardEmbeddingApprovedDomains()
	ResetEffectiveAutomaticClusterUpdateWorkspace()
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

// The jsii proxy struct for AccountSettingV2
type jsiiProxy_AccountSettingV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccountSettingV2) AibiDashboardEmbeddingAccessPolicy() AccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns AccountSettingV2AibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) AibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) AibiDashboardEmbeddingApprovedDomains() AccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns AccountSettingV2AibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) AibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) AutomaticClusterUpdateWorkspace() AccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference {
	var returns AccountSettingV2AutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) AutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) BooleanVal() AccountSettingV2BooleanValOutputReference {
	var returns AccountSettingV2BooleanValOutputReference
	_jsii_.Get(
		j,
		"booleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) BooleanValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicy() AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference {
	var returns AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAibiDashboardEmbeddingAccessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomains() AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference {
	var returns AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomainsOutputReference
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAibiDashboardEmbeddingApprovedDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAibiDashboardEmbeddingApprovedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAutomaticClusterUpdateWorkspace() AccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference {
	var returns AccountSettingV2EffectiveAutomaticClusterUpdateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveAutomaticClusterUpdateWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveAutomaticClusterUpdateWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveBooleanVal() AccountSettingV2EffectiveBooleanValOutputReference {
	var returns AccountSettingV2EffectiveBooleanValOutputReference
	_jsii_.Get(
		j,
		"effectiveBooleanVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveIntegerVal() AccountSettingV2EffectiveIntegerValOutputReference {
	var returns AccountSettingV2EffectiveIntegerValOutputReference
	_jsii_.Get(
		j,
		"effectiveIntegerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectivePersonalCompute() AccountSettingV2EffectivePersonalComputeOutputReference {
	var returns AccountSettingV2EffectivePersonalComputeOutputReference
	_jsii_.Get(
		j,
		"effectivePersonalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectivePersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectivePersonalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveRestrictWorkspaceAdmins() AccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference {
	var returns AccountSettingV2EffectiveRestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveRestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"effectiveRestrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) EffectiveStringVal() AccountSettingV2EffectiveStringValOutputReference {
	var returns AccountSettingV2EffectiveStringValOutputReference
	_jsii_.Get(
		j,
		"effectiveStringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) IntegerVal() AccountSettingV2IntegerValOutputReference {
	var returns AccountSettingV2IntegerValOutputReference
	_jsii_.Get(
		j,
		"integerVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) IntegerValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) PersonalCompute() AccountSettingV2PersonalComputeOutputReference {
	var returns AccountSettingV2PersonalComputeOutputReference
	_jsii_.Get(
		j,
		"personalCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) PersonalComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) RestrictWorkspaceAdmins() AccountSettingV2RestrictWorkspaceAdminsOutputReference {
	var returns AccountSettingV2RestrictWorkspaceAdminsOutputReference
	_jsii_.Get(
		j,
		"restrictWorkspaceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) RestrictWorkspaceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictWorkspaceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) StringVal() AccountSettingV2StringValOutputReference {
	var returns AccountSettingV2StringValOutputReference
	_jsii_.Get(
		j,
		"stringVal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) StringValInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringValInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountSettingV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/account_setting_v2 databricks_account_setting_v2} Resource.
func NewAccountSettingV2(scope constructs.Construct, id *string, config *AccountSettingV2Config) AccountSettingV2 {
	_init_.Initialize()

	if err := validateNewAccountSettingV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountSettingV2{}

	_jsii_.Create(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.96.0/docs/resources/account_setting_v2 databricks_account_setting_v2} Resource.
func NewAccountSettingV2_Override(a AccountSettingV2, scope constructs.Construct, id *string, config *AccountSettingV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccountSettingV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a AccountSettingV2 resource upon running "cdktf plan <stack-name>".
func AccountSettingV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAccountSettingV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
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
func AccountSettingV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountSettingV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccountSettingV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountSettingV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccountSettingV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountSettingV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccountSettingV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.accountSettingV2.AccountSettingV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccountSettingV2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AccountSettingV2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccountSettingV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccountSettingV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccountSettingV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccountSettingV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccountSettingV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccountSettingV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccountSettingV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccountSettingV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccountSettingV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccountSettingV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AccountSettingV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccountSettingV2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AccountSettingV2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccountSettingV2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutAibiDashboardEmbeddingAccessPolicy(value *AccountSettingV2AibiDashboardEmbeddingAccessPolicy) {
	if err := a.validatePutAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutAibiDashboardEmbeddingApprovedDomains(value *AccountSettingV2AibiDashboardEmbeddingApprovedDomains) {
	if err := a.validatePutAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutAutomaticClusterUpdateWorkspace(value *AccountSettingV2AutomaticClusterUpdateWorkspace) {
	if err := a.validatePutAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutBooleanVal(value *AccountSettingV2BooleanVal) {
	if err := a.validatePutBooleanValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBooleanVal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutEffectiveAibiDashboardEmbeddingAccessPolicy(value *AccountSettingV2EffectiveAibiDashboardEmbeddingAccessPolicy) {
	if err := a.validatePutEffectiveAibiDashboardEmbeddingAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEffectiveAibiDashboardEmbeddingAccessPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutEffectiveAibiDashboardEmbeddingApprovedDomains(value *AccountSettingV2EffectiveAibiDashboardEmbeddingApprovedDomains) {
	if err := a.validatePutEffectiveAibiDashboardEmbeddingApprovedDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEffectiveAibiDashboardEmbeddingApprovedDomains",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutEffectiveAutomaticClusterUpdateWorkspace(value *AccountSettingV2EffectiveAutomaticClusterUpdateWorkspace) {
	if err := a.validatePutEffectiveAutomaticClusterUpdateWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEffectiveAutomaticClusterUpdateWorkspace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutEffectivePersonalCompute(value *AccountSettingV2EffectivePersonalCompute) {
	if err := a.validatePutEffectivePersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEffectivePersonalCompute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutEffectiveRestrictWorkspaceAdmins(value *AccountSettingV2EffectiveRestrictWorkspaceAdmins) {
	if err := a.validatePutEffectiveRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEffectiveRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutIntegerVal(value *AccountSettingV2IntegerVal) {
	if err := a.validatePutIntegerValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegerVal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutPersonalCompute(value *AccountSettingV2PersonalCompute) {
	if err := a.validatePutPersonalComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPersonalCompute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutRestrictWorkspaceAdmins(value *AccountSettingV2RestrictWorkspaceAdmins) {
	if err := a.validatePutRestrictWorkspaceAdminsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestrictWorkspaceAdmins",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) PutStringVal(value *AccountSettingV2StringVal) {
	if err := a.validatePutStringValParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringVal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetBooleanVal() {
	_jsii_.InvokeVoid(
		a,
		"resetBooleanVal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetEffectiveAibiDashboardEmbeddingAccessPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveAibiDashboardEmbeddingAccessPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetEffectiveAibiDashboardEmbeddingApprovedDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveAibiDashboardEmbeddingApprovedDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetEffectiveAutomaticClusterUpdateWorkspace() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveAutomaticClusterUpdateWorkspace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetEffectivePersonalCompute() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectivePersonalCompute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetEffectiveRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		a,
		"resetEffectiveRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetIntegerVal() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegerVal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetPersonalCompute() {
	_jsii_.InvokeVoid(
		a,
		"resetPersonalCompute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetRestrictWorkspaceAdmins() {
	_jsii_.InvokeVoid(
		a,
		"resetRestrictWorkspaceAdmins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) ResetStringVal() {
	_jsii_.InvokeVoid(
		a,
		"resetStringVal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountSettingV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountSettingV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

