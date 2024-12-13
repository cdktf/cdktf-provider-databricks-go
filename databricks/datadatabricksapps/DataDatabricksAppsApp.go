// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksapps


type DataDatabricksAppsApp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#name DataDatabricksApps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#active_deployment DataDatabricksApps#active_deployment}.
	ActiveDeployment *DataDatabricksAppsAppActiveDeployment `field:"optional" json:"activeDeployment" yaml:"activeDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#app_status DataDatabricksApps#app_status}.
	AppStatus *DataDatabricksAppsAppAppStatus `field:"optional" json:"appStatus" yaml:"appStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#compute_status DataDatabricksApps#compute_status}.
	ComputeStatus *DataDatabricksAppsAppComputeStatus `field:"optional" json:"computeStatus" yaml:"computeStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#create_time DataDatabricksApps#create_time}.
	CreateTime *string `field:"optional" json:"createTime" yaml:"createTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#creator DataDatabricksApps#creator}.
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#default_source_code_path DataDatabricksApps#default_source_code_path}.
	DefaultSourceCodePath *string `field:"optional" json:"defaultSourceCodePath" yaml:"defaultSourceCodePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#description DataDatabricksApps#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#pending_deployment DataDatabricksApps#pending_deployment}.
	PendingDeployment *DataDatabricksAppsAppPendingDeployment `field:"optional" json:"pendingDeployment" yaml:"pendingDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#resources DataDatabricksApps#resources}.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#service_principal_client_id DataDatabricksApps#service_principal_client_id}.
	ServicePrincipalClientId *string `field:"optional" json:"servicePrincipalClientId" yaml:"servicePrincipalClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#service_principal_id DataDatabricksApps#service_principal_id}.
	ServicePrincipalId *float64 `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#service_principal_name DataDatabricksApps#service_principal_name}.
	ServicePrincipalName *string `field:"optional" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#updater DataDatabricksApps#updater}.
	Updater *string `field:"optional" json:"updater" yaml:"updater"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#update_time DataDatabricksApps#update_time}.
	UpdateTime *string `field:"optional" json:"updateTime" yaml:"updateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.61.0/docs/data-sources/apps#url DataDatabricksApps#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

