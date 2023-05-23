package mount


type MountS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/mount#bucket_name Mount#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/resources/mount#instance_profile Mount#instance_profile}.
	InstanceProfile *string `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
}

