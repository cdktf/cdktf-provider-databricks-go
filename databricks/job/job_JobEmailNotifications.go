package job


type JobEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#alert_on_last_attempt Job#alert_on_last_attempt}.
	AlertOnLastAttempt interface{} `field:"optional" json:"alertOnLastAttempt" yaml:"alertOnLastAttempt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#no_alert_for_skipped_runs Job#no_alert_for_skipped_runs}.
	NoAlertForSkippedRuns interface{} `field:"optional" json:"noAlertForSkippedRuns" yaml:"noAlertForSkippedRuns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#on_failure Job#on_failure}.
	OnFailure *[]*string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#on_start Job#on_start}.
	OnStart *[]*string `field:"optional" json:"onStart" yaml:"onStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#on_success Job#on_success}.
	OnSuccess *[]*string `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}
