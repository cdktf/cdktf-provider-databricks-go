// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksaccountnetworkpolicies


type DataDatabricksAccountNetworkPoliciesItemsEgressNetworkAccessAllowedStorageDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/account_network_policies#azure_storage_account DataDatabricksAccountNetworkPolicies#azure_storage_account}.
	AzureStorageAccount *string `field:"optional" json:"azureStorageAccount" yaml:"azureStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/account_network_policies#azure_storage_service DataDatabricksAccountNetworkPolicies#azure_storage_service}.
	AzureStorageService *string `field:"optional" json:"azureStorageService" yaml:"azureStorageService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/account_network_policies#bucket_name DataDatabricksAccountNetworkPolicies#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/account_network_policies#region DataDatabricksAccountNetworkPolicies#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/data-sources/account_network_policies#storage_destination_type DataDatabricksAccountNetworkPolicies#storage_destination_type}.
	StorageDestinationType *string `field:"optional" json:"storageDestinationType" yaml:"storageDestinationType"`
}

