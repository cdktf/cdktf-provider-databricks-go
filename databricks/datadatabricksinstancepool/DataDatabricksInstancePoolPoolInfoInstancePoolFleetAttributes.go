package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributes struct {
	// launch_template_override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#launch_template_override DataDatabricksInstancePool#launch_template_override}
	LaunchTemplateOverride interface{} `field:"required" json:"launchTemplateOverride" yaml:"launchTemplateOverride"`
	// fleet_on_demand_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#fleet_on_demand_option DataDatabricksInstancePool#fleet_on_demand_option}
	FleetOnDemandOption *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetOnDemandOption `field:"optional" json:"fleetOnDemandOption" yaml:"fleetOnDemandOption"`
	// fleet_spot_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/instance_pool#fleet_spot_option DataDatabricksInstancePool#fleet_spot_option}
	FleetSpotOption *DataDatabricksInstancePoolPoolInfoInstancePoolFleetAttributesFleetSpotOption `field:"optional" json:"fleetSpotOption" yaml:"fleetSpotOption"`
}

