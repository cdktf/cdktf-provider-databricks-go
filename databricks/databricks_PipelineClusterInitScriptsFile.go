// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PipelineClusterInitScriptsFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#destination Pipeline#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

