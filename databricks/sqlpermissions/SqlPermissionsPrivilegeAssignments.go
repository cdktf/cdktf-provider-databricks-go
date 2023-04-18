package sqlpermissions


type SqlPermissionsPrivilegeAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_permissions#principal SqlPermissions#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.14.3/docs/resources/sql_permissions#privileges SqlPermissions#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

