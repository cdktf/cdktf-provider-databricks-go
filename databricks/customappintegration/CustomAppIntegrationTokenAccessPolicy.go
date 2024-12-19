// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customappintegration


type CustomAppIntegrationTokenAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/custom_app_integration#access_token_ttl_in_minutes CustomAppIntegration#access_token_ttl_in_minutes}.
	AccessTokenTtlInMinutes *float64 `field:"optional" json:"accessTokenTtlInMinutes" yaml:"accessTokenTtlInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.62.0/docs/resources/custom_app_integration#refresh_token_ttl_in_minutes CustomAppIntegration#refresh_token_ttl_in_minutes}.
	RefreshTokenTtlInMinutes *float64 `field:"optional" json:"refreshTokenTtlInMinutes" yaml:"refreshTokenTtlInMinutes"`
}

