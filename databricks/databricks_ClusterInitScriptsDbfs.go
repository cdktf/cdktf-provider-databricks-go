// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type ClusterInitScriptsDbfs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#destination Cluster#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

