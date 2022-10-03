package job


type JobLibraryMaven struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#coordinates Job#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#exclusions Job#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#repo Job#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

