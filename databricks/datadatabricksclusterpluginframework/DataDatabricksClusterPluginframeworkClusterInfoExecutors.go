// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfoExecutors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#host_private_ip DataDatabricksClusterPluginframework#host_private_ip}.
	HostPrivateIp *string `field:"optional" json:"hostPrivateIp" yaml:"hostPrivateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#instance_id DataDatabricksClusterPluginframework#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#node_aws_attributes DataDatabricksClusterPluginframework#node_aws_attributes}.
	NodeAwsAttributes interface{} `field:"optional" json:"nodeAwsAttributes" yaml:"nodeAwsAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#node_id DataDatabricksClusterPluginframework#node_id}.
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#private_ip DataDatabricksClusterPluginframework#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#public_dns DataDatabricksClusterPluginframework#public_dns}.
	PublicDns *string `field:"optional" json:"publicDns" yaml:"publicDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/data-sources/cluster_pluginframework#start_timestamp DataDatabricksClusterPluginframework#start_timestamp}.
	StartTimestamp *float64 `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
}

