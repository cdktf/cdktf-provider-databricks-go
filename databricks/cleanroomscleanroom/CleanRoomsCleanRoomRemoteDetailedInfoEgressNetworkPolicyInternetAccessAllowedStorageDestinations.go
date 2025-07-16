// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfoEgressNetworkPolicyInternetAccessAllowedStorageDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#allowed_paths CleanRoomsCleanRoom#allowed_paths}.
	AllowedPaths *[]*string `field:"optional" json:"allowedPaths" yaml:"allowedPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#azure_container CleanRoomsCleanRoom#azure_container}.
	AzureContainer *string `field:"optional" json:"azureContainer" yaml:"azureContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#azure_dns_zone CleanRoomsCleanRoom#azure_dns_zone}.
	AzureDnsZone *string `field:"optional" json:"azureDnsZone" yaml:"azureDnsZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#azure_storage_account CleanRoomsCleanRoom#azure_storage_account}.
	AzureStorageAccount *string `field:"optional" json:"azureStorageAccount" yaml:"azureStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#azure_storage_service CleanRoomsCleanRoom#azure_storage_service}.
	AzureStorageService *string `field:"optional" json:"azureStorageService" yaml:"azureStorageService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#bucket_name CleanRoomsCleanRoom#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#region CleanRoomsCleanRoom#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/resources/clean_rooms_clean_room#type CleanRoomsCleanRoom#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

