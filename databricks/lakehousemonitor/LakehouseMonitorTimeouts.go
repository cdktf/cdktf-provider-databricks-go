// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.45.0/docs/resources/lakehouse_monitor#create LakehouseMonitor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

