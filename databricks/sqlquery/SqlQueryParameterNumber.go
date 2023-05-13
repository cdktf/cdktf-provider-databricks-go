package sqlquery


type SqlQueryParameterNumber struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.0/docs/resources/sql_query#value SqlQuery#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

