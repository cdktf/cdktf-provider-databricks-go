package cluster


type ClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/cluster#clients Cluster#clients}
	Clients *ClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

