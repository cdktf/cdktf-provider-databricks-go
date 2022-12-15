package job


type JobTaskDbtTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#commands Job#commands}.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#catalog Job#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#profiles_directory Job#profiles_directory}.
	ProfilesDirectory *string `field:"optional" json:"profilesDirectory" yaml:"profilesDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#project_directory Job#project_directory}.
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#schema Job#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#warehouse_id Job#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

