package cluster


type ClusterInitScriptsAbfss struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/cluster#destination Cluster#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

