// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type JobTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#create Job#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/job#update Job#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

