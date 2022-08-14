// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobGitSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#url Job#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#branch Job#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#commit Job#commit}.
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#provider Job#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#tag Job#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

