package cluster


type ClusterLibraryCran struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#package Cluster#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#repo Cluster#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}
