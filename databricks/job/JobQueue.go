// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobQueue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.34.0/docs/resources/job#enabled Job#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

