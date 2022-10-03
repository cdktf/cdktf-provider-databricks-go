package job


type JobJobClusterNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#clients Job#clients}
	Clients *JobJobClusterNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

