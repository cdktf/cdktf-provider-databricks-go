// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#created_time DataDatabricksJob#created_time}.
	CreatedTime *float64 `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#creator_user_name DataDatabricksJob#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#job_id DataDatabricksJob#job_id}.
	JobId *float64 `field:"optional" json:"jobId" yaml:"jobId"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#settings DataDatabricksJob#settings}
	Settings *DataDatabricksJobJobSettingsSettings `field:"optional" json:"settings" yaml:"settings"`
}

