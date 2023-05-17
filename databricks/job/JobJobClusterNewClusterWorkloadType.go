package job


type JobJobClusterNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/job#clients Job#clients}
	Clients *JobJobClusterNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

