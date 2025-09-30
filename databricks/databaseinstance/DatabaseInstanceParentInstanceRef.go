// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseinstance


type DatabaseInstanceParentInstanceRef struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/database_instance#branch_time DatabaseInstance#branch_time}.
	BranchTime *string `field:"optional" json:"branchTime" yaml:"branchTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/database_instance#lsn DatabaseInstance#lsn}.
	Lsn *string `field:"optional" json:"lsn" yaml:"lsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/database_instance#name DatabaseInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

