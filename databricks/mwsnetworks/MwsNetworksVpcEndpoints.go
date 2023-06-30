package mwsnetworks


type MwsNetworksVpcEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/mws_networks#dataplane_relay MwsNetworks#dataplane_relay}.
	DataplaneRelay *[]*string `field:"required" json:"dataplaneRelay" yaml:"dataplaneRelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/mws_networks#rest_api MwsNetworks#rest_api}.
	RestApi *[]*string `field:"required" json:"restApi" yaml:"restApi"`
}

