package sqlquery


type SqlQueryParameterNumber struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#value SqlQuery#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

