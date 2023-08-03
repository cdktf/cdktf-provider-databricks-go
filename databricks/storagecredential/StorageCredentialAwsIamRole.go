package storagecredential


type StorageCredentialAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/storage_credential#role_arn StorageCredential#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

