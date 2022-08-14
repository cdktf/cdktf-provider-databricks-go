// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type ClusterInitScriptsGcs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/cluster#destination Cluster#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

