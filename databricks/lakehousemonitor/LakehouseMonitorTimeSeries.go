// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorTimeSeries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/lakehouse_monitor#granularities LakehouseMonitor#granularities}.
	Granularities *[]*string `field:"required" json:"granularities" yaml:"granularities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.57.0/docs/resources/lakehouse_monitor#timestamp_col LakehouseMonitor#timestamp_col}.
	TimestampCol *string `field:"required" json:"timestampCol" yaml:"timestampCol"`
}

