package recipient


type RecipientIpAccessList struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/recipient#allowed_ip_addresses Recipient#allowed_ip_addresses}.
	AllowedIpAddresses *[]*string `field:"required" json:"allowedIpAddresses" yaml:"allowedIpAddresses"`
}

