// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsLibraryMaven struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#coordinates DataDatabricksJob#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#exclusions DataDatabricksJob#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#repo DataDatabricksJob#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

