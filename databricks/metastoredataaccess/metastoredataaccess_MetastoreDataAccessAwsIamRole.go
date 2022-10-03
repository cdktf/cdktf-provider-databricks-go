package metastoredataaccess


type MetastoreDataAccessAwsIamRole struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#role_arn MetastoreDataAccess#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

