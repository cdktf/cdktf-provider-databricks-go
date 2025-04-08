// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskLibraryMaven struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/data-sources/job#coordinates DataDatabricksJob#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/data-sources/job#exclusions DataDatabricksJob#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.72.0/docs/data-sources/job#repo DataDatabricksJob#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

