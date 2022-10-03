package job


type JobTaskSqlTaskQuery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#query_id Job#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
}

