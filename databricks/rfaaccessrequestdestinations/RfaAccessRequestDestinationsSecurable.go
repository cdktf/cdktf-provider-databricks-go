// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rfaaccessrequestdestinations


type RfaAccessRequestDestinationsSecurable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/rfa_access_request_destinations#full_name RfaAccessRequestDestinations#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/rfa_access_request_destinations#provider_share RfaAccessRequestDestinations#provider_share}.
	ProviderShare *string `field:"optional" json:"providerShare" yaml:"providerShare"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.97.0/docs/resources/rfa_access_request_destinations#type RfaAccessRequestDestinations#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

