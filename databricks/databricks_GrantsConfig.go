// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GrantsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#grant Grants#grant}
	Grant interface{} `field:"required" json:"grant" yaml:"grant"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#catalog Grants#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#external_location Grants#external_location}.
	ExternalLocation *string `field:"optional" json:"externalLocation" yaml:"externalLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#function Grants#function}.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#id Grants#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#metastore Grants#metastore}.
	Metastore *string `field:"optional" json:"metastore" yaml:"metastore"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#schema Grants#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#storage_credential Grants#storage_credential}.
	StorageCredential *string `field:"optional" json:"storageCredential" yaml:"storageCredential"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#table Grants#table}.
	Table *string `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#view Grants#view}.
	View *string `field:"optional" json:"view" yaml:"view"`
}

