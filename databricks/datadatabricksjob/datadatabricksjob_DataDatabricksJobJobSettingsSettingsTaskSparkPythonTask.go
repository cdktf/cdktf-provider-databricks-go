package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskSparkPythonTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#python_file DataDatabricksJob#python_file}.
	PythonFile *string `field:"required" json:"pythonFile" yaml:"pythonFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

