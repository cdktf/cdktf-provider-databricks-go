// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsvpcendpoint


type MwsVpcEndpointGcpVpcEndpointInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/mws_vpc_endpoint#endpoint_region MwsVpcEndpoint#endpoint_region}.
	EndpointRegion *string `field:"required" json:"endpointRegion" yaml:"endpointRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/mws_vpc_endpoint#project_id MwsVpcEndpoint#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/mws_vpc_endpoint#psc_endpoint_name MwsVpcEndpoint#psc_endpoint_name}.
	PscEndpointName *string `field:"required" json:"pscEndpointName" yaml:"pscEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/mws_vpc_endpoint#psc_connection_id MwsVpcEndpoint#psc_connection_id}.
	PscConnectionId *string `field:"optional" json:"pscConnectionId" yaml:"pscConnectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/mws_vpc_endpoint#service_attachment_id MwsVpcEndpoint#service_attachment_id}.
	ServiceAttachmentId *string `field:"optional" json:"serviceAttachmentId" yaml:"serviceAttachmentId"`
}

