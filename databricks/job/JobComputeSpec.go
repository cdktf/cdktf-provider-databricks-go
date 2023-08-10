package job


type JobComputeSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/job#kind Job#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

