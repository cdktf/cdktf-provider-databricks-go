// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscurrentmetastore


type DataDatabricksCurrentMetastoreMetastoreInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#cloud DataDatabricksCurrentMetastore#cloud}.
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#created_at DataDatabricksCurrentMetastore#created_at}.
	CreatedAt *float64 `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#created_by DataDatabricksCurrentMetastore#created_by}.
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#default_data_access_config_id DataDatabricksCurrentMetastore#default_data_access_config_id}.
	DefaultDataAccessConfigId *string `field:"optional" json:"defaultDataAccessConfigId" yaml:"defaultDataAccessConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#delta_sharing_organization_name DataDatabricksCurrentMetastore#delta_sharing_organization_name}.
	DeltaSharingOrganizationName *string `field:"optional" json:"deltaSharingOrganizationName" yaml:"deltaSharingOrganizationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#delta_sharing_recipient_token_lifetime_in_seconds DataDatabricksCurrentMetastore#delta_sharing_recipient_token_lifetime_in_seconds}.
	DeltaSharingRecipientTokenLifetimeInSeconds *float64 `field:"optional" json:"deltaSharingRecipientTokenLifetimeInSeconds" yaml:"deltaSharingRecipientTokenLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#delta_sharing_scope DataDatabricksCurrentMetastore#delta_sharing_scope}.
	DeltaSharingScope *string `field:"optional" json:"deltaSharingScope" yaml:"deltaSharingScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#external_access_enabled DataDatabricksCurrentMetastore#external_access_enabled}.
	ExternalAccessEnabled interface{} `field:"optional" json:"externalAccessEnabled" yaml:"externalAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#global_metastore_id DataDatabricksCurrentMetastore#global_metastore_id}.
	GlobalMetastoreId *string `field:"optional" json:"globalMetastoreId" yaml:"globalMetastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#metastore_id DataDatabricksCurrentMetastore#metastore_id}.
	MetastoreId *string `field:"optional" json:"metastoreId" yaml:"metastoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#name DataDatabricksCurrentMetastore#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#owner DataDatabricksCurrentMetastore#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#privilege_model_version DataDatabricksCurrentMetastore#privilege_model_version}.
	PrivilegeModelVersion *string `field:"optional" json:"privilegeModelVersion" yaml:"privilegeModelVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#region DataDatabricksCurrentMetastore#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#storage_root DataDatabricksCurrentMetastore#storage_root}.
	StorageRoot *string `field:"optional" json:"storageRoot" yaml:"storageRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#storage_root_credential_id DataDatabricksCurrentMetastore#storage_root_credential_id}.
	StorageRootCredentialId *string `field:"optional" json:"storageRootCredentialId" yaml:"storageRootCredentialId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#storage_root_credential_name DataDatabricksCurrentMetastore#storage_root_credential_name}.
	StorageRootCredentialName *string `field:"optional" json:"storageRootCredentialName" yaml:"storageRootCredentialName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#updated_at DataDatabricksCurrentMetastore#updated_at}.
	UpdatedAt *float64 `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.95.0/docs/data-sources/current_metastore#updated_by DataDatabricksCurrentMetastore#updated_by}.
	UpdatedBy *string `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

