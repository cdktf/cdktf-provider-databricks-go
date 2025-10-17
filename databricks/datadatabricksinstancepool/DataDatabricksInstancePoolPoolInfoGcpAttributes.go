// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/instance_pool#gcp_availability DataDatabricksInstancePool#gcp_availability}.
	GcpAvailability *string `field:"optional" json:"gcpAvailability" yaml:"gcpAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/instance_pool#local_ssd_count DataDatabricksInstancePool#local_ssd_count}.
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/instance_pool#zone_id DataDatabricksInstancePool#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

