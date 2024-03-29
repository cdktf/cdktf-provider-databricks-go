// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.39.0/docs/resources/lakehouse_monitor#on_failure LakehouseMonitor#on_failure}
	OnFailure *LakehouseMonitorNotificationsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
}

