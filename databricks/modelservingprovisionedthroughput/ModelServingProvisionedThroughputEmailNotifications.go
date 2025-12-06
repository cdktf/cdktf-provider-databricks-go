// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelservingprovisionedthroughput


type ModelServingProvisionedThroughputEmailNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#on_update_failure ModelServingProvisionedThroughput#on_update_failure}.
	OnUpdateFailure *[]*string `field:"optional" json:"onUpdateFailure" yaml:"onUpdateFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.99.0/docs/resources/model_serving_provisioned_throughput#on_update_success ModelServingProvisionedThroughput#on_update_success}.
	OnUpdateSuccess *[]*string `field:"optional" json:"onUpdateSuccess" yaml:"onUpdateSuccess"`
}

