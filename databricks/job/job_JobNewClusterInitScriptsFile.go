package job


type JobNewClusterInitScriptsFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#destination Job#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}
