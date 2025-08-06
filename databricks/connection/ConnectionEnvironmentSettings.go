// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connection


type ConnectionEnvironmentSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/connection#environment_version Connection#environment_version}.
	EnvironmentVersion *string `field:"optional" json:"environmentVersion" yaml:"environmentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.86.0/docs/resources/connection#java_dependencies Connection#java_dependencies}.
	JavaDependencies *[]*string `field:"optional" json:"javaDependencies" yaml:"javaDependencies"`
}

