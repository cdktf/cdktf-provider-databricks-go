package job


type JobNewClusterInitScriptsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#destination Job#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

