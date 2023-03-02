package sqlquery


type SqlQueryScheduleDaily struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#interval_days SqlQuery#interval_days}.
	IntervalDays *float64 `field:"required" json:"intervalDays" yaml:"intervalDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#time_of_day SqlQuery#time_of_day}.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#until_date SqlQuery#until_date}.
	UntilDate *string `field:"optional" json:"untilDate" yaml:"untilDate"`
}

