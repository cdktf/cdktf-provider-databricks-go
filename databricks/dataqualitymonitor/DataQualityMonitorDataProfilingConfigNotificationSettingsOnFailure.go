// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataqualitymonitor


type DataQualityMonitorDataProfilingConfigNotificationSettingsOnFailure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/data_quality_monitor#email_addresses DataQualityMonitor#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
}

