// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DatabricksProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#account_id DatabricksProvider#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#actions_id_token_request_token DatabricksProvider#actions_id_token_request_token}.
	ActionsIdTokenRequestToken *string `field:"optional" json:"actionsIdTokenRequestToken" yaml:"actionsIdTokenRequestToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#actions_id_token_request_url DatabricksProvider#actions_id_token_request_url}.
	ActionsIdTokenRequestUrl *string `field:"optional" json:"actionsIdTokenRequestUrl" yaml:"actionsIdTokenRequestUrl"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#alias DatabricksProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#auth_type DatabricksProvider#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_client_id DatabricksProvider#azure_client_id}.
	AzureClientId *string `field:"optional" json:"azureClientId" yaml:"azureClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_client_secret DatabricksProvider#azure_client_secret}.
	AzureClientSecret *string `field:"optional" json:"azureClientSecret" yaml:"azureClientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_environment DatabricksProvider#azure_environment}.
	AzureEnvironment *string `field:"optional" json:"azureEnvironment" yaml:"azureEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_login_app_id DatabricksProvider#azure_login_app_id}.
	AzureLoginAppId *string `field:"optional" json:"azureLoginAppId" yaml:"azureLoginAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_tenant_id DatabricksProvider#azure_tenant_id}.
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_use_msi DatabricksProvider#azure_use_msi}.
	AzureUseMsi interface{} `field:"optional" json:"azureUseMsi" yaml:"azureUseMsi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#azure_workspace_resource_id DatabricksProvider#azure_workspace_resource_id}.
	AzureWorkspaceResourceId *string `field:"optional" json:"azureWorkspaceResourceId" yaml:"azureWorkspaceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#client_id DatabricksProvider#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#client_secret DatabricksProvider#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#cluster_id DatabricksProvider#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#config_file DatabricksProvider#config_file}.
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#databricks_cli_path DatabricksProvider#databricks_cli_path}.
	DatabricksCliPath *string `field:"optional" json:"databricksCliPath" yaml:"databricksCliPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#debug_headers DatabricksProvider#debug_headers}.
	DebugHeaders interface{} `field:"optional" json:"debugHeaders" yaml:"debugHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#debug_truncate_bytes DatabricksProvider#debug_truncate_bytes}.
	DebugTruncateBytes *float64 `field:"optional" json:"debugTruncateBytes" yaml:"debugTruncateBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#google_credentials DatabricksProvider#google_credentials}.
	GoogleCredentials *string `field:"optional" json:"googleCredentials" yaml:"googleCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#google_service_account DatabricksProvider#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#host DatabricksProvider#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#http_timeout_seconds DatabricksProvider#http_timeout_seconds}.
	HttpTimeoutSeconds *float64 `field:"optional" json:"httpTimeoutSeconds" yaml:"httpTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#metadata_service_url DatabricksProvider#metadata_service_url}.
	MetadataServiceUrl *string `field:"optional" json:"metadataServiceUrl" yaml:"metadataServiceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#password DatabricksProvider#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#profile DatabricksProvider#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#rate_limit DatabricksProvider#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#retry_timeout_seconds DatabricksProvider#retry_timeout_seconds}.
	RetryTimeoutSeconds *float64 `field:"optional" json:"retryTimeoutSeconds" yaml:"retryTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#serverless_compute_id DatabricksProvider#serverless_compute_id}.
	ServerlessComputeId *string `field:"optional" json:"serverlessComputeId" yaml:"serverlessComputeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#skip_verify DatabricksProvider#skip_verify}.
	SkipVerify interface{} `field:"optional" json:"skipVerify" yaml:"skipVerify"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#token DatabricksProvider#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#username DatabricksProvider#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.52.0/docs#warehouse_id DatabricksProvider#warehouse_id}.
	WarehouseId *string `field:"optional" json:"warehouseId" yaml:"warehouseId"`
}

