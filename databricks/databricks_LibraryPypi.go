// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type LibraryPypi struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/library#package Library#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/library#repo Library#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

