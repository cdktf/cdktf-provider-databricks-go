// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorCustomMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#definition LakehouseMonitor#definition}.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#input_columns LakehouseMonitor#input_columns}.
	InputColumns *[]*string `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#name LakehouseMonitor#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#output_data_type LakehouseMonitor#output_data_type}.
	OutputDataType *string `field:"optional" json:"outputDataType" yaml:"outputDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#type LakehouseMonitor#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
