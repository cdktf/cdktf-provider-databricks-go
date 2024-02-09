// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#availability DataDatabricksCluster#availability}.
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#boot_disk_size DataDatabricksCluster#boot_disk_size}.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#google_service_account DataDatabricksCluster#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#local_ssd_count DataDatabricksCluster#local_ssd_count}.
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#use_preemptible_executors DataDatabricksCluster#use_preemptible_executors}.
	UsePreemptibleExecutors interface{} `field:"optional" json:"usePreemptibleExecutors" yaml:"usePreemptibleExecutors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/data-sources/cluster#zone_id DataDatabricksCluster#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

