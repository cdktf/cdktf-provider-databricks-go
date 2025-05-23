// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customappintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomAppIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#client_id CustomAppIntegration#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#client_secret CustomAppIntegration#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#confidential CustomAppIntegration#confidential}.
	Confidential interface{} `field:"optional" json:"confidential" yaml:"confidential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#created_by CustomAppIntegration#created_by}.
	CreatedBy *float64 `field:"optional" json:"createdBy" yaml:"createdBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#create_time CustomAppIntegration#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#creator_username CustomAppIntegration#creator_username}.
	CreatorUsername *string `field:"optional" json:"creatorUsername" yaml:"creatorUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#id CustomAppIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#integration_id CustomAppIntegration#integration_id}.
	IntegrationId *string `field:"optional" json:"integrationId" yaml:"integrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#name CustomAppIntegration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#redirect_urls CustomAppIntegration#redirect_urls}.
	RedirectUrls *[]*string `field:"optional" json:"redirectUrls" yaml:"redirectUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#scopes CustomAppIntegration#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// token_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#token_access_policy CustomAppIntegration#token_access_policy}
	TokenAccessPolicy *CustomAppIntegrationTokenAccessPolicy `field:"optional" json:"tokenAccessPolicy" yaml:"tokenAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.0/docs/resources/custom_app_integration#user_authorized_scopes CustomAppIntegration#user_authorized_scopes}.
	UserAuthorizedScopes *[]*string `field:"optional" json:"userAuthorizedScopes" yaml:"userAuthorizedScopes"`
}

