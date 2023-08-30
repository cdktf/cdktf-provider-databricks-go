// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsJobClusterNewClusterDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/data-sources/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.24.1/docs/data-sources/job#basic_auth DataDatabricksJob#basic_auth}
	BasicAuth *DataDatabricksJobJobSettingsSettingsJobClusterNewClusterDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

