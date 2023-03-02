package share


type ShareObjectPartition struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#value Share#value}
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

