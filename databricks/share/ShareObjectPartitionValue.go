package share


type ShareObjectPartitionValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#name Share#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#op Share#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#recipient_property_key Share#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#value Share#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

