package job


type JobNewClusterInitScriptsFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.15.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

