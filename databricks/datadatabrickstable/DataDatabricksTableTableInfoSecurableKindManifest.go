// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickstable


type DataDatabricksTableTableInfoSecurableKindManifest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/table#assignable_privileges DataDatabricksTable#assignable_privileges}.
	AssignablePrivileges *[]*string `field:"optional" json:"assignablePrivileges" yaml:"assignablePrivileges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/table#capabilities DataDatabricksTable#capabilities}.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/table#options DataDatabricksTable#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/table#securable_kind DataDatabricksTable#securable_kind}.
	SecurableKind *string `field:"optional" json:"securableKind" yaml:"securableKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.88.0/docs/data-sources/table#securable_type DataDatabricksTable#securable_type}.
	SecurableType *string `field:"optional" json:"securableType" yaml:"securableType"`
}

