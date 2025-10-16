// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-databricks-go/databricks/v15/pipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline databricks_pipeline}.
type Pipeline interface {
	cdktf.TerraformResource
	AllowDuplicateNames() interface{}
	SetAllowDuplicateNames(val interface{})
	AllowDuplicateNamesInput() interface{}
	BudgetPolicyId() *string
	SetBudgetPolicyId(val *string)
	BudgetPolicyIdInput() *string
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
	Cause() *string
	SetCause(val *string)
	CauseInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
	Cluster() PipelineClusterList
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterInput() interface{}
	Configuration() *map[string]*string
	SetConfiguration(val *map[string]*string)
	ConfigurationInput() *map[string]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Continuous() interface{}
	SetContinuous(val interface{})
	ContinuousInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatorUserName() *string
	SetCreatorUserName(val *string)
	CreatorUserNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() PipelineDeploymentOutputReference
	DeploymentInput() *PipelineDeployment
	Development() interface{}
	SetDevelopment(val interface{})
	DevelopmentInput() interface{}
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	Environment() PipelineEnvironmentOutputReference
	EnvironmentInput() *PipelineEnvironment
	EventLog() PipelineEventLogOutputReference
	EventLogInput() *PipelineEventLog
	ExpectedLastModified() *float64
	SetExpectedLastModified(val *float64)
	ExpectedLastModifiedInput() *float64
	Filters() PipelineFiltersOutputReference
	FiltersInput() *PipelineFilters
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayDefinition() PipelineGatewayDefinitionOutputReference
	GatewayDefinitionInput() *PipelineGatewayDefinition
	Health() *string
	SetHealth(val *string)
	HealthInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IngestionDefinition() PipelineIngestionDefinitionOutputReference
	IngestionDefinitionInput() *PipelineIngestionDefinition
	LastModified() *float64
	SetLastModified(val *float64)
	LastModifiedInput() *float64
	LatestUpdates() PipelineLatestUpdatesList
	LatestUpdatesInput() interface{}
	Library() PipelineLibraryList
	LibraryInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notification() PipelineNotificationList
	NotificationInput() interface{}
	Photon() interface{}
	SetPhoton(val interface{})
	PhotonInput() interface{}
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
	RestartWindow() PipelineRestartWindowOutputReference
	RestartWindowInput() *PipelineRestartWindow
	RootPath() *string
	SetRootPath(val *string)
	RootPathInput() *string
	RunAs() PipelineRunAsOutputReference
	RunAsInput() *PipelineRunAs
	RunAsUserName() *string
	SetRunAsUserName(val *string)
	RunAsUserNameInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Serverless() interface{}
	SetServerless(val interface{})
	ServerlessInput() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
	Storage() *string
	SetStorage(val *string)
	StorageInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PipelineTimeoutsOutputReference
	TimeoutsInput() interface{}
	Trigger() PipelineTriggerOutputReference
	TriggerInput() *PipelineTrigger
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutCluster(value interface{})
	PutDeployment(value *PipelineDeployment)
	PutEnvironment(value *PipelineEnvironment)
	PutEventLog(value *PipelineEventLog)
	PutFilters(value *PipelineFilters)
	PutGatewayDefinition(value *PipelineGatewayDefinition)
	PutIngestionDefinition(value *PipelineIngestionDefinition)
	PutLatestUpdates(value interface{})
	PutLibrary(value interface{})
	PutNotification(value interface{})
	PutRestartWindow(value *PipelineRestartWindow)
	PutRunAs(value *PipelineRunAs)
	PutTimeouts(value *PipelineTimeouts)
	PutTrigger(value *PipelineTrigger)
	ResetAllowDuplicateNames()
	ResetBudgetPolicyId()
	ResetCatalog()
	ResetCause()
	ResetChannel()
	ResetCluster()
	ResetClusterId()
	ResetConfiguration()
	ResetContinuous()
	ResetCreatorUserName()
	ResetDeployment()
	ResetDevelopment()
	ResetEdition()
	ResetEnvironment()
	ResetEventLog()
	ResetExpectedLastModified()
	ResetFilters()
	ResetGatewayDefinition()
	ResetHealth()
	ResetId()
	ResetIngestionDefinition()
	ResetLastModified()
	ResetLatestUpdates()
	ResetLibrary()
	ResetName()
	ResetNotification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhoton()
	ResetRestartWindow()
	ResetRootPath()
	ResetRunAs()
	ResetRunAsUserName()
	ResetSchema()
	ResetServerless()
	ResetState()
	ResetStorage()
	ResetTags()
	ResetTarget()
	ResetTimeouts()
	ResetTrigger()
	ResetUrl()
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

// The jsii proxy struct for Pipeline
type jsiiProxy_Pipeline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Pipeline) AllowDuplicateNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) AllowDuplicateNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) BudgetPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) BudgetPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Cause() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CauseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"causeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Cluster() PipelineClusterList {
	var returns PipelineClusterList
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Configuration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Continuous() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ContinuousInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CreatorUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CreatorUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creatorUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Deployment() PipelineDeploymentOutputReference {
	var returns PipelineDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) DeploymentInput() *PipelineDeployment {
	var returns *PipelineDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Development() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"development",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) DevelopmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Environment() PipelineEnvironmentOutputReference {
	var returns PipelineEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) EnvironmentInput() *PipelineEnvironment {
	var returns *PipelineEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) EventLog() PipelineEventLogOutputReference {
	var returns PipelineEventLogOutputReference
	_jsii_.Get(
		j,
		"eventLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) EventLogInput() *PipelineEventLog {
	var returns *PipelineEventLog
	_jsii_.Get(
		j,
		"eventLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ExpectedLastModified() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expectedLastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ExpectedLastModifiedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expectedLastModifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Filters() PipelineFiltersOutputReference {
	var returns PipelineFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) FiltersInput() *PipelineFilters {
	var returns *PipelineFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) GatewayDefinition() PipelineGatewayDefinitionOutputReference {
	var returns PipelineGatewayDefinitionOutputReference
	_jsii_.Get(
		j,
		"gatewayDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) GatewayDefinitionInput() *PipelineGatewayDefinition {
	var returns *PipelineGatewayDefinition
	_jsii_.Get(
		j,
		"gatewayDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) HealthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) IngestionDefinition() PipelineIngestionDefinitionOutputReference {
	var returns PipelineIngestionDefinitionOutputReference
	_jsii_.Get(
		j,
		"ingestionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) IngestionDefinitionInput() *PipelineIngestionDefinition {
	var returns *PipelineIngestionDefinition
	_jsii_.Get(
		j,
		"ingestionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) LastModified() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) LastModifiedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) LatestUpdates() PipelineLatestUpdatesList {
	var returns PipelineLatestUpdatesList
	_jsii_.Get(
		j,
		"latestUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) LatestUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Library() PipelineLibraryList {
	var returns PipelineLibraryList
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) LibraryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Notification() PipelineNotificationList {
	var returns PipelineNotificationList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Photon() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PhotonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"photonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RestartWindow() PipelineRestartWindowOutputReference {
	var returns PipelineRestartWindowOutputReference
	_jsii_.Get(
		j,
		"restartWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RestartWindowInput() *PipelineRestartWindow {
	var returns *PipelineRestartWindow
	_jsii_.Get(
		j,
		"restartWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RootPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RootPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RunAs() PipelineRunAsOutputReference {
	var returns PipelineRunAsOutputReference
	_jsii_.Get(
		j,
		"runAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RunAsInput() *PipelineRunAs {
	var returns *PipelineRunAs
	_jsii_.Get(
		j,
		"runAsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RunAsUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) RunAsUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Serverless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) ServerlessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Storage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) StorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Timeouts() PipelineTimeoutsOutputReference {
	var returns PipelineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Trigger() PipelineTriggerOutputReference {
	var returns PipelineTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) TriggerInput() *PipelineTrigger {
	var returns *PipelineTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline databricks_pipeline} Resource.
func NewPipeline(scope constructs.Construct, id *string, config *PipelineConfig) Pipeline {
	_init_.Initialize()

	if err := validateNewPipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pipeline{}

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/databricks/databricks/1.92.0/docs/resources/pipeline databricks_pipeline} Resource.
func NewPipeline_Override(p Pipeline, scope constructs.Construct, id *string, config *PipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Pipeline)SetAllowDuplicateNames(val interface{}) {
	if err := j.validateSetAllowDuplicateNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicateNames",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetBudgetPolicyId(val *string) {
	if err := j.validateSetBudgetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetPolicyId",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetCause(val *string) {
	if err := j.validateSetCauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cause",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetConfiguration(val *map[string]*string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetContinuous(val interface{}) {
	if err := j.validateSetContinuousParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuous",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetCreatorUserName(val *string) {
	if err := j.validateSetCreatorUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creatorUserName",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetDevelopment(val interface{}) {
	if err := j.validateSetDevelopmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"development",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetExpectedLastModified(val *float64) {
	if err := j.validateSetExpectedLastModifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedLastModified",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetHealth(val *string) {
	if err := j.validateSetHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"health",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetLastModified(val *float64) {
	if err := j.validateSetLastModifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModified",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetPhoton(val interface{}) {
	if err := j.validateSetPhotonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"photon",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetRootPath(val *string) {
	if err := j.validateSetRootPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPath",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetRunAsUserName(val *string) {
	if err := j.validateSetRunAsUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUserName",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetServerless(val interface{}) {
	if err := j.validateSetServerlessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverless",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetStorage(val *string) {
	if err := j.validateSetStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storage",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_Pipeline)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTF code for importing a Pipeline resource upon running "cdktf plan <stack-name>".
func Pipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.pipeline.Pipeline",
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
func Pipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Pipeline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipeline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Pipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-databricks.pipeline.Pipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Pipeline) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_Pipeline) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Pipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_Pipeline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pipeline) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_Pipeline) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Pipeline) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Pipeline) PutCluster(value interface{}) {
	if err := p.validatePutClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCluster",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutDeployment(value *PipelineDeployment) {
	if err := p.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDeployment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutEnvironment(value *PipelineEnvironment) {
	if err := p.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutEventLog(value *PipelineEventLog) {
	if err := p.validatePutEventLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventLog",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutFilters(value *PipelineFilters) {
	if err := p.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutGatewayDefinition(value *PipelineGatewayDefinition) {
	if err := p.validatePutGatewayDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGatewayDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutIngestionDefinition(value *PipelineIngestionDefinition) {
	if err := p.validatePutIngestionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIngestionDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutLatestUpdates(value interface{}) {
	if err := p.validatePutLatestUpdatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLatestUpdates",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutLibrary(value interface{}) {
	if err := p.validatePutLibraryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLibrary",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutNotification(value interface{}) {
	if err := p.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNotification",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutRestartWindow(value *PipelineRestartWindow) {
	if err := p.validatePutRestartWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRestartWindow",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutRunAs(value *PipelineRunAs) {
	if err := p.validatePutRunAsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunAs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutTimeouts(value *PipelineTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) PutTrigger(value *PipelineTrigger) {
	if err := p.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTrigger",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Pipeline) ResetAllowDuplicateNames() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowDuplicateNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetBudgetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetBudgetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetCatalog() {
	_jsii_.InvokeVoid(
		p,
		"resetCatalog",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetCause() {
	_jsii_.InvokeVoid(
		p,
		"resetCause",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetChannel() {
	_jsii_.InvokeVoid(
		p,
		"resetChannel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetCluster() {
	_jsii_.InvokeVoid(
		p,
		"resetCluster",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetClusterId() {
	_jsii_.InvokeVoid(
		p,
		"resetClusterId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetConfiguration() {
	_jsii_.InvokeVoid(
		p,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetContinuous() {
	_jsii_.InvokeVoid(
		p,
		"resetContinuous",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetCreatorUserName() {
	_jsii_.InvokeVoid(
		p,
		"resetCreatorUserName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetDeployment() {
	_jsii_.InvokeVoid(
		p,
		"resetDeployment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetDevelopment() {
	_jsii_.InvokeVoid(
		p,
		"resetDevelopment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetEdition() {
	_jsii_.InvokeVoid(
		p,
		"resetEdition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetEventLog() {
	_jsii_.InvokeVoid(
		p,
		"resetEventLog",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetExpectedLastModified() {
	_jsii_.InvokeVoid(
		p,
		"resetExpectedLastModified",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetGatewayDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetGatewayDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetHealth() {
	_jsii_.InvokeVoid(
		p,
		"resetHealth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetIngestionDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetIngestionDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetLastModified() {
	_jsii_.InvokeVoid(
		p,
		"resetLastModified",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetLatestUpdates() {
	_jsii_.InvokeVoid(
		p,
		"resetLatestUpdates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetLibrary() {
	_jsii_.InvokeVoid(
		p,
		"resetLibrary",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetNotification() {
	_jsii_.InvokeVoid(
		p,
		"resetNotification",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetPhoton() {
	_jsii_.InvokeVoid(
		p,
		"resetPhoton",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetRestartWindow() {
	_jsii_.InvokeVoid(
		p,
		"resetRestartWindow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetRootPath() {
	_jsii_.InvokeVoid(
		p,
		"resetRootPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetRunAs() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetRunAsUserName() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsUserName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetSchema() {
	_jsii_.InvokeVoid(
		p,
		"resetSchema",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetServerless() {
	_jsii_.InvokeVoid(
		p,
		"resetServerless",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetState() {
	_jsii_.InvokeVoid(
		p,
		"resetState",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetStorage() {
	_jsii_.InvokeVoid(
		p,
		"resetStorage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetTarget() {
	_jsii_.InvokeVoid(
		p,
		"resetTarget",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetTrigger() {
	_jsii_.InvokeVoid(
		p,
		"resetTrigger",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) ResetUrl() {
	_jsii_.InvokeVoid(
		p,
		"resetUrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

