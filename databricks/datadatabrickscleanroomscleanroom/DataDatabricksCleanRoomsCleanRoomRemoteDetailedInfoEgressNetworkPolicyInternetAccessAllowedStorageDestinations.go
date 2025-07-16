// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanroom


type DataDatabricksCleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#allowed_paths DataDatabricksCleanRoomsCleanRoom#allowed_paths}.
	AllowedPaths *[]*string `field:"optional" json:"allowedPaths" yaml:"allowedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#azure_container DataDatabricksCleanRoomsCleanRoom#azure_container}.
	AzureContainer *string `field:"optional" json:"azureContainer" yaml:"azureContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#azure_dns_zone DataDatabricksCleanRoomsCleanRoom#azure_dns_zone}.
	AzureDnsZone *string `field:"optional" json:"azureDnsZone" yaml:"azureDnsZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#azure_storage_account DataDatabricksCleanRoomsCleanRoom#azure_storage_account}.
	AzureStorageAccount *string `field:"optional" json:"azureStorageAccount" yaml:"azureStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#azure_storage_service DataDatabricksCleanRoomsCleanRoom#azure_storage_service}.
	AzureStorageService *string `field:"optional" json:"azureStorageService" yaml:"azureStorageService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#bucket_name DataDatabricksCleanRoomsCleanRoom#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#region DataDatabricksCleanRoomsCleanRoom#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/clean_rooms_clean_room#type DataDatabricksCleanRoomsCleanRoom#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

