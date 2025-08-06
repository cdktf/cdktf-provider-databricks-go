// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobEnvironmentSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#client Job#client}.
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#dependencies Job#dependencies}.
	Dependencies *[]*string `field:"optional" json:"dependencies" yaml:"dependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#environment_version Job#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/job#jar_dependencies Job#jar_dependencies}.
	JarDependencies *[]*string `field:"optional" json:"jarDependencies" yaml:"jarDependencies"`
}

