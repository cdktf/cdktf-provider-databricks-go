package job


type JobTaskSqlTaskQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.19.0/docs/resources/job#query_id Job#query_id}.
	QueryId *string `field:"required" json:"queryId" yaml:"queryId"`
}

