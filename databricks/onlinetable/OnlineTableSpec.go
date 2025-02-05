// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinetable


type OnlineTableSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#perform_full_copy OnlineTable#perform_full_copy}.
	PerformFullCopy interface{} `field:"optional" json:"performFullCopy" yaml:"performFullCopy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#primary_key_columns OnlineTable#primary_key_columns}.
	PrimaryKeyColumns *[]*string `field:"optional" json:"primaryKeyColumns" yaml:"primaryKeyColumns"`
	// run_continuously block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#run_continuously OnlineTable#run_continuously}
	RunContinuously *OnlineTableSpecRunContinuously `field:"optional" json:"runContinuously" yaml:"runContinuously"`
	// run_triggered block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#run_triggered OnlineTable#run_triggered}
	RunTriggered *OnlineTableSpecRunTriggered `field:"optional" json:"runTriggered" yaml:"runTriggered"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#source_table_full_name OnlineTable#source_table_full_name}.
	SourceTableFullName *string `field:"optional" json:"sourceTableFullName" yaml:"sourceTableFullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.65.0/docs/resources/online_table#timeseries_key OnlineTable#timeseries_key}.
	TimeseriesKey *string `field:"optional" json:"timeseriesKey" yaml:"timeseriesKey"`
}

