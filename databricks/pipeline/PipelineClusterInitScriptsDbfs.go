package pipeline


type PipelineClusterInitScriptsDbfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/pipeline#destination Pipeline#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

