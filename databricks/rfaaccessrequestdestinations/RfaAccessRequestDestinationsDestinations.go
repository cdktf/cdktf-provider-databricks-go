// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rfaaccessrequestdestinations


type RfaAccessRequestDestinationsDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/rfa_access_request_destinations#destination_id RfaAccessRequestDestinations#destination_id}.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/rfa_access_request_destinations#destination_type RfaAccessRequestDestinations#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.93.0/docs/resources/rfa_access_request_destinations#special_destination RfaAccessRequestDestinations#special_destination}.
	SpecialDestination *string `field:"optional" json:"specialDestination" yaml:"specialDestination"`
}

