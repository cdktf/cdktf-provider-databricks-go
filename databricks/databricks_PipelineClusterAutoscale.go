// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PipelineClusterAutoscale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#max_workers Pipeline#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#min_workers Pipeline#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#mode Pipeline#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

