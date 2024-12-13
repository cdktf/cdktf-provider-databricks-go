// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#cluster_size SqlEndpoint#cluster_size}.
	ClusterSize *string `field:"required" json:"clusterSize" yaml:"clusterSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#name SqlEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#auto_stop_mins SqlEndpoint#auto_stop_mins}.
	AutoStopMins *float64 `field:"optional" json:"autoStopMins" yaml:"autoStopMins"`
	// channel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#channel SqlEndpoint#channel}
	Channel *SqlEndpointChannel `field:"optional" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#data_source_id SqlEndpoint#data_source_id}.
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#enable_photon SqlEndpoint#enable_photon}.
	EnablePhoton interface{} `field:"optional" json:"enablePhoton" yaml:"enablePhoton"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#enable_serverless_compute SqlEndpoint#enable_serverless_compute}.
	EnableServerlessCompute interface{} `field:"optional" json:"enableServerlessCompute" yaml:"enableServerlessCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#id SqlEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#instance_profile_arn SqlEndpoint#instance_profile_arn}.
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#max_num_clusters SqlEndpoint#max_num_clusters}.
	MaxNumClusters *float64 `field:"optional" json:"maxNumClusters" yaml:"maxNumClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#min_num_clusters SqlEndpoint#min_num_clusters}.
	MinNumClusters *float64 `field:"optional" json:"minNumClusters" yaml:"minNumClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#spot_instance_policy SqlEndpoint#spot_instance_policy}.
	SpotInstancePolicy *string `field:"optional" json:"spotInstancePolicy" yaml:"spotInstancePolicy"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#tags SqlEndpoint#tags}
	Tags *SqlEndpointTags `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#timeouts SqlEndpoint#timeouts}
	Timeouts *SqlEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/resources/sql_endpoint#warehouse_type SqlEndpoint#warehouse_type}.
	WarehouseType *string `field:"optional" json:"warehouseType" yaml:"warehouseType"`
}

