// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#availability DataDatabricksClusterPluginframework#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#boot_disk_size DataDatabricksClusterPluginframework#boot_disk_size}.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#google_service_account DataDatabricksClusterPluginframework#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#local_ssd_count DataDatabricksClusterPluginframework#local_ssd_count}.
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#use_preemptible_executors DataDatabricksClusterPluginframework#use_preemptible_executors}.
	UsePreemptibleExecutors interface{} `field:"optional" json:"usePreemptibleExecutors" yaml:"usePreemptibleExecutors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.66.0/docs/data-sources/cluster_pluginframework#zone_id DataDatabricksClusterPluginframework#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

