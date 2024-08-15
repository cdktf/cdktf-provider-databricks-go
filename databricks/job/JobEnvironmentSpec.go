// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/job#client Job#client}.
	Client *string `field:"required" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.50.0/docs/resources/job#dependencies Job#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
}

