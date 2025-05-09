// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v13/datadatabricksgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/group databricks_group}.
type DataDatabricksGroup interface {
	cdktf.TerraformDataSource
	AclPrincipalId() *string
	SetAclPrincipalId(val *string)
	AclPrincipalIdInput() *string
	AllowClusterCreate() interface{}
	SetAllowClusterCreate(val interface{})
	AllowClusterCreateInput() interface{}
	AllowInstancePoolCreate() interface{}
	SetAllowInstancePoolCreate(val interface{})
	AllowInstancePoolCreateInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChildGroups() *[]*string
	SetChildGroups(val *[]*string)
	ChildGroupsInput() *[]*string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceProfiles() *[]*string
	SetInstanceProfiles(val *[]*string)
	InstanceProfilesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Recursive() interface{}
	SetRecursive(val interface{})
	RecursiveInput() interface{}
	ServicePrincipals() *[]*string
	SetServicePrincipals(val *[]*string)
	ServicePrincipalsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
	WorkspaceAccess() interface{}
	SetWorkspaceAccess(val interface{})
	WorkspaceAccessInput() interface{}
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
	ResetAclPrincipalId()
	ResetAllowClusterCreate()
	ResetAllowInstancePoolCreate()
	ResetChildGroups()
	ResetDatabricksSqlAccess()
	ResetExternalId()
	ResetGroups()
	ResetId()
	ResetInstanceProfiles()
	ResetMembers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecursive()
	ResetServicePrincipals()
	ResetUsers()
	ResetWorkspaceAccess()
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

// The jsii proxy struct for DataDatabricksGroup
type jsiiProxy_DataDatabricksGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatabricksGroup) AclPrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) AclPrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclPrincipalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) AllowClusterCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClusterCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) AllowClusterCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowClusterCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) AllowInstancePoolCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstancePoolCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) AllowInstancePoolCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstancePoolCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ChildGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ChildGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) DatabricksSqlAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksSqlAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) DatabricksSqlAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databricksSqlAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) InstanceProfiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) InstanceProfilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Recursive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recursive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) RecursiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recursiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ServicePrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicePrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) ServicePrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicePrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) WorkspaceAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatabricksGroup) WorkspaceAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceAccessInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/group databricks_group} Data Source.
func NewDataDatabricksGroup(scope constructs.Construct, id *string, config *DataDatabricksGroupConfig) DataDatabricksGroup {
	_init_.Initialize()

	if err := validateNewDataDatabricksGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatabricksGroup{}

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/group databricks_group} Data Source.
func NewDataDatabricksGroup_Override(d DataDatabricksGroup, scope constructs.Construct, id *string, config *DataDatabricksGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetAclPrincipalId(val *string) {
	if err := j.validateSetAclPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclPrincipalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetAllowClusterCreate(val interface{}) {
	if err := j.validateSetAllowClusterCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowClusterCreate",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetAllowInstancePoolCreate(val interface{}) {
	if err := j.validateSetAllowInstancePoolCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInstancePoolCreate",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetChildGroups(val *[]*string) {
	if err := j.validateSetChildGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childGroups",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetDatabricksSqlAccess(val interface{}) {
	if err := j.validateSetDatabricksSqlAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksSqlAccess",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetInstanceProfiles(val *[]*string) {
	if err := j.validateSetInstanceProfilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfiles",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetRecursive(val interface{}) {
	if err := j.validateSetRecursiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recursive",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetServicePrincipals(val *[]*string) {
	if err := j.validateSetServicePrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipals",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetUsers(val *[]*string) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

func (j *jsiiProxy_DataDatabricksGroup)SetWorkspaceAccess(val interface{}) {
	if err := j.validateSetWorkspaceAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceAccess",
		val,
	)
}

// Generates CDKTF code for importing a DataDatabricksGroup resource upon running "cdktf plan <stack-name>".
func DataDatabricksGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatabricksGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
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
func DataDatabricksGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksGroup_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksGroup_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatabricksGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatabricksGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatabricksGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.dataDatabricksGroup.DataDatabricksGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatabricksGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatabricksGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatabricksGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatabricksGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatabricksGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatabricksGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatabricksGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatabricksGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatabricksGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatabricksGroup) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetAclPrincipalId() {
	_jsii_.InvokeVoid(
		d,
		"resetAclPrincipalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetAllowClusterCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowClusterCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetAllowInstancePoolCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowInstancePoolCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetChildGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetChildGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetDatabricksSqlAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabricksSqlAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetExternalId() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetInstanceProfiles() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetMembers() {
	_jsii_.InvokeVoid(
		d,
		"resetMembers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetRecursive() {
	_jsii_.InvokeVoid(
		d,
		"resetRecursive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetServicePrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetUsers() {
	_jsii_.InvokeVoid(
		d,
		"resetUsers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) ResetWorkspaceAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatabricksGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatabricksGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

