// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobTaskNewClusterInitScriptsS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#destination Job#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#canned_acl Job#canned_acl}.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#enable_encryption Job#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#encryption_type Job#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#endpoint Job#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#kms_key Job#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#region Job#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

