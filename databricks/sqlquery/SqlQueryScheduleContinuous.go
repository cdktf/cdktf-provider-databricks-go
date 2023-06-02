package sqlquery


type SqlQueryScheduleContinuous struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/sql_query#interval_seconds SqlQuery#interval_seconds}.
	IntervalSeconds *float64 `field:"required" json:"intervalSeconds" yaml:"intervalSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/sql_query#until_date SqlQuery#until_date}.
	UntilDate *string `field:"optional" json:"untilDate" yaml:"untilDate"`
}

