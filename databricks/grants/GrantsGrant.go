package grants


type GrantsGrant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/grants#principal Grants#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.20.0/docs/resources/grants#privileges Grants#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

