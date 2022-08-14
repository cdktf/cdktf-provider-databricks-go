// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MwsNetworksErrorMessages struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_networks#error_message MwsNetworks#error_message}.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_networks#error_type MwsNetworks#error_type}.
	ErrorType *string `field:"optional" json:"errorType" yaml:"errorType"`
}

