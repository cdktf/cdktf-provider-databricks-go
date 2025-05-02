// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontrolruleset


type AccessControlRuleSetGrantRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/resources/access_control_rule_set#role AccessControlRuleSet#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.76.0/docs/resources/access_control_rule_set#principals AccessControlRuleSet#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

