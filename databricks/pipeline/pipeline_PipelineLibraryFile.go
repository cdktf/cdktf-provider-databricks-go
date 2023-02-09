package pipeline


type PipelineLibraryFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#path Pipeline#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

