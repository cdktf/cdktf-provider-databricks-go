// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PipelineClusterGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#google_service_account Pipeline#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
}

