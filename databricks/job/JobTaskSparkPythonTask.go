package job


type JobTaskSparkPythonTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#python_file Job#python_file}.
	PythonFile *string `field:"required" json:"pythonFile" yaml:"pythonFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}
