// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomscleanroom


type CleanRoomsCleanRoomRemoteDetailedInfoComplianceSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#compliance_standards CleanRoomsCleanRoom#compliance_standards}.
	ComplianceStandards *[]*string `field:"optional" json:"complianceStandards" yaml:"complianceStandards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.89.0/docs/resources/clean_rooms_clean_room#is_enabled CleanRoomsCleanRoom#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

