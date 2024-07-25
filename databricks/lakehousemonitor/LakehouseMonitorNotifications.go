// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakehousemonitor


type LakehouseMonitorNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/resources/lakehouse_monitor#on_failure LakehouseMonitor#on_failure}
	OnFailure *LakehouseMonitorNotificationsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_new_classification_tag_detected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/resources/lakehouse_monitor#on_new_classification_tag_detected LakehouseMonitor#on_new_classification_tag_detected}
	OnNewClassificationTagDetected *LakehouseMonitorNotificationsOnNewClassificationTagDetected `field:"optional" json:"onNewClassificationTagDetected" yaml:"onNewClassificationTagDetected"`
}

