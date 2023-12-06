// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretscope


type SecretScopeKeyvaultMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/secret_scope#dns_name SecretScope#dns_name}.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.31.0/docs/resources/secret_scope#resource_id SecretScope#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

