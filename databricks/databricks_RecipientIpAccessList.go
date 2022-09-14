// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type RecipientIpAccessList struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#allowed_ip_addresses Recipient#allowed_ip_addresses}.
	AllowedIpAddresses *[]*string `field:"required" json:"allowedIpAddresses" yaml:"allowedIpAddresses"`
}

