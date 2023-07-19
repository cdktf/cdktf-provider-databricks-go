package job


type JobTaskNewClusterClusterLogConfDbfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

