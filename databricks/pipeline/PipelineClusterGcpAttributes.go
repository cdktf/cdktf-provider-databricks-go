package pipeline


type PipelineClusterGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#availability Pipeline#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#google_service_account Pipeline#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#zone_id Pipeline#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

