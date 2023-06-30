package sqlquery


type SqlQueryParameterDatetimeRangeRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/sql_query#end SqlQuery#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/sql_query#start SqlQuery#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

