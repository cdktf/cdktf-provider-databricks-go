package mwsnetworks


type MwsNetworksVpcEndpoints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_networks#dataplane_relay MwsNetworks#dataplane_relay}.
	DataplaneRelay *[]*string `field:"required" json:"dataplaneRelay" yaml:"dataplaneRelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_networks#rest_api MwsNetworks#rest_api}.
	RestApi *[]*string `field:"required" json:"restApi" yaml:"restApi"`
}

