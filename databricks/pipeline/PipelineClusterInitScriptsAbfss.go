package pipeline


type PipelineClusterInitScriptsAbfss struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/pipeline#destination Pipeline#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

