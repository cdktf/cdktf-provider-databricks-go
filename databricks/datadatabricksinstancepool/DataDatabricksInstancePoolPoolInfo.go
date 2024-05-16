// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstancepool


type DataDatabricksInstancePoolPoolInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#idle_instance_autotermination_minutes DataDatabricksInstancePool#idle_instance_autotermination_minutes}.
	IdleInstanceAutoterminationMinutes *float64 `field:"required" json:"idleInstanceAutoterminationMinutes" yaml:"idleInstanceAutoterminationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#instance_pool_name DataDatabricksInstancePool#instance_pool_name}.
	InstancePoolName *string `field:"required" json:"instancePoolName" yaml:"instancePoolName"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#aws_attributes DataDatabricksInstancePool#aws_attributes}
	AwsAttributes *DataDatabricksInstancePoolPoolInfoAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#azure_attributes DataDatabricksInstancePool#azure_attributes}
	AzureAttributes *DataDatabricksInstancePoolPoolInfoAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#custom_tags DataDatabricksInstancePool#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#default_tags DataDatabricksInstancePool#default_tags}.
	DefaultTags *map[string]*string `field:"optional" json:"defaultTags" yaml:"defaultTags"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#disk_spec DataDatabricksInstancePool#disk_spec}
	DiskSpec *DataDatabricksInstancePoolPoolInfoDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#enable_elastic_disk DataDatabricksInstancePool#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#gcp_attributes DataDatabricksInstancePool#gcp_attributes}
	GcpAttributes *DataDatabricksInstancePoolPoolInfoGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// instance_pool_fleet_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#instance_pool_fleet_attributes DataDatabricksInstancePool#instance_pool_fleet_attributes}
	InstancePoolFleetAttributes interface{} `field:"optional" json:"instancePoolFleetAttributes" yaml:"instancePoolFleetAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#instance_pool_id DataDatabricksInstancePool#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#max_capacity DataDatabricksInstancePool#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#min_idle_instances DataDatabricksInstancePool#min_idle_instances}.
	MinIdleInstances *float64 `field:"optional" json:"minIdleInstances" yaml:"minIdleInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#node_type_id DataDatabricksInstancePool#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// preloaded_docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#preloaded_docker_image DataDatabricksInstancePool#preloaded_docker_image}
	PreloadedDockerImage interface{} `field:"optional" json:"preloadedDockerImage" yaml:"preloadedDockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#preloaded_spark_versions DataDatabricksInstancePool#preloaded_spark_versions}.
	PreloadedSparkVersions *[]*string `field:"optional" json:"preloadedSparkVersions" yaml:"preloadedSparkVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#state DataDatabricksInstancePool#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// stats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.44.0/docs/data-sources/instance_pool#stats DataDatabricksInstancePool#stats}
	Stats *DataDatabricksInstancePoolPoolInfoStats `field:"optional" json:"stats" yaml:"stats"`
}

