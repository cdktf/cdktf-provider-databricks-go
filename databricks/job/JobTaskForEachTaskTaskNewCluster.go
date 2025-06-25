// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskForEachTaskTaskNewCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#apply_policy_default_values Job#apply_policy_default_values}.
	ApplyPolicyDefaultValues interface{} `field:"optional" json:"applyPolicyDefaultValues" yaml:"applyPolicyDefaultValues"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#autoscale Job#autoscale}
	Autoscale *JobTaskForEachTaskTaskNewClusterAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#aws_attributes Job#aws_attributes}
	AwsAttributes *JobTaskForEachTaskTaskNewClusterAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#azure_attributes Job#azure_attributes}
	AzureAttributes *JobTaskForEachTaskTaskNewClusterAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#cluster_id Job#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#cluster_log_conf Job#cluster_log_conf}
	ClusterLogConf *JobTaskForEachTaskTaskNewClusterClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// cluster_mount_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#cluster_mount_info Job#cluster_mount_info}
	ClusterMountInfo interface{} `field:"optional" json:"clusterMountInfo" yaml:"clusterMountInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#cluster_name Job#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#custom_tags Job#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#data_security_mode Job#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#docker_image Job#docker_image}
	DockerImage *JobTaskForEachTaskTaskNewClusterDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#driver_instance_pool_id Job#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#driver_node_type_id Job#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#enable_elastic_disk Job#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#enable_local_disk_encryption Job#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#gcp_attributes Job#gcp_attributes}
	GcpAttributes *JobTaskForEachTaskTaskNewClusterGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#idempotency_token Job#idempotency_token}.
	IdempotencyToken *string `field:"optional" json:"idempotencyToken" yaml:"idempotencyToken"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#init_scripts Job#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#instance_pool_id Job#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#is_single_node Job#is_single_node}.
	IsSingleNode interface{} `field:"optional" json:"isSingleNode" yaml:"isSingleNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#kind Job#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#library Job#library}
	Library interface{} `field:"optional" json:"library" yaml:"library"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#node_type_id Job#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#num_workers Job#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#policy_id Job#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#remote_disk_throughput Job#remote_disk_throughput}.
	RemoteDiskThroughput *float64 `field:"optional" json:"remoteDiskThroughput" yaml:"remoteDiskThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#runtime_engine Job#runtime_engine}.
	RuntimeEngine *string `field:"optional" json:"runtimeEngine" yaml:"runtimeEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#single_user_name Job#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#spark_conf Job#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#spark_env_vars Job#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#spark_version Job#spark_version}.
	SparkVersion *string `field:"optional" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#ssh_public_keys Job#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#total_initial_remote_disk_size Job#total_initial_remote_disk_size}.
	TotalInitialRemoteDiskSize *float64 `field:"optional" json:"totalInitialRemoteDiskSize" yaml:"totalInitialRemoteDiskSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#use_ml_runtime Job#use_ml_runtime}.
	UseMlRuntime interface{} `field:"optional" json:"useMlRuntime" yaml:"useMlRuntime"`
	// workload_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.84.0/docs/resources/job#workload_type Job#workload_type}
	WorkloadType *JobTaskForEachTaskTaskNewClusterWorkloadType `field:"optional" json:"workloadType" yaml:"workloadType"`
}

