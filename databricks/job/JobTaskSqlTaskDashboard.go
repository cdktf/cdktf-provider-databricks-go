package job


type JobTaskSqlTaskDashboard struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.16.1/docs/resources/job#dashboard_id Job#dashboard_id}.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
}

