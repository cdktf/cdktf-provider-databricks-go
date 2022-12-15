package sqlquery


type SqlQueryParameterDateRangeRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#end SqlQuery#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#start SqlQuery#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

