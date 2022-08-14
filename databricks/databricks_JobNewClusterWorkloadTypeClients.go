// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobNewClusterWorkloadTypeClients struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#jobs Job#jobs}.
	Jobs interface{} `field:"optional" json:"jobs" yaml:"jobs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#notebooks Job#notebooks}.
	Notebooks interface{} `field:"optional" json:"notebooks" yaml:"notebooks"`
}

