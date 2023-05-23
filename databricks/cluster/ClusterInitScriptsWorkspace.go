package cluster


type ClusterInitScriptsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/cluster#destination Cluster#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

