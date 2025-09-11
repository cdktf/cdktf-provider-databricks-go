// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscleanroomscleanrooms


type DataDatabricksCleanRoomsCleanRoomsCleanRoomsRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#allowed_paths DataDatabricksCleanRoomsCleanRooms#allowed_paths}.
	AllowedPaths *[]*string `field:"optional" json:"allowedPaths" yaml:"allowedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#azure_container DataDatabricksCleanRoomsCleanRooms#azure_container}.
	AzureContainer *string `field:"optional" json:"azureContainer" yaml:"azureContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#azure_dns_zone DataDatabricksCleanRoomsCleanRooms#azure_dns_zone}.
	AzureDnsZone *string `field:"optional" json:"azureDnsZone" yaml:"azureDnsZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#azure_storage_account DataDatabricksCleanRoomsCleanRooms#azure_storage_account}.
	AzureStorageAccount *string `field:"optional" json:"azureStorageAccount" yaml:"azureStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#azure_storage_service DataDatabricksCleanRoomsCleanRooms#azure_storage_service}.
	AzureStorageService *string `field:"optional" json:"azureStorageService" yaml:"azureStorageService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#bucket_name DataDatabricksCleanRoomsCleanRooms#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#region DataDatabricksCleanRoomsCleanRooms#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/data-sources/clean_rooms_clean_rooms#type DataDatabricksCleanRoomsCleanRooms#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

