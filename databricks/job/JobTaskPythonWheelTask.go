package job


type JobTaskPythonWheelTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#entry_point Job#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#named_parameters Job#named_parameters}.
	NamedParameters *map[string]*string `field:"optional" json:"namedParameters" yaml:"namedParameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#package_name Job#package_name}.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

