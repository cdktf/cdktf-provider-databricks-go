// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolInstancePoolFleetAttributes struct {
	// launch_template_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/instance_pool#launch_template_override InstancePool#launch_template_override}
	LaunchTemplateOverride interface{} `field:"required" json:"launchTemplateOverride" yaml:"launchTemplateOverride"`
	// fleet_on_demand_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/instance_pool#fleet_on_demand_option InstancePool#fleet_on_demand_option}
	FleetOnDemandOption *InstancePoolInstancePoolFleetAttributesFleetOnDemandOption `field:"optional" json:"fleetOnDemandOption" yaml:"fleetOnDemandOption"`
	// fleet_spot_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/instance_pool#fleet_spot_option InstancePool#fleet_spot_option}
	FleetSpotOption *InstancePoolInstancePoolFleetAttributesFleetSpotOption `field:"optional" json:"fleetSpotOption" yaml:"fleetSpotOption"`
}

