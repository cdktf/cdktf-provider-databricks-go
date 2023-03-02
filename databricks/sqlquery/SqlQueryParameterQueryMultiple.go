package sqlquery


type SqlQueryParameterQueryMultiple struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#prefix SqlQuery#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#separator SqlQuery#separator}.
	Separator *string `field:"required" json:"separator" yaml:"separator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#suffix SqlQuery#suffix}.
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
}

