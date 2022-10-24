package share


type ShareObject struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#data_object_type Share#data_object_type}.
	DataObjectType *string `field:"required" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#name Share#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#added_at Share#added_at}.
	AddedAt *float64 `field:"optional" json:"addedAt" yaml:"addedAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#added_by Share#added_by}.
	AddedBy *string `field:"optional" json:"addedBy" yaml:"addedBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#comment Share#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/share#shared_as Share#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
}

