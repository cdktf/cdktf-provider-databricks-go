package mwscustomermanagedkeys


type MwsCustomerManagedKeysAwsKeyInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_customer_managed_keys#key_alias MwsCustomerManagedKeys#key_alias}.
	KeyAlias *string `field:"required" json:"keyAlias" yaml:"keyAlias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_customer_managed_keys#key_arn MwsCustomerManagedKeys#key_arn}.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_customer_managed_keys#key_region MwsCustomerManagedKeys#key_region}.
	KeyRegion *string `field:"optional" json:"keyRegion" yaml:"keyRegion"`
}
