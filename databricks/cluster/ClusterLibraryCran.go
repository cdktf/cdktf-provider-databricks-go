// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterLibraryCran struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/cluster#package Cluster#package}.
	Package *string `field:"required" json:"package" yaml:"package"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.37.1/docs/resources/cluster#repo Cluster#repo}.
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

