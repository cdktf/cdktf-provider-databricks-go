package datadatabrickscluster


type DataDatabricksClusterClusterInfoAutoscale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#max_workers DataDatabricksCluster#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#min_workers DataDatabricksCluster#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
}

