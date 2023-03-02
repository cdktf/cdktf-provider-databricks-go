package datadatabricksshare


type DataDatabricksShareObjectPartitionValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#name DataDatabricksShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#op DataDatabricksShare#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#recipient_property_key DataDatabricksShare#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#value DataDatabricksShare#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

