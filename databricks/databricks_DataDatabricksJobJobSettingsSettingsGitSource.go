// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsGitSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#branch DataDatabricksJob#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#commit DataDatabricksJob#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#provider DataDatabricksJob#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#tag DataDatabricksJob#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

