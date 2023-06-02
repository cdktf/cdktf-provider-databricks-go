package job


type JobSparkPythonTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#python_file Job#python_file}.
	PythonFile *string `field:"required" json:"pythonFile" yaml:"pythonFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.18.0/docs/resources/job#source Job#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

