// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsNewClusterGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#availability DataDatabricksJob#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#boot_disk_size DataDatabricksJob#boot_disk_size}.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#google_service_account DataDatabricksJob#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#use_preemptible_executors DataDatabricksJob#use_preemptible_executors}.
	UsePreemptibleExecutors interface{} `field:"optional" json:"usePreemptibleExecutors" yaml:"usePreemptibleExecutors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#zone_id DataDatabricksJob#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

