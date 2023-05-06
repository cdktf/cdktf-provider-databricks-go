package job


type JobNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.15.0/docs/resources/job#clients Job#clients}
	Clients *JobNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

