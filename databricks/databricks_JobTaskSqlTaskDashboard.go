// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobTaskSqlTaskDashboard struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#dashboard_id Job#dashboard_id}.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
}

