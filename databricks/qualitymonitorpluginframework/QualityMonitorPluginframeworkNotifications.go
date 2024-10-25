// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkNotifications struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#on_failure QualityMonitorPluginframework#on_failure}
	OnFailure interface{} `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_new_classification_tag_detected block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.55.0/docs/resources/quality_monitor_pluginframework#on_new_classification_tag_detected QualityMonitorPluginframework#on_new_classification_tag_detected}
	OnNewClassificationTagDetected interface{} `field:"optional" json:"onNewClassificationTagDetected" yaml:"onNewClassificationTagDetected"`
}

