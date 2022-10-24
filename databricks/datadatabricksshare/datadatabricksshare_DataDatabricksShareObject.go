package datadatabricksshare


type DataDatabricksShareObject struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#data_object_type DataDatabricksShare#data_object_type}.
	DataObjectType *string `field:"required" json:"dataObjectType" yaml:"dataObjectType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#name DataDatabricksShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#added_at DataDatabricksShare#added_at}.
	AddedAt *float64 `field:"optional" json:"addedAt" yaml:"addedAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#added_by DataDatabricksShare#added_by}.
	AddedBy *string `field:"optional" json:"addedBy" yaml:"addedBy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#comment DataDatabricksShare#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/share#shared_as DataDatabricksShare#shared_as}.
	SharedAs *string `field:"optional" json:"sharedAs" yaml:"sharedAs"`
}

