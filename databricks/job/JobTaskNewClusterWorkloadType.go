package job


type JobTaskNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.19.0/docs/resources/job#clients Job#clients}
	Clients *JobTaskNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

