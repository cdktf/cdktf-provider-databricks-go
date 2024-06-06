// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package compliancesecurityprofileworkspacesetting


type ComplianceSecurityProfileWorkspaceSettingComplianceSecurityProfileWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/resources/compliance_security_profile_workspace_setting#compliance_standards ComplianceSecurityProfileWorkspaceSetting#compliance_standards}.
	ComplianceStandards *[]*string `field:"optional" json:"complianceStandards" yaml:"complianceStandards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.47.0/docs/resources/compliance_security_profile_workspace_setting#is_enabled ComplianceSecurityProfileWorkspaceSetting#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

