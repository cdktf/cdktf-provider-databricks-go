// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlpermissions


type SqlPermissionsPrivilegeAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/sql_permissions#principal SqlPermissions#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.83.0/docs/resources/sql_permissions#privileges SqlPermissions#privileges}.
	Privileges *[]*string `field:"required" json:"privileges" yaml:"privileges"`
}

