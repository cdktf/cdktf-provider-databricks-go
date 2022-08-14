// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MountS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#bucket_name Mount#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#instance_profile Mount#instance_profile}.
	InstanceProfile *string `field:"optional" json:"instanceProfile" yaml:"instanceProfile"`
}

