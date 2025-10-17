// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfigNotificationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/data_quality_monitor#on_failure DataQualityMonitor#on_failure}.
	OnFailure *DataQualityMonitorDataProfilingConfigNotificationSettingsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
}

