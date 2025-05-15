// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolGcpAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/instance_pool#gcp_availability InstancePool#gcp_availability}.
	GcpAvailability *string `field:"optional" json:"gcpAvailability" yaml:"gcpAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/instance_pool#local_ssd_count InstancePool#local_ssd_count}.
	LocalSsdCount *float64 `field:"optional" json:"localSsdCount" yaml:"localSsdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.79.1/docs/resources/instance_pool#zone_id InstancePool#zone_id}.
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

