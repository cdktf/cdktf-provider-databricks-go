// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#url DataDatabricksJob#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.85.0/docs/data-sources/job#basic_auth DataDatabricksJob#basic_auth}
	BasicAuth *DataDatabricksJobJobSettingsSettingsTaskForEachTaskTaskNewClusterDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

