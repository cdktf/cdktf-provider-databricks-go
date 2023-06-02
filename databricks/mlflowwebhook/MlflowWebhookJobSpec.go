package mlflowwebhook


type MlflowWebhookJobSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/mlflow_webhook#access_token MlflowWebhook#access_token}.
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/mlflow_webhook#job_id MlflowWebhook#job_id}.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/mlflow_webhook#workspace_url MlflowWebhook#workspace_url}.
	WorkspaceUrl *string `field:"optional" json:"workspaceUrl" yaml:"workspaceUrl"`
}

