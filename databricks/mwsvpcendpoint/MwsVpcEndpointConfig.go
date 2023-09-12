// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsvpcendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsVpcEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#vpc_endpoint_name MwsVpcEndpoint#vpc_endpoint_name}.
	VpcEndpointName *string `field:"required" json:"vpcEndpointName" yaml:"vpcEndpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#account_id MwsVpcEndpoint#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#aws_account_id MwsVpcEndpoint#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#aws_endpoint_service_id MwsVpcEndpoint#aws_endpoint_service_id}.
	AwsEndpointServiceId *string `field:"optional" json:"awsEndpointServiceId" yaml:"awsEndpointServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#aws_vpc_endpoint_id MwsVpcEndpoint#aws_vpc_endpoint_id}.
	AwsVpcEndpointId *string `field:"optional" json:"awsVpcEndpointId" yaml:"awsVpcEndpointId"`
	// gcp_vpc_endpoint_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#gcp_vpc_endpoint_info MwsVpcEndpoint#gcp_vpc_endpoint_info}
	GcpVpcEndpointInfo *MwsVpcEndpointGcpVpcEndpointInfo `field:"optional" json:"gcpVpcEndpointInfo" yaml:"gcpVpcEndpointInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#id MwsVpcEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#region MwsVpcEndpoint#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#state MwsVpcEndpoint#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#use_case MwsVpcEndpoint#use_case}.
	UseCase *string `field:"optional" json:"useCase" yaml:"useCase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.25.0/docs/resources/mws_vpc_endpoint#vpc_endpoint_id MwsVpcEndpoint#vpc_endpoint_id}.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

