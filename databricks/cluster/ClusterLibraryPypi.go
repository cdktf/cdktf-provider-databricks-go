package cluster


type ClusterLibraryPypi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/cluster#package Cluster#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/cluster#repo Cluster#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

