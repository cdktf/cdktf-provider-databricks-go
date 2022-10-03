package job


type JobTaskNotebookTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#notebook_path Job#notebook_path}.
	NotebookPath *string `field:"required" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#base_parameters Job#base_parameters}.
	BaseParameters *map[string]*string `field:"optional" json:"baseParameters" yaml:"baseParameters"`
}

