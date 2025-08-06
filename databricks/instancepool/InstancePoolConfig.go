// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstancePoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#idle_instance_autotermination_minutes InstancePool#idle_instance_autotermination_minutes}.
	IdleInstanceAutoterminationMinutes *float64 `field:"required" json:"idleInstanceAutoterminationMinutes" yaml:"idleInstanceAutoterminationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#instance_pool_name InstancePool#instance_pool_name}.
	InstancePoolName *string `field:"required" json:"instancePoolName" yaml:"instancePoolName"`
	// aws_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#aws_attributes InstancePool#aws_attributes}
	AwsAttributes *InstancePoolAwsAttributes `field:"optional" json:"awsAttributes" yaml:"awsAttributes"`
	// azure_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#azure_attributes InstancePool#azure_attributes}
	AzureAttributes *InstancePoolAzureAttributes `field:"optional" json:"azureAttributes" yaml:"azureAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#custom_tags InstancePool#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#disk_spec InstancePool#disk_spec}
	DiskSpec *InstancePoolDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#enable_elastic_disk InstancePool#enable_elastic_disk}.
	EnableElasticDisk interface{} `field:"optional" json:"enableElasticDisk" yaml:"enableElasticDisk"`
	// gcp_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#gcp_attributes InstancePool#gcp_attributes}
	GcpAttributes *InstancePoolGcpAttributes `field:"optional" json:"gcpAttributes" yaml:"gcpAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#id InstancePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_pool_fleet_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#instance_pool_fleet_attributes InstancePool#instance_pool_fleet_attributes}
	InstancePoolFleetAttributes *InstancePoolInstancePoolFleetAttributes `field:"optional" json:"instancePoolFleetAttributes" yaml:"instancePoolFleetAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#instance_pool_id InstancePool#instance_pool_id}.
	InstancePoolId *string `field:"optional" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#max_capacity InstancePool#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#min_idle_instances InstancePool#min_idle_instances}.
	MinIdleInstances *float64 `field:"optional" json:"minIdleInstances" yaml:"minIdleInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#node_type_id InstancePool#node_type_id}.
	NodeTypeId *string `field:"optional" json:"nodeTypeId" yaml:"nodeTypeId"`
	// preloaded_docker_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#preloaded_docker_image InstancePool#preloaded_docker_image}
	PreloadedDockerImage interface{} `field:"optional" json:"preloadedDockerImage" yaml:"preloadedDockerImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/instance_pool#preloaded_spark_versions InstancePool#preloaded_spark_versions}.
	PreloadedSparkVersions *[]*string `field:"optional" json:"preloadedSparkVersions" yaml:"preloadedSparkVersions"`
}

