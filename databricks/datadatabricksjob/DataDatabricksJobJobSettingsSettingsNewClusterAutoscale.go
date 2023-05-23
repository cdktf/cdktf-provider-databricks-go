package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsNewClusterAutoscale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/job#max_workers DataDatabricksJob#max_workers}.
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/job#min_workers DataDatabricksJob#min_workers}.
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
}

