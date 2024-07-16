// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.3/docs/resources/quality_monitor#create QualityMonitor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

