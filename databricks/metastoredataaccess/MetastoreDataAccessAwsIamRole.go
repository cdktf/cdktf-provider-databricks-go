package metastoredataaccess


type MetastoreDataAccessAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/resources/metastore_data_access#role_arn MetastoreDataAccess#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

