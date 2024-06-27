// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoExecutors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#host_private_ip DataDatabricksCluster#host_private_ip}.
	HostPrivateIp *string `field:"optional" json:"hostPrivateIp" yaml:"hostPrivateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#instance_id DataDatabricksCluster#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// node_aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#node_aws_attributes DataDatabricksCluster#node_aws_attributes}
	NodeAwsAttributes *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes `field:"optional" json:"nodeAwsAttributes" yaml:"nodeAwsAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#node_id DataDatabricksCluster#node_id}.
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#private_ip DataDatabricksCluster#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#public_dns DataDatabricksCluster#public_dns}.
	PublicDns *string `field:"optional" json:"publicDns" yaml:"publicDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.48.0/docs/data-sources/cluster#start_timestamp DataDatabricksCluster#start_timestamp}.
	StartTimestamp *float64 `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
}

