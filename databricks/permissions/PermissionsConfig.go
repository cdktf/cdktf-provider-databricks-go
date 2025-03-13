// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package permissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PermissionsConfig struct {
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
	// access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#access_control Permissions#access_control}
	AccessControl interface{} `field:"required" json:"accessControl" yaml:"accessControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#app_name Permissions#app_name}.
	AppName *string `field:"optional" json:"appName" yaml:"appName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#authorization Permissions#authorization}.
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#cluster_id Permissions#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#cluster_policy_id Permissions#cluster_policy_id}.
	ClusterPolicyId *string `field:"optional" json:"clusterPolicyId" yaml:"clusterPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#dashboard_id Permissions#dashboard_id}.
	DashboardId *string `field:"optional" json:"dashboardId" yaml:"dashboardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#directory_id Permissions#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#directory_path Permissions#directory_path}.
	DirectoryPath *string `field:"optional" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#experiment_id Permissions#experiment_id}.
	ExperimentId *string `field:"optional" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#id Permissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#instance_pool_id Permissions#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#job_id Permissions#job_id}.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#notebook_id Permissions#notebook_id}.
	NotebookId *string `field:"optional" json:"notebookId" yaml:"notebookId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#notebook_path Permissions#notebook_path}.
	NotebookPath *string `field:"optional" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#object_type Permissions#object_type}.
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#pipeline_id Permissions#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#registered_model_id Permissions#registered_model_id}.
	RegisteredModelId *string `field:"optional" json:"registeredModelId" yaml:"registeredModelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#repo_id Permissions#repo_id}.
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#repo_path Permissions#repo_path}.
	RepoPath *string `field:"optional" json:"repoPath" yaml:"repoPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#serving_endpoint_id Permissions#serving_endpoint_id}.
	ServingEndpointId *string `field:"optional" json:"servingEndpointId" yaml:"servingEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#sql_alert_id Permissions#sql_alert_id}.
	SqlAlertId *string `field:"optional" json:"sqlAlertId" yaml:"sqlAlertId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#sql_dashboard_id Permissions#sql_dashboard_id}.
	SqlDashboardId *string `field:"optional" json:"sqlDashboardId" yaml:"sqlDashboardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#sql_endpoint_id Permissions#sql_endpoint_id}.
	SqlEndpointId *string `field:"optional" json:"sqlEndpointId" yaml:"sqlEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#sql_query_id Permissions#sql_query_id}.
	SqlQueryId *string `field:"optional" json:"sqlQueryId" yaml:"sqlQueryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#vector_search_endpoint_id Permissions#vector_search_endpoint_id}.
	VectorSearchEndpointId *string `field:"optional" json:"vectorSearchEndpointId" yaml:"vectorSearchEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#workspace_file_id Permissions#workspace_file_id}.
	WorkspaceFileId *string `field:"optional" json:"workspaceFileId" yaml:"workspaceFileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.70.0/docs/resources/permissions#workspace_file_path Permissions#workspace_file_path}.
	WorkspaceFilePath *string `field:"optional" json:"workspaceFilePath" yaml:"workspaceFilePath"`
}

