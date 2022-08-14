// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlEndpointTags struct {
	// custom_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_endpoint#custom_tags SqlEndpoint#custom_tags}
	CustomTags interface{} `field:"required" json:"customTags" yaml:"customTags"`
}

