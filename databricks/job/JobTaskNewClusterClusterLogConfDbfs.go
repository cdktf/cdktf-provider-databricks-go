package job


type JobTaskNewClusterClusterLogConfDbfs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#destination Job#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

