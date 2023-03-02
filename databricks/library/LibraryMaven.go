package library


type LibraryMaven struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/library#coordinates Library#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/library#exclusions Library#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/library#repo Library#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

