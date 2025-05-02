// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsworkspaces


type MwsWorkspacesGcpManagedNetworkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/resources/mws_workspaces#subnet_cidr MwsWorkspaces#subnet_cidr}.
	SubnetCidr *string `field:"required" json:"subnetCidr" yaml:"subnetCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/resources/mws_workspaces#gke_cluster_pod_ip_range MwsWorkspaces#gke_cluster_pod_ip_range}.
	GkeClusterPodIpRange *string `field:"optional" json:"gkeClusterPodIpRange" yaml:"gkeClusterPodIpRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/resources/mws_workspaces#gke_cluster_service_ip_range MwsWorkspaces#gke_cluster_service_ip_range}.
	GkeClusterServiceIpRange *string `field:"optional" json:"gkeClusterServiceIpRange" yaml:"gkeClusterServiceIpRange"`
}

