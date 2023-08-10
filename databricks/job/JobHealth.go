package job


type JobHealth struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.23.0/docs/resources/job#rules Job#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

