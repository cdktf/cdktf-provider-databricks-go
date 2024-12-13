// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapp


type DataDatabricksAppApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#name DataDatabricksApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#active_deployment DataDatabricksApp#active_deployment}.
	ActiveDeployment *DataDatabricksAppAppActiveDeployment `field:"optional" json:"activeDeployment" yaml:"activeDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#app_status DataDatabricksApp#app_status}.
	AppStatus *DataDatabricksAppAppAppStatus `field:"optional" json:"appStatus" yaml:"appStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#compute_status DataDatabricksApp#compute_status}.
	ComputeStatus *DataDatabricksAppAppComputeStatus `field:"optional" json:"computeStatus" yaml:"computeStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#create_time DataDatabricksApp#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#creator DataDatabricksApp#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#default_source_code_path DataDatabricksApp#default_source_code_path}.
	DefaultSourceCodePath *string `field:"optional" json:"defaultSourceCodePath" yaml:"defaultSourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#description DataDatabricksApp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#pending_deployment DataDatabricksApp#pending_deployment}.
	PendingDeployment *DataDatabricksAppAppPendingDeployment `field:"optional" json:"pendingDeployment" yaml:"pendingDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#resources DataDatabricksApp#resources}.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#service_principal_client_id DataDatabricksApp#service_principal_client_id}.
	ServicePrincipalClientId *string `field:"optional" json:"servicePrincipalClientId" yaml:"servicePrincipalClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#service_principal_id DataDatabricksApp#service_principal_id}.
	ServicePrincipalId *float64 `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#service_principal_name DataDatabricksApp#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#updater DataDatabricksApp#updater}.
	Updater *string `field:"optional" json:"updater" yaml:"updater"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#update_time DataDatabricksApp#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/app#url DataDatabricksApp#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

