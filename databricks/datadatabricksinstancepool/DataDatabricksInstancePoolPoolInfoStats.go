// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoStats struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/instance_pool#idle_count DataDatabricksInstancePool#idle_count}.
	IdleCount *float64 `field:"optional" json:"idleCount" yaml:"idleCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/instance_pool#pending_idle_count DataDatabricksInstancePool#pending_idle_count}.
	PendingIdleCount *float64 `field:"optional" json:"pendingIdleCount" yaml:"pendingIdleCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/instance_pool#pending_used_count DataDatabricksInstancePool#pending_used_count}.
	PendingUsedCount *float64 `field:"optional" json:"pendingUsedCount" yaml:"pendingUsedCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/instance_pool#used_count DataDatabricksInstancePool#used_count}.
	UsedCount *float64 `field:"optional" json:"usedCount" yaml:"usedCount"`
}

