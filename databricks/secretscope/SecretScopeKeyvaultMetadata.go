package secretscope


type SecretScopeKeyvaultMetadata struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/secret_scope#dns_name SecretScope#dns_name}.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/secret_scope#resource_id SecretScope#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

