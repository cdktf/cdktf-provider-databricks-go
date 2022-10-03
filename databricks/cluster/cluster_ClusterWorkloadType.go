package cluster


type ClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#clients Cluster#clients}
	Clients *ClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

