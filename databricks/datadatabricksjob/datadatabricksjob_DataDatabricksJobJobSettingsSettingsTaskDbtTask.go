package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskDbtTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#commands DataDatabricksJob#commands}.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#profiles_directory DataDatabricksJob#profiles_directory}.
	ProfilesDirectory *string `field:"optional" json:"profilesDirectory" yaml:"profilesDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#project_directory DataDatabricksJob#project_directory}.
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#schema DataDatabricksJob#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#warehouse_id DataDatabricksJob#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}
