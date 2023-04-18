package library


type LibraryCran struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/library#package Library#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/library#repo Library#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

