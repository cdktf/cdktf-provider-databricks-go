// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertv2


type AlertV2EvaluationNotificationSubscriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/alert_v2#destination_id AlertV2#destination_id}.
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.81.1/docs/resources/alert_v2#user_email AlertV2#user_email}.
	UserEmail *string `field:"optional" json:"userEmail" yaml:"userEmail"`
}

