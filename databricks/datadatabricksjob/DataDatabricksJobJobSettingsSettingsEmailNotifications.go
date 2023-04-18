package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/job#alert_on_last_attempt DataDatabricksJob#alert_on_last_attempt}.
	AlertOnLastAttempt interface{} `field:"optional" json:"alertOnLastAttempt" yaml:"alertOnLastAttempt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/job#no_alert_for_skipped_runs DataDatabricksJob#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/job#on_failure DataDatabricksJob#on_failure}.
	OnFailure *[]*string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/job#on_start DataDatabricksJob#on_start}.
	OnStart *[]*string `field:"optional" json:"onStart" yaml:"onStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/data-sources/job#on_success DataDatabricksJob#on_success}.
	OnSuccess *[]*string `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}

