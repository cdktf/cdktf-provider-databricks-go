// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworkconnectivityconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsNetworkConnectivityConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#name MwsNetworkConnectivityConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#region MwsNetworkConnectivityConfig#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#account_id MwsNetworkConnectivityConfig#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#creation_time MwsNetworkConnectivityConfig#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// egress_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#egress_config MwsNetworkConnectivityConfig#egress_config}
	EgressConfig *MwsNetworkConnectivityConfigEgressConfig `field:"optional" json:"egressConfig" yaml:"egressConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#id MwsNetworkConnectivityConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#network_connectivity_config_id MwsNetworkConnectivityConfig#network_connectivity_config_id}.
	NetworkConnectivityConfigId *string `field:"optional" json:"networkConnectivityConfigId" yaml:"networkConnectivityConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.42.0/docs/resources/mws_network_connectivity_config#updated_time MwsNetworkConnectivityConfig#updated_time}.
	UpdatedTime *float64 `field:"optional" json:"updatedTime" yaml:"updatedTime"`
}

