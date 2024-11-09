// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesLaunchTemplateOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/instance_pool#availability_zone DataDatabricksInstancePool#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.58.0/docs/data-sources/instance_pool#instance_type DataDatabricksInstancePool#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

