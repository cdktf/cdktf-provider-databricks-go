package permissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PermissionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#access_control Permissions#access_control}
	AccessControl interface{} `field:"required" json:"accessControl" yaml:"accessControl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#authorization Permissions#authorization}.
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#cluster_id Permissions#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#cluster_policy_id Permissions#cluster_policy_id}.
	ClusterPolicyId *string `field:"optional" json:"clusterPolicyId" yaml:"clusterPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#directory_id Permissions#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#directory_path Permissions#directory_path}.
	DirectoryPath *string `field:"optional" json:"directoryPath" yaml:"directoryPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#experiment_id Permissions#experiment_id}.
	ExperimentId *string `field:"optional" json:"experimentId" yaml:"experimentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#id Permissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#instance_pool_id Permissions#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#job_id Permissions#job_id}.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#notebook_id Permissions#notebook_id}.
	NotebookId *string `field:"optional" json:"notebookId" yaml:"notebookId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#notebook_path Permissions#notebook_path}.
	NotebookPath *string `field:"optional" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#object_type Permissions#object_type}.
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#pipeline_id Permissions#pipeline_id}.
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#registered_model_id Permissions#registered_model_id}.
	RegisteredModelId *string `field:"optional" json:"registeredModelId" yaml:"registeredModelId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#repo_id Permissions#repo_id}.
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#repo_path Permissions#repo_path}.
	RepoPath *string `field:"optional" json:"repoPath" yaml:"repoPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#sql_alert_id Permissions#sql_alert_id}.
	SqlAlertId *string `field:"optional" json:"sqlAlertId" yaml:"sqlAlertId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#sql_dashboard_id Permissions#sql_dashboard_id}.
	SqlDashboardId *string `field:"optional" json:"sqlDashboardId" yaml:"sqlDashboardId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#sql_endpoint_id Permissions#sql_endpoint_id}.
	SqlEndpointId *string `field:"optional" json:"sqlEndpointId" yaml:"sqlEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/permissions#sql_query_id Permissions#sql_query_id}.
	SqlQueryId *string `field:"optional" json:"sqlQueryId" yaml:"sqlQueryId"`
}

