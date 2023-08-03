package job


type JobTaskNewClusterInitScriptsAbfss struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

