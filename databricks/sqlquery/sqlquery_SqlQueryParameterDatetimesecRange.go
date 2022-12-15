package sqlquery


type SqlQueryParameterDatetimesecRange struct {
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#range SqlQuery#range}
	Range *SqlQueryParameterDatetimesecRangeRange `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#value SqlQuery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

