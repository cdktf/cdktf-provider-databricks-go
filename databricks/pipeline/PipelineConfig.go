// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#allow_duplicate_names Pipeline#allow_duplicate_names}.
	AllowDuplicateNames interface{} `field:"optional" json:"allowDuplicateNames" yaml:"allowDuplicateNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#budget_policy_id Pipeline#budget_policy_id}.
	BudgetPolicyId *string `field:"optional" json:"budgetPolicyId" yaml:"budgetPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#catalog Pipeline#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#cause Pipeline#cause}.
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#channel Pipeline#channel}.
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#cluster Pipeline#cluster}
	Cluster interface{} `field:"optional" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#cluster_id Pipeline#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#configuration Pipeline#configuration}.
	Configuration *map[string]*string `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#continuous Pipeline#continuous}.
	Continuous interface{} `field:"optional" json:"continuous" yaml:"continuous"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#creator_user_name Pipeline#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#deployment Pipeline#deployment}
	Deployment *PipelineDeployment `field:"optional" json:"deployment" yaml:"deployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#development Pipeline#development}.
	Development interface{} `field:"optional" json:"development" yaml:"development"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#edition Pipeline#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#environment Pipeline#environment}
	Environment *PipelineEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// event_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#event_log Pipeline#event_log}
	EventLog *PipelineEventLog `field:"optional" json:"eventLog" yaml:"eventLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#expected_last_modified Pipeline#expected_last_modified}.
	ExpectedLastModified *float64 `field:"optional" json:"expectedLastModified" yaml:"expectedLastModified"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#filters Pipeline#filters}
	Filters *PipelineFilters `field:"optional" json:"filters" yaml:"filters"`
	// gateway_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#gateway_definition Pipeline#gateway_definition}
	GatewayDefinition *PipelineGatewayDefinition `field:"optional" json:"gatewayDefinition" yaml:"gatewayDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#health Pipeline#health}.
	Health *string `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#id Pipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ingestion_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#ingestion_definition Pipeline#ingestion_definition}
	IngestionDefinition *PipelineIngestionDefinition `field:"optional" json:"ingestionDefinition" yaml:"ingestionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#last_modified Pipeline#last_modified}.
	LastModified *float64 `field:"optional" json:"lastModified" yaml:"lastModified"`
	// latest_updates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#latest_updates Pipeline#latest_updates}
	LatestUpdates interface{} `field:"optional" json:"latestUpdates" yaml:"latestUpdates"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#library Pipeline#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#name Pipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#notification Pipeline#notification}
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#photon Pipeline#photon}.
	Photon interface{} `field:"optional" json:"photon" yaml:"photon"`
	// restart_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#restart_window Pipeline#restart_window}
	RestartWindow *PipelineRestartWindow `field:"optional" json:"restartWindow" yaml:"restartWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#root_path Pipeline#root_path}.
	RootPath *string `field:"optional" json:"rootPath" yaml:"rootPath"`
	// run_as block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#run_as Pipeline#run_as}
	RunAs *PipelineRunAs `field:"optional" json:"runAs" yaml:"runAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#run_as_user_name Pipeline#run_as_user_name}.
	RunAsUserName *string `field:"optional" json:"runAsUserName" yaml:"runAsUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#schema Pipeline#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#serverless Pipeline#serverless}.
	Serverless interface{} `field:"optional" json:"serverless" yaml:"serverless"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#state Pipeline#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#storage Pipeline#storage}.
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#tags Pipeline#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#target Pipeline#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#timeouts Pipeline#timeouts}
	Timeouts *PipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#trigger Pipeline#trigger}
	Trigger *PipelineTrigger `field:"optional" json:"trigger" yaml:"trigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/resources/pipeline#url Pipeline#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

