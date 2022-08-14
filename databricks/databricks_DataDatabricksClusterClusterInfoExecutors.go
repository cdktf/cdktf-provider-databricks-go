// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksClusterClusterInfoExecutors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#host_private_ip DataDatabricksCluster#host_private_ip}.
	HostPrivateIp *string `field:"optional" json:"hostPrivateIp" yaml:"hostPrivateIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#instance_id DataDatabricksCluster#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// node_aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#node_aws_attributes DataDatabricksCluster#node_aws_attributes}
	NodeAwsAttributes *DataDatabricksClusterClusterInfoExecutorsNodeAwsAttributes `field:"optional" json:"nodeAwsAttributes" yaml:"nodeAwsAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#node_id DataDatabricksCluster#node_id}.
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#private_ip DataDatabricksCluster#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#public_dns DataDatabricksCluster#public_dns}.
	PublicDns *string `field:"optional" json:"publicDns" yaml:"publicDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/cluster#start_timestamp DataDatabricksCluster#start_timestamp}.
	StartTimestamp *float64 `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
}

