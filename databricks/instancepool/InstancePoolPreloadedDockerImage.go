// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instancepool


type InstancePoolPreloadedDockerImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/instance_pool#url InstancePool#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.30.0/docs/resources/instance_pool#basic_auth InstancePool#basic_auth}
	BasicAuth *InstancePoolPreloadedDockerImageBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

