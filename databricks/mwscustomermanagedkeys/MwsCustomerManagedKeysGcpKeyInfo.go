package mwscustomermanagedkeys


type MwsCustomerManagedKeysGcpKeyInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mws_customer_managed_keys#kms_key_id MwsCustomerManagedKeys#kms_key_id}.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

