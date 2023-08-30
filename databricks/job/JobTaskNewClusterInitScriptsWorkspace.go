// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobTaskNewClusterInitScriptsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/resources/job#destination Job#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

