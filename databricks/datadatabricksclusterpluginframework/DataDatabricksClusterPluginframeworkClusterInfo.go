// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksclusterpluginframework


type DataDatabricksClusterPluginframeworkClusterInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#autoscale DataDatabricksClusterPluginframework#autoscale}.
	Autoscale interface{} `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#autotermination_minutes DataDatabricksClusterPluginframework#autotermination_minutes}.
	AutoterminationMinutes *float64 `field:"optional" json:"autoterminationMinutes" yaml:"autoterminationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#aws_attributes DataDatabricksClusterPluginframework#aws_attributes}.
	AwsAttributes interface{} `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#azure_attributes DataDatabricksClusterPluginframework#azure_attributes}.
	AzureAttributes interface{} `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_cores DataDatabricksClusterPluginframework#cluster_cores}.
	ClusterCores *float64 `field:"optional" json:"clusterCores" yaml:"clusterCores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_id DataDatabricksClusterPluginframework#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_log_conf DataDatabricksClusterPluginframework#cluster_log_conf}.
	ClusterLogConf interface{} `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_log_status DataDatabricksClusterPluginframework#cluster_log_status}.
	ClusterLogStatus interface{} `field:"optional" json:"clusterLogStatus" yaml:"clusterLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_memory_mb DataDatabricksClusterPluginframework#cluster_memory_mb}.
	ClusterMemoryMb *float64 `field:"optional" json:"clusterMemoryMb" yaml:"clusterMemoryMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_name DataDatabricksClusterPluginframework#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#cluster_source DataDatabricksClusterPluginframework#cluster_source}.
	ClusterSource *string `field:"optional" json:"clusterSource" yaml:"clusterSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#creator_user_name DataDatabricksClusterPluginframework#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#custom_tags DataDatabricksClusterPluginframework#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#data_security_mode DataDatabricksClusterPluginframework#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#default_tags DataDatabricksClusterPluginframework#default_tags}.
	DefaultTags *map[string]*string `field:"optional" json:"defaultTags" yaml:"defaultTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#docker_image DataDatabricksClusterPluginframework#docker_image}.
	DockerImage interface{} `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#driver DataDatabricksClusterPluginframework#driver}.
	Driver interface{} `field:"optional" json:"driver" yaml:"driver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#driver_instance_pool_id DataDatabricksClusterPluginframework#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#driver_node_type_id DataDatabricksClusterPluginframework#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#enable_elastic_disk DataDatabricksClusterPluginframework#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#enable_local_disk_encryption DataDatabricksClusterPluginframework#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#executors DataDatabricksClusterPluginframework#executors}.
	Executors interface{} `field:"optional" json:"executors" yaml:"executors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#gcp_attributes DataDatabricksClusterPluginframework#gcp_attributes}.
	GcpAttributes interface{} `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#init_scripts DataDatabricksClusterPluginframework#init_scripts}.
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#instance_pool_id DataDatabricksClusterPluginframework#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#is_single_node DataDatabricksClusterPluginframework#is_single_node}.
	IsSingleNode interface{} `field:"optional" json:"isSingleNode" yaml:"isSingleNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#jdbc_port DataDatabricksClusterPluginframework#jdbc_port}.
	JdbcPort *float64 `field:"optional" json:"jdbcPort" yaml:"jdbcPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#kind DataDatabricksClusterPluginframework#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#last_restarted_time DataDatabricksClusterPluginframework#last_restarted_time}.
	LastRestartedTime *float64 `field:"optional" json:"lastRestartedTime" yaml:"lastRestartedTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#last_state_loss_time DataDatabricksClusterPluginframework#last_state_loss_time}.
	LastStateLossTime *float64 `field:"optional" json:"lastStateLossTime" yaml:"lastStateLossTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#node_type_id DataDatabricksClusterPluginframework#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#num_workers DataDatabricksClusterPluginframework#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#policy_id DataDatabricksClusterPluginframework#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#runtime_engine DataDatabricksClusterPluginframework#runtime_engine}.
	RuntimeEngine *string `field:"optional" json:"runtimeEngine" yaml:"runtimeEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#single_user_name DataDatabricksClusterPluginframework#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spark_conf DataDatabricksClusterPluginframework#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spark_context_id DataDatabricksClusterPluginframework#spark_context_id}.
	SparkContextId *float64 `field:"optional" json:"sparkContextId" yaml:"sparkContextId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spark_env_vars DataDatabricksClusterPluginframework#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spark_version DataDatabricksClusterPluginframework#spark_version}.
	SparkVersion *string `field:"optional" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#spec DataDatabricksClusterPluginframework#spec}.
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#ssh_public_keys DataDatabricksClusterPluginframework#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#start_time DataDatabricksClusterPluginframework#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#state DataDatabricksClusterPluginframework#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#state_message DataDatabricksClusterPluginframework#state_message}.
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#terminated_time DataDatabricksClusterPluginframework#terminated_time}.
	TerminatedTime *float64 `field:"optional" json:"terminatedTime" yaml:"terminatedTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#termination_reason DataDatabricksClusterPluginframework#termination_reason}.
	TerminationReason interface{} `field:"optional" json:"terminationReason" yaml:"terminationReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#use_ml_runtime DataDatabricksClusterPluginframework#use_ml_runtime}.
	UseMlRuntime interface{} `field:"optional" json:"useMlRuntime" yaml:"useMlRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.77.0/docs/data-sources/cluster_pluginframework#workload_type DataDatabricksClusterPluginframework#workload_type}.
	WorkloadType interface{} `field:"optional" json:"workloadType" yaml:"workloadType"`
}

