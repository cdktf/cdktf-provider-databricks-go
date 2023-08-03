package pipeline


type PipelineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/pipeline#default Pipeline#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

