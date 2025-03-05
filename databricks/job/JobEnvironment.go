// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/job#environment_key Job#environment_key}.
	EnvironmentKey *string `field:"required" json:"environmentKey" yaml:"environmentKey"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.69.0/docs/resources/job#spec Job#spec}
	Spec *JobEnvironmentSpec `field:"optional" json:"spec" yaml:"spec"`
}

