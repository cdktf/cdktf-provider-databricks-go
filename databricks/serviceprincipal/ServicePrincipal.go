// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/serviceprincipal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/service_principal databricks_service_principal}.
type ServicePrincipal interface {
	cdktf.TerraformResource
	AclPrincipalId() *string
	SetAclPrincipalId(val *string)
	AclPrincipalIdInput() *string
	Active() interface{}
	SetActive(val interface{})
	ActiveInput() interface{}
	AllowClusterCreate() interface{}
	SetAllowClusterCreate(val interface{})
	AllowClusterCreateInput() interface{}
	AllowInstancePoolCreate() interface{}
	SetAllowInstancePoolCreate(val interface{})
	AllowInstancePoolCreateInput() interface{}
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
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
	DatabricksSqlAccess() interface{}
	SetDatabricksSqlAccess(val interface{})
	DatabricksSqlAccessInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableAsUserDeletion() interface{}
	SetDisableAsUserDeletion(val interface{})
	DisableAsUserDeletionInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	Force() interface{}
	SetForce(val interface{})
	ForceDeleteHomeDir() interface{}
	SetForceDeleteHomeDir(val interface{})
	ForceDeleteHomeDirInput() interface{}
	ForceDeleteRepos() interface{}
	SetForceDeleteRepos(val interface{})
	ForceDeleteReposInput() interface{}
	ForceInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Home() *string
	SetHome(val *string)
	HomeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	Repos() *string
	SetRepos(val *string)
	ReposInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkspaceAccess() interface{}
	SetWorkspaceAccess(val interface{})
	WorkspaceAccessInput() interface{}
	WorkspaceConsume() interface{}
	SetWorkspaceConsume(val interface{})
	WorkspaceConsumeInput() interface{}
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
	ResetAclPrincipalId()
	ResetActive()
	ResetAllowClusterCreate()
	ResetAllowInstancePoolCreate()
	ResetApplicationId()
	ResetDatabricksSqlAccess()
	ResetDisableAsUserDeletion()
	ResetDisplayName()
	ResetExternalId()
	ResetForce()
	ResetForceDeleteHomeDir()
	ResetForceDeleteRepos()
	ResetHome()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRepos()
	ResetWorkspaceAccess()
	ResetWorkspaceConsume()
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

// The jsii proxy struct for ServicePrincipal
type jsiiProxy_ServicePrincipal struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicePrincipal) AclPrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AclPrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AllowClusterCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClusterCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AllowClusterCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClusterCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AllowInstancePoolCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstancePoolCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AllowInstancePoolCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstancePoolCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DatabricksSqlAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksSqlAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DatabricksSqlAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksSqlAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DisableAsUserDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAsUserDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DisableAsUserDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAsUserDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Force() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForceDeleteHomeDir() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteHomeDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForceDeleteHomeDirInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteHomeDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForceDeleteRepos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteRepos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForceDeleteReposInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteReposInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Home() *string {
	var returns *string
	_jsii_.Get(
		j,
		"home",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) HomeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Repos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ReposInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reposInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) WorkspaceAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) WorkspaceAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) WorkspaceConsume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceConsume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) WorkspaceConsumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceConsumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/service_principal databricks_service_principal} Resource.
func NewServicePrincipal(scope constructs.Construct, id *string, config *ServicePrincipalConfig) ServicePrincipal {
	_init_.Initialize()

	if err := validateNewServicePrincipalParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipal{}

	_jsii_.Create(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/service_principal databricks_service_principal} Resource.
func NewServicePrincipal_Override(s ServicePrincipal, scope constructs.Construct, id *string, config *ServicePrincipalConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAclPrincipalId(val *string) {
	if err := j.validateSetAclPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclPrincipalId",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAllowClusterCreate(val interface{}) {
	if err := j.validateSetAllowClusterCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClusterCreate",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAllowInstancePoolCreate(val interface{}) {
	if err := j.validateSetAllowInstancePoolCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInstancePoolCreate",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDatabricksSqlAccess(val interface{}) {
	if err := j.validateSetDatabricksSqlAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksSqlAccess",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDisableAsUserDeletion(val interface{}) {
	if err := j.validateSetDisableAsUserDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAsUserDeletion",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetForce(val interface{}) {
	if err := j.validateSetForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"force",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetForceDeleteHomeDir(val interface{}) {
	if err := j.validateSetForceDeleteHomeDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteHomeDir",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetForceDeleteRepos(val interface{}) {
	if err := j.validateSetForceDeleteReposParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteRepos",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetHome(val *string) {
	if err := j.validateSetHomeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"home",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetRepos(val *string) {
	if err := j.validateSetReposParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repos",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetWorkspaceAccess(val interface{}) {
	if err := j.validateSetWorkspaceAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceAccess",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetWorkspaceConsume(val interface{}) {
	if err := j.validateSetWorkspaceConsumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceConsume",
		val,
	)
}

// Generates CDKTF code for importing a ServicePrincipal resource upon running "cdktf plan <stack-name>".
func ServicePrincipal_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicePrincipal_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
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
func ServicePrincipal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicePrincipal_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicePrincipal_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicePrincipal_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.servicePrincipal.ServicePrincipal",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicePrincipal) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicePrincipal) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicePrincipal) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicePrincipal) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicePrincipal) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicePrincipal) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicePrincipal) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAclPrincipalId() {
	_jsii_.InvokeVoid(
		s,
		"resetAclPrincipalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetActive() {
	_jsii_.InvokeVoid(
		s,
		"resetActive",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAllowClusterCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowClusterCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAllowInstancePoolCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowInstancePoolCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetApplicationId() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetDatabricksSqlAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetDatabricksSqlAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetDisableAsUserDeletion() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableAsUserDeletion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetExternalId() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetForce() {
	_jsii_.InvokeVoid(
		s,
		"resetForce",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetForceDeleteHomeDir() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDeleteHomeDir",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetForceDeleteRepos() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDeleteRepos",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetHome() {
	_jsii_.InvokeVoid(
		s,
		"resetHome",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetRepos() {
	_jsii_.InvokeVoid(
		s,
		"resetRepos",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetWorkspaceAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkspaceAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetWorkspaceConsume() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkspaceConsume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

