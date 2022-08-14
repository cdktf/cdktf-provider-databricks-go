// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobTaskDbtTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#commands Job#commands}.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#project_directory Job#project_directory}.
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#schema Job#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

