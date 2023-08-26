// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metastoredataaccess


type MetastoreDataAccessAzureManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.0/docs/resources/metastore_data_access#access_connector_id MetastoreDataAccess#access_connector_id}.
	AccessConnectorId *string `field:"required" json:"accessConnectorId" yaml:"accessConnectorId"`
}

