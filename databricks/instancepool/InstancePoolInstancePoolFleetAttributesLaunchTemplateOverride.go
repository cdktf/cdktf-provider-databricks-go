// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolInstancePoolFleetAttributesLaunchTemplateOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/resources/instance_pool#availability_zone InstancePool#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.29.0/docs/resources/instance_pool#instance_type InstancePool#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

