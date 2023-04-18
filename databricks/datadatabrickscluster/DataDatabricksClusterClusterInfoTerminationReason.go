package datadatabrickscluster


type DataDatabricksClusterClusterInfoTerminationReason struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/cluster#code DataDatabricksCluster#code}.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/cluster#parameters DataDatabricksCluster#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/cluster#type DataDatabricksCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

