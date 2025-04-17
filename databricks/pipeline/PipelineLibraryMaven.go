// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipeline


type PipelineLibraryMaven struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/pipeline#coordinates Pipeline#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/pipeline#exclusions Pipeline#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.74.0/docs/resources/pipeline#repo Pipeline#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

