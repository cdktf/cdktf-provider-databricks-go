package mlflowwebhook


type MlflowWebhookHttpUrlSpec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#url MlflowWebhook#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#authorization MlflowWebhook#authorization}.
	Authorization *string `field:"optional" json:"authorization" yaml:"authorization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#enable_ssl_verification MlflowWebhook#enable_ssl_verification}.
	EnableSslVerification interface{} `field:"optional" json:"enableSslVerification" yaml:"enableSslVerification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mlflow_webhook#secret MlflowWebhook#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}
