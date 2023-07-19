package share


type ShareObjectPartition struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.21.0/docs/resources/share#value Share#value}
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

