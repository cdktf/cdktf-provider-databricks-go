package pipeline


type PipelineLibraryNotebook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/pipeline#path Pipeline#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

