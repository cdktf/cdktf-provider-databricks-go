package sqlquery


type SqlQuerySchedule struct {
	// continuous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/sql_query#continuous SqlQuery#continuous}
	Continuous *SqlQueryScheduleContinuous `field:"optional" json:"continuous" yaml:"continuous"`
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/sql_query#daily SqlQuery#daily}
	Daily *SqlQueryScheduleDaily `field:"optional" json:"daily" yaml:"daily"`
	// weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/sql_query#weekly SqlQuery#weekly}
	Weekly *SqlQueryScheduleWeekly `field:"optional" json:"weekly" yaml:"weekly"`
}

