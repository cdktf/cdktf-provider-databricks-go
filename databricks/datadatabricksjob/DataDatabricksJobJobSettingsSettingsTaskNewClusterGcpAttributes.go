// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskNewClusterGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#availability DataDatabricksJob#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#boot_disk_size DataDatabricksJob#boot_disk_size}.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#google_service_account DataDatabricksJob#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#local_ssd_count DataDatabricksJob#local_ssd_count}.
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#use_preemptible_executors DataDatabricksJob#use_preemptible_executors}.
	UsePreemptibleExecutors interface{} `field:"optional" json:"usePreemptibleExecutors" yaml:"usePreemptibleExecutors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.0/docs/data-sources/job#zone_id DataDatabricksJob#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

