package cluster


type ClusterClusterLogConfDbfs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#destination Cluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}
