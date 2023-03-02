package job


type JobLibraryCran struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#package Job#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#repo Job#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

