// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksinstanceprofiles


type DataDatabricksInstanceProfilesInstanceProfiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/data-sources/instance_profiles#arn DataDatabricksInstanceProfiles#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/data-sources/instance_profiles#is_meta DataDatabricksInstanceProfiles#is_meta}.
	IsMeta interface{} `field:"optional" json:"isMeta" yaml:"isMeta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/data-sources/instance_profiles#name DataDatabricksInstanceProfiles#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.49.1/docs/data-sources/instance_profiles#role_arn DataDatabricksInstanceProfiles#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

