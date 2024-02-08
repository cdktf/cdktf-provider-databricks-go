// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess


type MetastoreDataAccessDatabricksGcpServiceAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/metastore_data_access#credential_id MetastoreDataAccess#credential_id}.
	CredentialId *string `field:"optional" json:"credentialId" yaml:"credentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.1/docs/resources/metastore_data_access#email MetastoreDataAccess#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
}

