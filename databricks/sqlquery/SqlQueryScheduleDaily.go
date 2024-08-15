// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlquery


type SqlQueryScheduleDaily struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/sql_query#interval_days SqlQuery#interval_days}.
	IntervalDays *float64 `field:"required" json:"intervalDays" yaml:"intervalDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/sql_query#time_of_day SqlQuery#time_of_day}.
	TimeOfDay *string `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/sql_query#until_date SqlQuery#until_date}.
	UntilDate *string `field:"optional" json:"untilDate" yaml:"untilDate"`
}

