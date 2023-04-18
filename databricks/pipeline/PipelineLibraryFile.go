package pipeline


type PipelineLibraryFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/pipeline#path Pipeline#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

