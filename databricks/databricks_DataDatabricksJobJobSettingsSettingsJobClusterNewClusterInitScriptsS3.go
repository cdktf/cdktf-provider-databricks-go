// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterInitScriptsS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#destination DataDatabricksJob#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#canned_acl DataDatabricksJob#canned_acl}.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#enable_encryption DataDatabricksJob#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#encryption_type DataDatabricksJob#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#endpoint DataDatabricksJob#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#kms_key DataDatabricksJob#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#region DataDatabricksJob#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

