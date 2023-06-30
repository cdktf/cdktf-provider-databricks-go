package job


type JobTaskSqlTaskFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/job#path Job#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

