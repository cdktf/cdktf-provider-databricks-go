package job


type JobNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#clients Job#clients}
	Clients *JobNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

