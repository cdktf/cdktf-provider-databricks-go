package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/job#compute_key DataDatabricksJob#compute_key}.
	ComputeKey *string `field:"optional" json:"computeKey" yaml:"computeKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.22.0/docs/data-sources/job#spec DataDatabricksJob#spec}
	Spec *DataDatabricksJobJobSettingsSettingsComputeSpec `field:"optional" json:"spec" yaml:"spec"`
}

