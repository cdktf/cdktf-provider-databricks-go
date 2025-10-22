// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworks


type MwsNetworksGcpNetworkInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#network_project_id MwsNetworks#network_project_id}.
	NetworkProjectId *string `field:"required" json:"networkProjectId" yaml:"networkProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#subnet_id MwsNetworks#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#subnet_region MwsNetworks#subnet_region}.
	SubnetRegion *string `field:"required" json:"subnetRegion" yaml:"subnetRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#vpc_id MwsNetworks#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#pod_ip_range_name MwsNetworks#pod_ip_range_name}.
	PodIpRangeName *string `field:"optional" json:"podIpRangeName" yaml:"podIpRangeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.94.0/docs/resources/mws_networks#service_ip_range_name MwsNetworks#service_ip_range_name}.
	ServiceIpRangeName *string `field:"optional" json:"serviceIpRangeName" yaml:"serviceIpRangeName"`
}

