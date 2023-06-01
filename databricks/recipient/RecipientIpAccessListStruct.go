package recipient


type RecipientIpAccessListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/recipient#allowed_ip_addresses Recipient#allowed_ip_addresses}.
	AllowedIpAddresses *[]*string `field:"required" json:"allowedIpAddresses" yaml:"allowedIpAddresses"`
}

