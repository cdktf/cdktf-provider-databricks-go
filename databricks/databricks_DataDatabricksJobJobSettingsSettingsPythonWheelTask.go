// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsPythonWheelTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#entry_point DataDatabricksJob#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#named_parameters DataDatabricksJob#named_parameters}.
	NamedParameters *map[string]*string `field:"optional" json:"namedParameters" yaml:"namedParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#package_name DataDatabricksJob#package_name}.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

