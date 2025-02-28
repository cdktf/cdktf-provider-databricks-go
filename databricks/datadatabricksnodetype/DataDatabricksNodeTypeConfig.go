// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksnodetype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatabricksNodeTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#category DataDatabricksNodeType#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#fleet DataDatabricksNodeType#fleet}.
	Fleet interface{} `field:"optional" json:"fleet" yaml:"fleet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#gb_per_core DataDatabricksNodeType#gb_per_core}.
	GbPerCore *float64 `field:"optional" json:"gbPerCore" yaml:"gbPerCore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#graviton DataDatabricksNodeType#graviton}.
	Graviton interface{} `field:"optional" json:"graviton" yaml:"graviton"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#id DataDatabricksNodeType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#is_io_cache_enabled DataDatabricksNodeType#is_io_cache_enabled}.
	IsIoCacheEnabled interface{} `field:"optional" json:"isIoCacheEnabled" yaml:"isIoCacheEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#local_disk DataDatabricksNodeType#local_disk}.
	LocalDisk interface{} `field:"optional" json:"localDisk" yaml:"localDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#local_disk_min_size DataDatabricksNodeType#local_disk_min_size}.
	LocalDiskMinSize *float64 `field:"optional" json:"localDiskMinSize" yaml:"localDiskMinSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#min_cores DataDatabricksNodeType#min_cores}.
	MinCores *float64 `field:"optional" json:"minCores" yaml:"minCores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#min_gpus DataDatabricksNodeType#min_gpus}.
	MinGpus *float64 `field:"optional" json:"minGpus" yaml:"minGpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#min_memory_gb DataDatabricksNodeType#min_memory_gb}.
	MinMemoryGb *float64 `field:"optional" json:"minMemoryGb" yaml:"minMemoryGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#photon_driver_capable DataDatabricksNodeType#photon_driver_capable}.
	PhotonDriverCapable interface{} `field:"optional" json:"photonDriverCapable" yaml:"photonDriverCapable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#photon_worker_capable DataDatabricksNodeType#photon_worker_capable}.
	PhotonWorkerCapable interface{} `field:"optional" json:"photonWorkerCapable" yaml:"photonWorkerCapable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.68.0/docs/data-sources/node_type#support_port_forwarding DataDatabricksNodeType#support_port_forwarding}.
	SupportPortForwarding interface{} `field:"optional" json:"supportPortForwarding" yaml:"supportPortForwarding"`
}

