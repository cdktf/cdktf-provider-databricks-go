package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterAutoscale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#max_workers DataDatabricksJob#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#min_workers DataDatabricksJob#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
}

