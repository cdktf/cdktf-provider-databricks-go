package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsSparkPythonTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#python_file DataDatabricksJob#python_file}.
	PythonFile *string `field:"required" json:"pythonFile" yaml:"pythonFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#parameters DataDatabricksJob#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#source DataDatabricksJob#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}
