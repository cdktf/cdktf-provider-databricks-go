// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksSqlWarehouseOdbcParams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/sql_warehouse#path DataDatabricksSqlWarehouse#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/sql_warehouse#port DataDatabricksSqlWarehouse#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/sql_warehouse#protocol DataDatabricksSqlWarehouse#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/sql_warehouse#host DataDatabricksSqlWarehouse#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/sql_warehouse#hostname DataDatabricksSqlWarehouse#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}
