package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskNotebookTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#notebook_path DataDatabricksJob#notebook_path}.
	NotebookPath *string `field:"required" json:"notebookPath" yaml:"notebookPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#base_parameters DataDatabricksJob#base_parameters}.
	BaseParameters *map[string]*string `field:"optional" json:"baseParameters" yaml:"baseParameters"`
}

