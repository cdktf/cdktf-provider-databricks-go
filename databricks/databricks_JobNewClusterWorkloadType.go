// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#clients Job#clients}
	Clients *JobNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

