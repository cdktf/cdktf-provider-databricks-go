// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type PipelineLibraryNotebook struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/pipeline#path Pipeline#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

