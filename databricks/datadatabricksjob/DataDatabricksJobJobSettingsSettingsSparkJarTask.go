package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsSparkJarTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#jar_uri DataDatabricksJob#jar_uri}.
	JarUri *string `field:"optional" json:"jarUri" yaml:"jarUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#main_class_name DataDatabricksJob#main_class_name}.
	MainClassName *string `field:"optional" json:"mainClassName" yaml:"mainClassName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

