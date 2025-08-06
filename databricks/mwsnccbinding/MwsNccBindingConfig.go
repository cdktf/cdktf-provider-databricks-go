// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnccbinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNccBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_ncc_binding#network_connectivity_config_id MwsNccBinding#network_connectivity_config_id}.
	NetworkConnectivityConfigId *string `field:"required" json:"networkConnectivityConfigId" yaml:"networkConnectivityConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_ncc_binding#workspace_id MwsNccBinding#workspace_id}.
	WorkspaceId *float64 `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/mws_ncc_binding#id MwsNccBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

