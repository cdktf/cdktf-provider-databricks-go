// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type MountGs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#bucket_name Mount#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/mount#service_account Mount#service_account}.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

