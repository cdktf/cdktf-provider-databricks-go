// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsTaskLibraryPypi struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#package DataDatabricksJob#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#repo DataDatabricksJob#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

