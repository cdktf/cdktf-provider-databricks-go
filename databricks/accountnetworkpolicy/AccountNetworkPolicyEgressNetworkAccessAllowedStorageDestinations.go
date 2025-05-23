// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountnetworkpolicy


type AccountNetworkPolicyEgressNetworkAccessAllowedStorageDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/account_network_policy#azure_storage_account AccountNetworkPolicy#azure_storage_account}.
	AzureStorageAccount *string `field:"optional" json:"azureStorageAccount" yaml:"azureStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/account_network_policy#azure_storage_service AccountNetworkPolicy#azure_storage_service}.
	AzureStorageService *string `field:"optional" json:"azureStorageService" yaml:"azureStorageService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/account_network_policy#bucket_name AccountNetworkPolicy#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/account_network_policy#region AccountNetworkPolicy#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/account_network_policy#storage_destination_type AccountNetworkPolicy#storage_destination_type}.
	StorageDestinationType *string `field:"optional" json:"storageDestinationType" yaml:"storageDestinationType"`
}

