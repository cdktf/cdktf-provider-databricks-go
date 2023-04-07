package job


type JobTaskSqlTaskFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#path Job#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

