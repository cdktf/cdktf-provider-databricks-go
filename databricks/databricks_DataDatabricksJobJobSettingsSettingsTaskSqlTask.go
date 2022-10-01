// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsTaskSqlTask struct {
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#alert DataDatabricksJob#alert}
	Alert *DataDatabricksJobJobSettingsSettingsTaskSqlTaskAlert `field:"optional" json:"alert" yaml:"alert"`
	// dashboard block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#dashboard DataDatabricksJob#dashboard}
	Dashboard *DataDatabricksJobJobSettingsSettingsTaskSqlTaskDashboard `field:"optional" json:"dashboard" yaml:"dashboard"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#parameters DataDatabricksJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#query DataDatabricksJob#query}
	Query *DataDatabricksJobJobSettingsSettingsTaskSqlTaskQuery `field:"optional" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

