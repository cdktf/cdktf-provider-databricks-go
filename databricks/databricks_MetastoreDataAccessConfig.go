// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetastoreDataAccessConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#metastore_id MetastoreDataAccess#metastore_id}.
	MetastoreId *string `field:"required" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#name MetastoreDataAccess#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// aws_iam_role block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#aws_iam_role MetastoreDataAccess#aws_iam_role}
	AwsIamRole *MetastoreDataAccessAwsIamRole `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// azure_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#azure_managed_identity MetastoreDataAccess#azure_managed_identity}
	AzureManagedIdentity *MetastoreDataAccessAzureManagedIdentity `field:"optional" json:"azureManagedIdentity" yaml:"azureManagedIdentity"`
	// azure_service_principal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#azure_service_principal MetastoreDataAccess#azure_service_principal}
	AzureServicePrincipal *MetastoreDataAccessAzureServicePrincipal `field:"optional" json:"azureServicePrincipal" yaml:"azureServicePrincipal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#configuration_type MetastoreDataAccess#configuration_type}.
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#id MetastoreDataAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/metastore_data_access#is_default MetastoreDataAccess#is_default}.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
}

