// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlWidgetPosition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_widget#size_x SqlWidget#size_x}.
	SizeX *float64 `field:"required" json:"sizeX" yaml:"sizeX"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_widget#size_y SqlWidget#size_y}.
	SizeY *float64 `field:"required" json:"sizeY" yaml:"sizeY"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_widget#auto_height SqlWidget#auto_height}.
	AutoHeight interface{} `field:"optional" json:"autoHeight" yaml:"autoHeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_widget#pos_x SqlWidget#pos_x}.
	PosX *float64 `field:"optional" json:"posX" yaml:"posX"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_widget#pos_y SqlWidget#pos_y}.
	PosY *float64 `field:"optional" json:"posY" yaml:"posY"`
}
