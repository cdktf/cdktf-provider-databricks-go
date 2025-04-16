// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitor


type QualityMonitorNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/quality_monitor#on_failure QualityMonitor#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_new_classification_tag_detected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.73.0/docs/resources/quality_monitor#on_new_classification_tag_detected QualityMonitor#on_new_classification_tag_detected}
	OnNewClassificationTagDetected interface{} `field:"optional" json:"onNewClassificationTagDetected" yaml:"onNewClassificationTagDetected"`
}

