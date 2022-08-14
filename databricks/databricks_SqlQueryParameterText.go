// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlQueryParameterText struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#value SqlQuery#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

