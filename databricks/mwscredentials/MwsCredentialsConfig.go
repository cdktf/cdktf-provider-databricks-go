// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwscredentials

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MwsCredentialsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#credentials_name MwsCredentials#credentials_name}.
	CredentialsName *string `field:"required" json:"credentialsName" yaml:"credentialsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#role_arn MwsCredentials#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#account_id MwsCredentials#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#creation_time MwsCredentials#creation_time}.
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#credentials_id MwsCredentials#credentials_id}.
	CredentialsId *string `field:"optional" json:"credentialsId" yaml:"credentialsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#external_id MwsCredentials#external_id}.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.36.2/docs/resources/mws_credentials#id MwsCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

