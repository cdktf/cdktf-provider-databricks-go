// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatabricksjob


type DataDatabricksJobJobSettingsSettingsEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.0/docs/data-sources/job#environment_key DataDatabricksJob#environment_key}.
	EnvironmentKey *string `field:"required" json:"environmentKey" yaml:"environmentKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.64.0/docs/data-sources/job#spec DataDatabricksJob#spec}
	Spec *DataDatabricksJobJobSettingsSettingsEnvironmentSpec `field:"optional" json:"spec" yaml:"spec"`
}

