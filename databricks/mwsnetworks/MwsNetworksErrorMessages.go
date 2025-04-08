// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mwsnetworks


type MwsNetworksErrorMessages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/mws_networks#error_message MwsNetworks#error_message}.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/resources/mws_networks#error_type MwsNetworks#error_type}.
	ErrorType *string `field:"optional" json:"errorType" yaml:"errorType"`
}

