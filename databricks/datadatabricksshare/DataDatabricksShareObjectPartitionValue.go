package datadatabricksshare


type DataDatabricksShareObjectPartitionValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/share#name DataDatabricksShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/share#op DataDatabricksShare#op}.
	Op *string `field:"required" json:"op" yaml:"op"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/share#recipient_property_key DataDatabricksShare#recipient_property_key}.
	RecipientPropertyKey *string `field:"optional" json:"recipientPropertyKey" yaml:"recipientPropertyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.17.0/docs/data-sources/share#value DataDatabricksShare#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

