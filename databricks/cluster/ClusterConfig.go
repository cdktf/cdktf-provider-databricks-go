// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#spark_version Cluster#spark_version}.
	SparkVersion *string `field:"required" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#apply_policy_default_values Cluster#apply_policy_default_values}.
	ApplyPolicyDefaultValues interface{} `field:"optional" json:"applyPolicyDefaultValues" yaml:"applyPolicyDefaultValues"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#autoscale Cluster#autoscale}
	Autoscale *ClusterAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#autotermination_minutes Cluster#autotermination_minutes}.
	AutoterminationMinutes *float64 `field:"optional" json:"autoterminationMinutes" yaml:"autoterminationMinutes"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#aws_attributes Cluster#aws_attributes}
	AwsAttributes *ClusterAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#azure_attributes Cluster#azure_attributes}
	AzureAttributes *ClusterAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#cluster_log_conf Cluster#cluster_log_conf}
	ClusterLogConf *ClusterClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// cluster_mount_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#cluster_mount_info Cluster#cluster_mount_info}
	ClusterMountInfo interface{} `field:"optional" json:"clusterMountInfo" yaml:"clusterMountInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#cluster_name Cluster#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#custom_tags Cluster#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#data_security_mode Cluster#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#docker_image Cluster#docker_image}
	DockerImage *ClusterDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#driver_instance_pool_id Cluster#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#driver_node_type_id Cluster#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#enable_elastic_disk Cluster#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#enable_local_disk_encryption Cluster#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#gcp_attributes Cluster#gcp_attributes}
	GcpAttributes *ClusterGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#id Cluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#idempotency_token Cluster#idempotency_token}.
	IdempotencyToken *string `field:"optional" json:"idempotencyToken" yaml:"idempotencyToken"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#init_scripts Cluster#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#instance_pool_id Cluster#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#is_pinned Cluster#is_pinned}.
	IsPinned interface{} `field:"optional" json:"isPinned" yaml:"isPinned"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#is_single_node Cluster#is_single_node}.
	IsSingleNode interface{} `field:"optional" json:"isSingleNode" yaml:"isSingleNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#kind Cluster#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#library Cluster#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#node_type_id Cluster#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#no_wait Cluster#no_wait}.
	NoWait interface{} `field:"optional" json:"noWait" yaml:"noWait"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#num_workers Cluster#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#policy_id Cluster#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#remote_disk_throughput Cluster#remote_disk_throughput}.
	RemoteDiskThroughput *float64 `field:"optional" json:"remoteDiskThroughput" yaml:"remoteDiskThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#runtime_engine Cluster#runtime_engine}.
	RuntimeEngine *string `field:"optional" json:"runtimeEngine" yaml:"runtimeEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#single_user_name Cluster#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#spark_conf Cluster#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#spark_env_vars Cluster#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#ssh_public_keys Cluster#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#timeouts Cluster#timeouts}
	Timeouts *ClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#total_initial_remote_disk_size Cluster#total_initial_remote_disk_size}.
	TotalInitialRemoteDiskSize *float64 `field:"optional" json:"totalInitialRemoteDiskSize" yaml:"totalInitialRemoteDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#use_ml_runtime Cluster#use_ml_runtime}.
	UseMlRuntime interface{} `field:"optional" json:"useMlRuntime" yaml:"useMlRuntime"`
	// workload_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.91.0/docs/resources/cluster#workload_type Cluster#workload_type}
	WorkloadType *ClusterWorkloadType `field:"optional" json:"workloadType" yaml:"workloadType"`
}

