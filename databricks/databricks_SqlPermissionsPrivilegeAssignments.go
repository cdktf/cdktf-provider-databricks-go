// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type SqlPermissionsPrivilegeAssignments struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_permissions#principal SqlPermissions#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/r/sql_permissions#privileges SqlPermissions#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

