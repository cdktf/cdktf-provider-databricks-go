// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoInitScriptsS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#destination DataDatabricksClusterPluginframework#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#canned_acl DataDatabricksClusterPluginframework#canned_acl}.
	CannedAcl *string `field:"optional" json:"cannedAcl" yaml:"cannedAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#enable_encryption DataDatabricksClusterPluginframework#enable_encryption}.
	EnableEncryption interface{} `field:"optional" json:"enableEncryption" yaml:"enableEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#encryption_type DataDatabricksClusterPluginframework#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#endpoint DataDatabricksClusterPluginframework#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#kms_key DataDatabricksClusterPluginframework#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.75.0/docs/data-sources/cluster_pluginframework#region DataDatabricksClusterPluginframework#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

