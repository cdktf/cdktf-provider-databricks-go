// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlQueryScheduleWeekly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#day_of_week SqlQuery#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#interval_weeks SqlQuery#interval_weeks}.
	IntervalWeeks *float64 `field:"required" json:"intervalWeeks" yaml:"intervalWeeks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#time_of_day SqlQuery#time_of_day}.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#until_date SqlQuery#until_date}.
	UntilDate *string `field:"optional" json:"untilDate" yaml:"untilDate"`
}
