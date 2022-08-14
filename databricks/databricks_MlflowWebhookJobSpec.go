// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MlflowWebhookJobSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#access_token MlflowWebhook#access_token}.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#job_id MlflowWebhook#job_id}.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#workspace_url MlflowWebhook#workspace_url}.
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

