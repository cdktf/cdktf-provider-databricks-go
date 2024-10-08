// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package qualitymonitorpluginframework


type QualityMonitorPluginframeworkNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/quality_monitor_pluginframework#on_failure QualityMonitorPluginframework#on_failure}.
	OnFailure *QualityMonitorPluginframeworkNotificationsOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.53.0/docs/resources/quality_monitor_pluginframework#on_new_classification_tag_detected QualityMonitorPluginframework#on_new_classification_tag_detected}.
	OnNewClassificationTagDetected *QualityMonitorPluginframeworkNotificationsOnNewClassificationTagDetected `field:"optional" json:"onNewClassificationTagDetected" yaml:"onNewClassificationTagDetected"`
}

