// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#apply_policy_default_values DataDatabricksCluster#apply_policy_default_values}.
	ApplyPolicyDefaultValues interface{} `field:"optional" json:"applyPolicyDefaultValues" yaml:"applyPolicyDefaultValues"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#autoscale DataDatabricksCluster#autoscale}
	Autoscale *DataDatabricksClusterClusterInfoSpecAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#aws_attributes DataDatabricksCluster#aws_attributes}
	AwsAttributes *DataDatabricksClusterClusterInfoSpecAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#azure_attributes DataDatabricksCluster#azure_attributes}
	AzureAttributes *DataDatabricksClusterClusterInfoSpecAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#cluster_log_conf DataDatabricksCluster#cluster_log_conf}
	ClusterLogConf *DataDatabricksClusterClusterInfoSpecClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// cluster_mount_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#cluster_mount_info DataDatabricksCluster#cluster_mount_info}
	ClusterMountInfo interface{} `field:"optional" json:"clusterMountInfo" yaml:"clusterMountInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#cluster_name DataDatabricksCluster#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#custom_tags DataDatabricksCluster#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#data_security_mode DataDatabricksCluster#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#docker_image DataDatabricksCluster#docker_image}
	DockerImage *DataDatabricksClusterClusterInfoSpecDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#driver_instance_pool_id DataDatabricksCluster#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#driver_node_type_id DataDatabricksCluster#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#enable_elastic_disk DataDatabricksCluster#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#enable_local_disk_encryption DataDatabricksCluster#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#gcp_attributes DataDatabricksCluster#gcp_attributes}
	GcpAttributes *DataDatabricksClusterClusterInfoSpecGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#idempotency_token DataDatabricksCluster#idempotency_token}.
	IdempotencyToken *string `field:"optional" json:"idempotencyToken" yaml:"idempotencyToken"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#init_scripts DataDatabricksCluster#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#instance_pool_id DataDatabricksCluster#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#is_single_node DataDatabricksCluster#is_single_node}.
	IsSingleNode interface{} `field:"optional" json:"isSingleNode" yaml:"isSingleNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#kind DataDatabricksCluster#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#library DataDatabricksCluster#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#node_type_id DataDatabricksCluster#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#num_workers DataDatabricksCluster#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#policy_id DataDatabricksCluster#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#runtime_engine DataDatabricksCluster#runtime_engine}.
	RuntimeEngine *string `field:"optional" json:"runtimeEngine" yaml:"runtimeEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#single_user_name DataDatabricksCluster#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#spark_conf DataDatabricksCluster#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#spark_env_vars DataDatabricksCluster#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#spark_version DataDatabricksCluster#spark_version}.
	SparkVersion *string `field:"optional" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#ssh_public_keys DataDatabricksCluster#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#use_ml_runtime DataDatabricksCluster#use_ml_runtime}.
	UseMlRuntime interface{} `field:"optional" json:"useMlRuntime" yaml:"useMlRuntime"`
	// workload_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/data-sources/cluster#workload_type DataDatabricksCluster#workload_type}
	WorkloadType *DataDatabricksClusterClusterInfoSpecWorkloadType `field:"optional" json:"workloadType" yaml:"workloadType"`
}

