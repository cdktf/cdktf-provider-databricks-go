// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabrickscluster


type DataDatabricksClusterClusterInfoSpecLibraryMaven struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/cluster#coordinates DataDatabricksCluster#coordinates}.
	Coordinates *string `field:"required" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/cluster#exclusions DataDatabricksCluster#exclusions}.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.1/docs/data-sources/cluster#repo DataDatabricksCluster#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

