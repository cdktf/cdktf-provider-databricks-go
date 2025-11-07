// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#apply_policy_default_values Pipeline#apply_policy_default_values}.
	ApplyPolicyDefaultValues interface{} `field:"optional" json:"applyPolicyDefaultValues" yaml:"applyPolicyDefaultValues"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#autoscale Pipeline#autoscale}
	Autoscale *PipelineClusterAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#aws_attributes Pipeline#aws_attributes}
	AwsAttributes *PipelineClusterAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#azure_attributes Pipeline#azure_attributes}
	AzureAttributes *PipelineClusterAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// cluster_log_conf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#cluster_log_conf Pipeline#cluster_log_conf}
	ClusterLogConf *PipelineClusterClusterLogConf `field:"optional" json:"clusterLogConf" yaml:"clusterLogConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#custom_tags Pipeline#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#driver_instance_pool_id Pipeline#driver_instance_pool_id}.
	DriverInstancePoolId *string `field:"optional" json:"driverInstancePoolId" yaml:"driverInstancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#driver_node_type_id Pipeline#driver_node_type_id}.
	DriverNodeTypeId *string `field:"optional" json:"driverNodeTypeId" yaml:"driverNodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#enable_local_disk_encryption Pipeline#enable_local_disk_encryption}.
	EnableLocalDiskEncryption interface{} `field:"optional" json:"enableLocalDiskEncryption" yaml:"enableLocalDiskEncryption"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#gcp_attributes Pipeline#gcp_attributes}
	GcpAttributes *PipelineClusterGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// init_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#init_scripts Pipeline#init_scripts}
	InitScripts interface{} `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#instance_pool_id Pipeline#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#label Pipeline#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#node_type_id Pipeline#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#num_workers Pipeline#num_workers}.
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#policy_id Pipeline#policy_id}.
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#spark_conf Pipeline#spark_conf}.
	SparkConf *map[string]*string `field:"optional" json:"sparkConf" yaml:"sparkConf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#spark_env_vars Pipeline#spark_env_vars}.
	SparkEnvVars *map[string]*string `field:"optional" json:"sparkEnvVars" yaml:"sparkEnvVars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/pipeline#ssh_public_keys Pipeline#ssh_public_keys}.
	SshPublicKeys *[]*string `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
}

