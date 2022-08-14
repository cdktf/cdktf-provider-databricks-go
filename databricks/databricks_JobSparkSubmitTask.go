// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobSparkSubmitTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#parameters Job#parameters}.
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

