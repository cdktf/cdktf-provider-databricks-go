// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type GrantsGrant struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#principal Grants#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/grants#privileges Grants#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

