// Prebuilt databricks Provider for Terraform CDK (cdktf)
package databricks


type DataDatabricksJobJobSettingsSettingsTaskNewCluster struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#num_workers DataDatabricksJob#num_workers}.
	NumWorkers *float64 `field:"required" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_version DataDatabricksJob#spark_version}.
	SparkVersion *string `field:"required" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#apply_policy_default_values DataDatabricksJob#apply_policy_default_values}.
	ApplyPolicyDefaultValues interface{} `field:"optional" json:"applyPolicyDefaultValues" yaml:"applyPolicyDefaultValues"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#autoscale DataDatabricksJob#autoscale}
	Autoscale *DataDatabricksJobJobSettingsSettingsTaskNewClusterAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#autotermination_minutes DataDatabricksJob#autotermination_minutes}.
	AutoterminationMinutes *float64 `field:"optional" json:"autoterminationMinutes" yaml:"autoterminationMinutes"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#aws_attributes DataDatabricksJob#aws_attributes}
	AwsAttributes *DataDatabricksJobJobSettingsSettingsTaskNewClusterAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#azure_attributes DataDatabricksJob#azure_attributes}
	AzureAttributes *DataDatabricksJobJobSettingsSettingsTaskNewClusterAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#cluster_id DataDatabricksJob#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#cluster_log_conf DataDatabricksJob#cluster_log_conf}
	ClusterLogConf *DataDatabricksJobJobSettingsSettingsTaskNewClusterClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#cluster_name DataDatabricksJob#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#custom_tags DataDatabricksJob#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#data_security_mode DataDatabricksJob#data_security_mode}.
	DataSecurityMode *string `field:"optional" json:"dataSecurityMode" yaml:"dataSecurityMode"`
	// docker_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#docker_image DataDatabricksJob#docker_image}
	DockerImage *DataDatabricksJobJobSettingsSettingsTaskNewClusterDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#driver_instance_pool_id DataDatabricksJob#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#driver_node_type_id DataDatabricksJob#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#enable_elastic_disk DataDatabricksJob#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#enable_local_disk_encryption DataDatabricksJob#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#gcp_attributes DataDatabricksJob#gcp_attributes}
	GcpAttributes *DataDatabricksJobJobSettingsSettingsTaskNewClusterGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#idempotency_token DataDatabricksJob#idempotency_token}.
	IdempotencyToken *string `field:"optional" json:"idempotencyToken" yaml:"idempotencyToken"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#init_scripts DataDatabricksJob#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#instance_pool_id DataDatabricksJob#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#node_type_id DataDatabricksJob#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#policy_id DataDatabricksJob#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#single_user_name DataDatabricksJob#single_user_name}.
	SingleUserName *string `field:"optional" json:"singleUserName" yaml:"singleUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_conf DataDatabricksJob#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#spark_env_vars DataDatabricksJob#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#ssh_public_keys DataDatabricksJob#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// workload_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks/d/job#workload_type DataDatabricksJob#workload_type}
	WorkloadType *DataDatabricksJobJobSettingsSettingsTaskNewClusterWorkloadType `field:"optional" json:"workloadType" yaml:"workloadType"`
}

