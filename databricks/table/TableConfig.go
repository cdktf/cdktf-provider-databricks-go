package table

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TableConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#catalog_name Table#catalog_name}.
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// column block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#column Table#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#data_source_format Table#data_source_format}.
	DataSourceFormat *string `field:"required" json:"dataSourceFormat" yaml:"dataSourceFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#name Table#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#schema_name Table#schema_name}.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#table_type Table#table_type}.
	TableType *string `field:"required" json:"tableType" yaml:"tableType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#comment Table#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#id Table#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#owner Table#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#properties Table#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#storage_credential_name Table#storage_credential_name}.
	StorageCredentialName *string `field:"optional" json:"storageCredentialName" yaml:"storageCredentialName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#storage_location Table#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/table#view_definition Table#view_definition}.
	ViewDefinition *string `field:"optional" json:"viewDefinition" yaml:"viewDefinition"`
}

