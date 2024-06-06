// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#default_tags DataDatabricksCluster#default_tags}.
	DefaultTags *map[string]*string `field:"required" json:"defaultTags" yaml:"defaultTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#spark_version DataDatabricksCluster#spark_version}.
	SparkVersion *string `field:"required" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#state DataDatabricksCluster#state}.
	State *string `field:"required" json:"state" yaml:"state"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#autoscale DataDatabricksCluster#autoscale}
	Autoscale *DataDatabricksClusterClusterInfoAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#autotermination_minutes DataDatabricksCluster#autotermination_minutes}.
	AutoterminationMinutes *float64 `field:"optional" json:"autoterminationMinutes" yaml:"autoterminationMinutes"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#aws_attributes DataDatabricksCluster#aws_attributes}
	AwsAttributes *DataDatabricksClusterClusterInfoAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#azure_attributes DataDatabricksCluster#azure_attributes}
	AzureAttributes *DataDatabricksClusterClusterInfoAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_cores DataDatabricksCluster#cluster_cores}.
	ClusterCores *float64 `field:"optional" json:"clusterCores" yaml:"clusterCores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_id DataDatabricksCluster#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_log_conf DataDatabricksCluster#cluster_log_conf}
	ClusterLogConf *DataDatabricksClusterClusterInfoClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// cluster_log_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_log_status DataDatabricksCluster#cluster_log_status}
	ClusterLogStatus *DataDatabricksClusterClusterInfoClusterLogStatus `field:"optional" json:"clusterLogStatus" yaml:"clusterLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_memory_mb DataDatabricksCluster#cluster_memory_mb}.
	ClusterMemoryMb *float64 `field:"optional" json:"clusterMemoryMb" yaml:"clusterMemoryMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#cluster_name DataDatabricksCluster#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#creator_user_name DataDatabricksCluster#creator_user_name}.
	CreatorUserName *string `field:"optional" json:"creatorUserName" yaml:"creatorUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#custom_tags DataDatabricksCluster#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#data_security_mode DataDatabricksCluster#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#docker_image DataDatabricksCluster#docker_image}
	DockerImage *DataDatabricksClusterClusterInfoDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#driver DataDatabricksCluster#driver}
	Driver *DataDatabricksClusterClusterInfoDriver `field:"optional" json:"driver" yaml:"driver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#driver_instance_pool_id DataDatabricksCluster#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#driver_node_type_id DataDatabricksCluster#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#enable_elastic_disk DataDatabricksCluster#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#enable_local_disk_encryption DataDatabricksCluster#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// executors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#executors DataDatabricksCluster#executors}
	Executors interface{} `field:"optional" json:"executors" yaml:"executors"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#gcp_attributes DataDatabricksCluster#gcp_attributes}
	GcpAttributes *DataDatabricksClusterClusterInfoGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#init_scripts DataDatabricksCluster#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#instance_pool_id DataDatabricksCluster#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#jdbc_port DataDatabricksCluster#jdbc_port}.
	JdbcPort *float64 `field:"optional" json:"jdbcPort" yaml:"jdbcPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#last_activity_time DataDatabricksCluster#last_activity_time}.
	LastActivityTime *float64 `field:"optional" json:"lastActivityTime" yaml:"lastActivityTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#last_state_loss_time DataDatabricksCluster#last_state_loss_time}.
	LastStateLossTime *float64 `field:"optional" json:"lastStateLossTime" yaml:"lastStateLossTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#node_type_id DataDatabricksCluster#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#num_workers DataDatabricksCluster#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#policy_id DataDatabricksCluster#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#runtime_engine DataDatabricksCluster#runtime_engine}.
	RuntimeEngine *string `field:"optional" json:"runtimeEngine" yaml:"runtimeEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#single_user_name DataDatabricksCluster#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#spark_conf DataDatabricksCluster#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#spark_context_id DataDatabricksCluster#spark_context_id}.
	SparkContextId *float64 `field:"optional" json:"sparkContextId" yaml:"sparkContextId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#spark_env_vars DataDatabricksCluster#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#ssh_public_keys DataDatabricksCluster#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#start_time DataDatabricksCluster#start_time}.
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#state_message DataDatabricksCluster#state_message}.
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#terminate_time DataDatabricksCluster#terminate_time}.
	TerminateTime *float64 `field:"optional" json:"terminateTime" yaml:"terminateTime"`
	// termination_reason block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/data-sources/cluster#termination_reason DataDatabricksCluster#termination_reason}
	TerminationReason *DataDatabricksClusterClusterInfoTerminationReason `field:"optional" json:"terminationReason" yaml:"terminationReason"`
}

