package provider


type DatabricksProviderConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#account_id DatabricksProvider#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#alias DatabricksProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#auth_type DatabricksProvider#auth_type}.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_client_id DatabricksProvider#azure_client_id}.
	AzureClientId *string `field:"optional" json:"azureClientId" yaml:"azureClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_client_secret DatabricksProvider#azure_client_secret}.
	AzureClientSecret *string `field:"optional" json:"azureClientSecret" yaml:"azureClientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_environment DatabricksProvider#azure_environment}.
	AzureEnvironment *string `field:"optional" json:"azureEnvironment" yaml:"azureEnvironment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_login_app_id DatabricksProvider#azure_login_app_id}.
	AzureLoginAppId *string `field:"optional" json:"azureLoginAppId" yaml:"azureLoginAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_tenant_id DatabricksProvider#azure_tenant_id}.
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_use_msi DatabricksProvider#azure_use_msi}.
	AzureUseMsi interface{} `field:"optional" json:"azureUseMsi" yaml:"azureUseMsi"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#azure_workspace_resource_id DatabricksProvider#azure_workspace_resource_id}.
	AzureWorkspaceResourceId *string `field:"optional" json:"azureWorkspaceResourceId" yaml:"azureWorkspaceResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#config_file DatabricksProvider#config_file}.
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#debug_headers DatabricksProvider#debug_headers}.
	DebugHeaders interface{} `field:"optional" json:"debugHeaders" yaml:"debugHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#debug_truncate_bytes DatabricksProvider#debug_truncate_bytes}.
	DebugTruncateBytes *float64 `field:"optional" json:"debugTruncateBytes" yaml:"debugTruncateBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#google_credentials DatabricksProvider#google_credentials}.
	GoogleCredentials *string `field:"optional" json:"googleCredentials" yaml:"googleCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#google_service_account DatabricksProvider#google_service_account}.
	GoogleServiceAccount *string `field:"optional" json:"googleServiceAccount" yaml:"googleServiceAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#host DatabricksProvider#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#http_timeout_seconds DatabricksProvider#http_timeout_seconds}.
	HttpTimeoutSeconds *float64 `field:"optional" json:"httpTimeoutSeconds" yaml:"httpTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#password DatabricksProvider#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#profile DatabricksProvider#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#rate_limit DatabricksProvider#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#skip_verify DatabricksProvider#skip_verify}.
	SkipVerify interface{} `field:"optional" json:"skipVerify" yaml:"skipVerify"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#token DatabricksProvider#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/databricks#username DatabricksProvider#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

