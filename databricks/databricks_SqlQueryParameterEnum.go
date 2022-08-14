// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlQueryParameterEnum struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#options SqlQuery#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
	// multiple block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#multiple SqlQuery#multiple}
	Multiple *SqlQueryParameterEnumMultiple `field:"optional" json:"multiple" yaml:"multiple"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#value SqlQuery#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_query#values SqlQuery#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

