// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecipientConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#authentication_type Recipient#authentication_type}.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#name Recipient#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#comment Recipient#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#data_recipient_global_metastore_id Recipient#data_recipient_global_metastore_id}.
	DataRecipientGlobalMetastoreId *string `field:"optional" json:"dataRecipientGlobalMetastoreId" yaml:"dataRecipientGlobalMetastoreId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#id Recipient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_access_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#ip_access_list Recipient#ip_access_list}
	IpAccessList *RecipientIpAccessList `field:"optional" json:"ipAccessList" yaml:"ipAccessList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#sharing_code Recipient#sharing_code}.
	SharingCode *string `field:"optional" json:"sharingCode" yaml:"sharingCode"`
	// tokens block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/recipient#tokens Recipient#tokens}
	Tokens interface{} `field:"optional" json:"tokens" yaml:"tokens"`
}

