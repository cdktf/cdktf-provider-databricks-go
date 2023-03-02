package storagecredential


type StorageCredentialAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/storage_credential#role_arn StorageCredential#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

